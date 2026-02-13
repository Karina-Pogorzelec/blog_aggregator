package main

import (
	"context"
	"net/http"
	"io"
	"encoding/xml"
	"html"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	feed, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var feedData RSSFeed
	if err := xml.Unmarshal(feed, &feedData); err != nil {
		return nil, err
	}

	feedData.Channel.Title = html.UnescapeString(feedData.Channel.Title)
	feedData.Channel.Description = html.UnescapeString(feedData.Channel.Description)
	for i := range feedData.Channel.Item {
		feedData.Channel.Item[i].Title = html.UnescapeString(feedData.Channel.Item[i].Title)
		feedData.Channel.Item[i].Description = html.UnescapeString(feedData.Channel.Item[i].Description)
	}
	return &feedData, nil
}