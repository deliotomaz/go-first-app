package domain

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Gravatar string `json:"gravatar"`
}

type Users []User

type Widget struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	Price     string `json:"price"`
	Inventory int    `json:"inventory"`
	Melts     bool   `json:"melts"`
}
type Widgets []Widget

type UserRepository interface {
	Count() (int, error)
	FindById(id int64) (User, error)
	List() (Users, error)
}
type WidgetRepository interface {
	Count() (int, error)
	FindById(id int64) (Widget, error)
	List() (Widgets, error)
	AddUpdate(widget Widget) error
}
