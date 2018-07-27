package service

import (
	"../domain"
)

//UserService ...Concrete implment.  for UserAppService
type UserService struct {
	repo domain.UserRepository
}

//WidgetService ...Concrete implment.  for WidgetAppService
type WidgetService struct {
	repo domain.WidgetRepository
}

func NewUserAppService(UserRepository domain.UserRepository) *UserService {
	return &UserService{
		repo: UserRepository,
	}
}
func (userService *UserService) GetCount() (int, error) {
	return userService.repo.Count()
}
func (userService *UserService) FindUserById(id int64) (domain.User, error) {
	return userService.repo.FindById(id)
}
func (userService *UserService) ListUsers() (domain.Users, error) {
	return userService.repo.List()
}

func NewWidgetAppService(WidgetRepository domain.WidgetRepository) *WidgetService {
	return &WidgetService{
		repo: WidgetRepository,
	}
}
func (widgetService *WidgetService) GetCount() (int, error) {
	return widgetService.repo.Count()
}
func (widgetService *WidgetService) FindWidgetById(id int64) (domain.Widget, error) {
	return widgetService.repo.FindById(id)
}
func (widgetService *WidgetService) ListWidgets() (domain.Widgets, error) {
	return widgetService.repo.List()
}
func (widgetService *WidgetService) AddUpdate(widget domain.Widget) error {
	return widgetService.repo.AddUpdate(widget)
}
