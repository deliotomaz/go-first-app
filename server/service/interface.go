package service

import (
	"../domain"
)

//UserAppService ...define all contracts for comunicate controllers and repository
type UserAppService interface {
	GetCount() (int, error)
	FindUserById(id int64) (domain.User, error)
	ListUsers() (domain.Users, error)
}

//WidgetAppService ...define all contracts for comunicate controllers and repository
type WidgetAppService interface {
	GetCount() (int, error)
	FindWidgetById(id int64) (domain.Widget, error)
	ListWidgets() (domain.Widgets, error)
	AddUpdate(widget domain.Widget) error
}
