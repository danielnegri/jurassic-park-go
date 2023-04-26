package model

import "strings"

type ID string

func NewID(prefix string, uuid string) ID {
	return ID(strings.Join([]string{prefix, uuid}, "_"))
}
