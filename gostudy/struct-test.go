package github

import "time"

const IssuesURL = "https://api.gitub.com/search/issues"

type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"created_at"`
	Body     string    //Markdown 格式
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
