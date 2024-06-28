package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"

	"github.com/gpr3211/blogger/internal/clog"
)

type RSSfeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urltoFeed(url string) (*RSSfeed, error) {
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}
	rsp, err := httpClient.Get(url)
	if err != nil {
		clog.Println("client failed to fetch url")
	}
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}
	dat, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	rssFeed := RSSfeed{}
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return nil, err
	}
	return &rssFeed, nil
}
