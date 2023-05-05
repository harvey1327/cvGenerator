package model

type skills struct {
	Sets []set
}

type set struct {
	Name   string
	Skills []string
}
