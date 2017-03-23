package main

import (
	"encoding/xml"
	"time"
)

type item struct {
	XMLName     xml.Name  `xml:"item"`
	Title       string    `xml:"title"`
	Description string    `xml:"description"`
	PubDate     time.Time `xml:"pubDate"`
	Link        string    `xml:"link"`
	Guid        string    `xml:"guid"`
	Content     string    `xml:"content:encoded"`
}

type image struct {
	XMLName xml.Name `xml:"image"`
	URL     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
}

type channel struct {
	XMLName       xml.Name  `xml:"channel"`
	Title         string    `xml:"title"`
	ink           string    `xml:"link"`
	Description   string    `xml:"description"`
	Language      string    `xml:"language"`
	Copyright     string    `xml:"copyright"`
	Generator     string    `xml:"generator"`
	LastBuildDate time.Time `xml:"lastBuildDate"`
	Image         image     `xml:"image"`
	Items         []item    `xml:"item"`
}

type rssDocument struct {
	XMLName xml.Name `xml:"rss"`
	Channel channel  `xml:"channel"`
}
