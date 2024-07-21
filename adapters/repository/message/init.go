package message

import (
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

var repo *Repo

func GetRepo(db *gorm.DB) *Repo {
	if repo != nil {
		return repo
	}

	repo = &Repo{
		db: db,
	}

	return repo
}
