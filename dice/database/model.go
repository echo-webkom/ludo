package database

import "time"

type List int

const (
	BACKLOG = iota
	TODO
	DOING
	REVIEW
	DONE
)

type Item struct {
	ID              uint
	Repo            Repo
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Title           string
	Description     string
	Creator         User
	Assignee        User
	ConnectedBranch string
	List            List
}

type User struct {
	ID   uint
	Name string
}

type Repo struct {
	ID   uint
	Name string
	URL  string
}
