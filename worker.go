package main

import (
	"context"
	"log"
	"time"

	"sync"

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
		clog.Print("Found post", item.Title)
		return
	}
	log.Printf("Feed %s collected, %v posts,\n", feed.Name, len(rssFeed.Channel.Item))

}
