package database

import (
	"time"
)

type model struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Board struct {
	model
	Items []Item `json:"items"`
	Users []User `gorm:"many2many:board_users;" json:"users"`
}

type Item struct {
	model
	RepoName        string `json:"repoName"`
	RepoURL         string `json:"repoURL"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	ConnectedBranch string `json:"connectedBranch"`
	List            uint   `json:"list"`
	Creator         User   `gorm:"many2many:item_user;" json:"creator"`
	Assignee        User   `gorm:"many2many:item_user;" json:"assignee"`
}

type User struct {
	model
	DisplayName    string `json:"displayName"`
	GithubUsername string `json:"githubUsername"`
}
