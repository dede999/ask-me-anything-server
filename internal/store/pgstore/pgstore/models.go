// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package pgstore

import (
	"github.com/google/uuid"
)

type Answer struct {
	ID        uuid.UUID
	MessageID uuid.UUID
	Answer    string
}

type Message struct {
	ID        uuid.UUID
	RoomID    uuid.UUID
	Message   string
	Reactions int64
	Answered  bool
}

type Room struct {
	ID    uuid.UUID
	Theme string
}
