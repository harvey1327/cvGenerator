package model

type social struct {
	Links []link
}

type link struct {
	Network string
	User    string
	Url     string
}
