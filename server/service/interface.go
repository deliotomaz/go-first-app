package service

import (
	"../domain"
)

type UserAppService interface {
	GetCount() (int, error)
	FindUserById(id int64) (domain.User, error)
	ListUsers() (domain.Users, error)
}
type WidgetAppService interface {
	GetCount() (int, error)
	FindWidgetById(id int64) (domain.Widget, error)
	ListWidgets() (domain.Widgets, error)
	AddUpdate(widget domain.Widget) error
}
