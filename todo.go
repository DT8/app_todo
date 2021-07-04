package todo

type ToDo struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Desctiption string `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Desctiption string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
