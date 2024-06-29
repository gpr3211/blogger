package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"sync"

	"github.com/google/uuid"
	"github.com/gpr3211/blogger/internal/clog"
	"github.com/gpr3211/blogger/internal/database"
)

func startScrapeWorker(
	db *database.Queries,
	concurrency int,
	periodRequest time.Duration,
) {
	clog.Printf("Scarping on %v routines every %s duration", concurrency, periodRequest)
	ticker := time.NewTicker(periodRequest)
	// for loops fires immediately and waits for ticker
	for ; ; <-ticker.C {
		feeds, err := db.MakeFetchList(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			clog.Fatal("Failed to create fetc")
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()

	}

}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		clog.Println("Err marking feed as fetched")
		return
	}
	rssFeed, err := urltoFeed(feed.Url)
	if err != nil {
		clog.Println("err fetching feed")
		return
	}
	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true

		} else {
			clog.Println("Empty item description")
			return
		}

		pubic, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			clog.Println("failed parse of post string to time")
			return
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubic,
			FeedID:      feed.ID,
		})
		if err != nil {
			clog.Printf("Failed to write post to DB\n,")
			return
		}
		clog.Printf("post fetched")
	}
	log.Printf("Feed %s collected, %v posts,\n", feed.Name, len(rssFeed.Channel.Item))
	return
}
