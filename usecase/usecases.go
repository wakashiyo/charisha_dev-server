package usecase

import (
	"github.com/wakashiyo/charisha_dev-server/domain"
)

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
