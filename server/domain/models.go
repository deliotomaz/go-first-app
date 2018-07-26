package domain

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Gravatar string `json:"gravatar"`
}

type Users []User

type Widget struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Inventory int64   `json:"inventory"`
	Melts     bool    `json:"melts"`
}
type Widgets []Widget

// type BaseReporsitory interface {
// 	Count() error
// 	FindById(id int64)
// 	Save(item struct{})
// 	List(filter struct{}, page int, pageSize int)
// }

type UserRepository interface {
	Count() (int, error)
	// FindById(id int64) (*User, error)
	// Save(item struct{}) (*User, error)
	// List(page int, pageSize int)
	// SearchByName(name string) (*Users, error)
}
type WidgetsRepository interface {
	Count() (int, error)
	FindById(id int64) (*Widget, error)
	Save(item struct{}) (*Widget, error)
	SearchByName(name string) (*Widgets, error)
	List(page int, pageSize int)
}
