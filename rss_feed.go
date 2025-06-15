package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"

	"github.com/ProbsPropps/gator/internal/database"
)

type RSSFeed struct {
	Channel struct {
		Title 		string 		`xml:"title"`
		Link 		string 		`xml:"link"`
		Description string		`xml:"description"`
		Item 		[]RSSItem 	`xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title 		string `xml:"title"`
	Link 		string `xml:"link"`
	Description	string `xml:"description"`
	PubDate 	string `xml:"pubDate"`
}


func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Error - fetchFeed: %v\n", err)
	}
	
	req.Header.Set("User-Agent", "gator")

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Error - fetchFeed: %v\n", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Error - fetchFeed: %v\n", err)
	}

	var feed RSSFeed
	if err = xml.Unmarshal(data, &feed); err != nil {
		return &RSSFeed{}, fmt.Errorf("Error - fetchFeed: %v\n", err)
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	
	for i, item := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		feed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}

	return &feed, nil
	
}

func printFeed(feed database.Feed) {
	fmt.Println("Feed Fields:")
	fmt.Printf("  -ID: %v\n", feed.ID)
	fmt.Printf("  -CreatedAt: %v\n", feed.CreatedAt)
	fmt.Printf("  -UpdatedAt: %v\n", feed.UpdatedAt)
	fmt.Printf("  -Name: %s\n", feed.Name)
	fmt.Printf("  -URL: %s\n", feed.Url)
	fmt.Printf("  -UserID: %v\n", feed.UserID)
}
