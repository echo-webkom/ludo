package database

import (
	"time"

	"gorm.io/gorm"
)

// Copy of gorm.Model with JSON tags
type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Used in API response when creating objects.
type ID struct {
	ID uint `json:"id"`
}

type Board struct {
	Model
	Items []Item `gorm:"foreignKey:BoardID" json:"items"`
	Users []User `gorm:"many2many:board_users;" json:"users"`
}

type Status uint

const (
	StatusBacklog    Status = iota // Item has been created with little to no details
	StatusReady                    // Item is ready to be worked on and has sufficient details
	StatusInProgress               // Item is being worked on and tracks a branch
	StatusInReview                 // Item has a pull request open
	StatusClosed                   // Items pull request was merged or closed
)

type Item struct {
	Model
	BoardID         uint   `json:"boardId"`
	RepoName        string `json:"repoName"`
	RepoURL         string `json:"repoURL"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	ConnectedBranch string `json:"connectedBranch"`
	Status          Status `json:"status"`

	CreatorID uint `json:"-"`
	Creator   User `gorm:"foreignKey:CreatorID" json:"creator"`

	AssigneeID uint `json:"-"`
	Assignee   User `gorm:"foreignKey:AssigneeID" json:"assignee"`

	Data string `json:"data"`
}

type User struct {
	Model
	DisplayName    string `json:"displayName"`
	GithubUsername string `json:"githubUsername"`
}
