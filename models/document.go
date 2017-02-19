package models

import "time"

type Document struct {
	id      uint64
	author  Author
	date    time.Time
	title   string
	content string
}
