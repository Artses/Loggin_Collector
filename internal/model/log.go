package model

import "time"

type LogItem struct {
	Order     int       `json:"order"`
	TimeStamp time.Time `json:"timestamp"`
	Line      string    `json:"line"`
}

type Log struct {
	Content []LogItem `json:"content"`
}
