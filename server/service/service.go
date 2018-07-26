package service

import (
	"../domain"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserAppService(UserRepository domain.UserRepository) *UserService {
	return &UserService{
		repo: UserRepository,
	}
}
func (userService *UserService) GetCount() (int, error) {
	return userService.repo.Count()
}

// func (userService *UserService) ListTopNameByName(string name) (*[]string, error) {

// 	var items = userService.repo.SearchByName

// }
