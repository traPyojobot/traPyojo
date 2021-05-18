package domain

import (
	"time"

	"github.com/gofrs/uuid"
)

type Tweet struct {
	ID        uuid.UUID
	Content   string
	CreatedAt time.Time
}

type Monologue struct {
	Content string
}
