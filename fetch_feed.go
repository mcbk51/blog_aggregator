package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"

)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Items       []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// Create HTTP request with context
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err 
	}
	
	req.Header.Set("User-Agent", "gator")
	
	// Create HTTP client and make the request
	client := &http.Client{
		Timeout: 10* time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err 
	}
	defer resp.Body.Close()
	
	// Read the entire response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err 
	}
	
	// Parse the XML into our RSSFeed struct
	var feed RSSFeed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, err 
	}
	
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i := range feed.Channel.Items {
		feed.Channel.Items[i].Title = html.UnescapeString(feed.Channel.Items[i].Title)
		feed.Channel.Items[i].Description = html.UnescapeString(feed.Channel.Items[i].Description)
	}
	
	return &feed, nil
}
