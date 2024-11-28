package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func parseCalendarHTML(doc *html.Node) (Calendar, error) {
	calendar := Calendar{ConditionsBySign: make(map[string]ConditionByDay)}

	var table *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "table" {
			table = n
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if table == nil {
		return calendar, fmt.Errorf("table not found")
	}

	for tr := table.FirstChild; tr != nil; tr = tr.NextSibling {
		// TODO
	}

	return calendar, nil
}
