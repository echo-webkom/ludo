package database

import (
	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
}

type Item struct {
	gorm.Model
	RepoName        string
	RepoURL         string
	Title           string
	Description     string
	Creator         User
	Assignee        User
	ConnectedBranch string
	List            uint
}

type User struct {
	gorm.Model
	DisplayName    string
	GithubUsername string
}
