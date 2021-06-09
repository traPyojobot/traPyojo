package domain

import (
	"net/http"
	"time"
)

type Tweet struct {
	ID        int64
	Response  *http.Response
	Content   string
	CreatedAt time.Time
}

type Monologue struct {
	Content string
}
