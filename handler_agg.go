package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/ProbsPropps/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[1])
	if err != nil {
		return fmt.Errorf("Error - handlerAgg: %v\n", err)
	}
	
	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)	
	
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			return err
		}
	}

}

func scrapeFeeds(s *state) error {
	nextFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Error - scrapeFeeds: %v\n", err)
	}
	
	err = s.db.MarkFeedFetched(context.Background(), nextFetch.ID)
	if err != nil {
		return fmt.Errorf("Error - scrapeFeeds: %v\n", err)
		}
	
	rssFeed, err := fetchFeed(context.Background(), nextFetch.Url)
	if err != nil {
		return fmt.Errorf("Error - scrapeFeeds: %v\n", err)
	}
	
	for _, item := range rssFeed.Channel.Item {
		fmtTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return fmt.Errorf("Error  scrapeFeeds %v\n", err)
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Url: item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: fmtTime,
			FeedID: nextFetch.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			fmt.Printf("Error - scrapeFeeds: %v\n", err)
			continue
		}
	}
	return nil
}

