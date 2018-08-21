package model

type Message struct {
	From        string
	FromName    string
	To          []string
	Subject     string
	ContentType string
	Content     string
	Attach      []string
}
