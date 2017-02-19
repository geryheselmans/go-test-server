package models

import "net/mail"

type Author struct {
	id        uint64
	firstName string
	lastName  string
	mail      mail.Address
}
