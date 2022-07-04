package link

import "time"

type Link struct {
	Id            string
	Url           string
	CreatedAt     time.Time
	RedirectCount int
}
