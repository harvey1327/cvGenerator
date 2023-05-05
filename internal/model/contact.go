package model

import (
	"strings"
)

type contact struct {
	Phone        string
	Email        string
	EmailDisplay string
}

func (c *contact) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type raw contact
	value := (*raw)(c)
	if err := unmarshal(value); err != nil {
		return err
	}
	if strings.Contains(value.Email, "_") {
		c.EmailDisplay = strings.ReplaceAll(value.Email, "_", "\\_")
	} else {
		c.EmailDisplay = value.Email
	}
	return nil
}
