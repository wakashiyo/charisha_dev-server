package usecase

import (
	"github.com/wakashiyo/charisha_dev-server/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

// func (i *UserInteractor) Add(u domain.User) (user domain.User, err error) {
// 	id, err := i.UserRepository.Store(u)
// 	if err != nil {
// 		return
// 	}
// 	user, err = i.UserRepository.FindById(id)
// 	return
// }

func (i *UserInteractor) Add(u domain.User) (id int, err error) {
	id, err = i.UserRepository.Store(u)
	return
}

func (i *UserInteractor) Users() (u domain.Users, err error) {
	u, err = i.UserRepository.FindAll()
	return
}

func (i *UserInteractor) UserById(id int) (u domain.User, err error) {
	u, err = i.UserRepository.FindById(id)
	return
}
