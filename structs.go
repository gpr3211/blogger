package main

import (
	//	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gpr3211/blogger/internal/database"
)

// FOLLOWS
type Follow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedId    uuid.UUID `json:"feed_id"`
	UserId    uuid.UUID `json:"user_id"`
}

func dbToFollow(follow database.Follow) Follow {
	return Follow{
		ID:        follow.ID,
		CreatedAt: follow.CreatedAt,
		UpdatedAt: follow.UpdatedAt,
		FeedId:    follow.FeedID,
		UserId:    follow.UserID,
	}
}

func FollowToFollows(follows []database.Follow) []Follow {
	results := make([]Follow, len(follows))
	for i, follow := range follows {
		results[i] = dbToFollow(follow)
	}
	return results
}

// FEEEEEEDS

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FetchedAt time.Time `json:"last_fetch"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserId    uuid.UUID `json:"user_id"`
}

func dbToFeed(feed database.Feed) Feed {

	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		FetchedAt: feed.LastFetch.Time,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserId:    feed.UserID,
	}
}
func FeedToFeeds(feeds []database.Feed) []Feed {
	results := make([]Feed, len(feeds))
	for i, feed := range feeds {
		results[i] = dbToFeed(feed)
	}
	return results
}

// USERS

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func dbToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}

}
