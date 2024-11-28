package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
	"golang.org/x/net/http2"
)

type CalendarClient struct {
	url    string
	client *http.Client
}

func NewCalendarClient(url string) *CalendarClient {
	return &CalendarClient{
		url:    url,
		client: &http.Client{Transport: &http2.Transport{}},
	}
}

func (c *CalendarClient) Get() (*html.Node, error) {
	response, err := c.client.Get(c.url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", response.StatusCode)
	}

	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
