package repo

import (
	"errors"

	model "github.com/aviabird/go-echo-seed/app/models"
	"gorm.io/gorm"
)

func (s *Repo) ListUsers() ([]model.User, error) {
	var users = []model.User{}
	if err := s.db.Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return users, nil
}
