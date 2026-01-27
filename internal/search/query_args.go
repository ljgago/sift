package search

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	DefaultPage = 1
	DefaultSize = 20
	DefaultFrom = 0
)

type QueryArgs struct {
	Text        string
	Page        int
	Size        int
	From        int
	Quality     float64
	Popularity  float64
	Maintenance float64
}

func ParseQuery(query string) (*QueryArgs, error) {
	params, err := url.ParseQuery(query)
	if err != nil {
		return nil, fmt.Errorf("ParseQuery -> %w", err)
	}

	text := getText(params)
	size := getSize(params)
	from := getFrom(params)
	page := getPage(params)

	return &QueryArgs{
		Text: text,
		Size: size,
		From: from,
		Page: page,
	}, nil
}

func getText(m url.Values) string {
	return m.Get("q")
}

func getSize(m url.Values) int {
	v := m.Get("size")
	if v == "" || v == "0" {
		return DefaultSize
	}

	size, err := strconv.Atoi(v)
	if err != nil {
		return DefaultSize
	}

	return size
}

func getPage(m url.Values) int {
	v := m.Get("page")
	if v == "" || v == "0" {
		return DefaultPage
	}

	page, err := strconv.Atoi(v)
	if err != nil {
		return DefaultPage
	}

	return page
}

func getFrom(m url.Values) int {
	page := getPage(m)
	size := getSize(m)
	return (page - 1) * size
}
