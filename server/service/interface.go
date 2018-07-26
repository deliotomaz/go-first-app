package service

type UserAppService interface {
	GetCount() (int, error)
	//	ListTopNameByName(name string) (*[]string, error)
}
type WidgetAppService interface {
	GetCount() (int, error)
	//	ListTopNameByName(name string) (*[]string, error)
}
