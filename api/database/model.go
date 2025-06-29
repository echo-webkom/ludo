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

type Board struct {
	Model
	Items []Item `gorm:"foreignKey:BoardID" json:"items"`
	Users []User `gorm:"many2many:board_users;" json:"users"`
}

type Item struct {
	Model
	BoardID         uint   `json:"boardId"`
	RepoName        string `json:"repoName"`
	RepoURL         string `json:"repoURL"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	ConnectedBranch string `json:"connectedBranch"`
	List            uint   `json:"list"`

	CreatorID uint `json:"-"`
	Creator   User `gorm:"foreignKey:CreatorID" json:"creator"`

	AssigneeID uint `json:"-"`
	Assignee   User `gorm:"foreignKey:AssigneeID" json:"assignee"`
}

type User struct {
	Model
	DisplayName    string `json:"displayName"`
	GithubUsername string `json:"githubUsername"`
}
