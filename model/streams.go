package model

import "time"

type Stream struct {
	Labels  string  `json:"labels"`
	Entries []Entry `json:"entries"`
}

// Entry is a log entry with a timestamp.
type Entry struct {
	Timestamp time.Time `json:"ts"`
	Line      string    `json:"line"`
}

type PushRequest struct {
	Streams []Stream `json:"streams"`
}

type Label struct {
	Key, Value string
}

type LabelRules struct {
	Label, Cond, Value string
}
