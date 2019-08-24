package domain

import "time"

type Tweet interface {
	Print() string
	GetUser() *User
	GetText() string
	GetID() int64
	GetDate() *time.Time
	SetID(id int64)
	SetText(text string)
}
