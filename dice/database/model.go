package database

import (
	"gorm.io/gorm"
)

type List int

const (
	BACKLOG = iota
	TODO
	DOING
	REVIEW
	DONE
)

type Item struct {
	gorm.Model
	RepoID			uint
	Repo            Repo
	Title           string
	Description     string
	CreatorID		uint
	Creator         User
	AssigneeID 		uint
	Assignee        User
	ConnectedBranch string
	List            List
}

type User struct {
	gorm.Model
	Name string
}

type Repo struct {
	gorm.Model
	Name string
	URL  string
}
