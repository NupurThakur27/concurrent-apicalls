package model

type Data struct {
	Meta  MetaData `json:"meta"`
	Users []User   `json:"data"`
}
type MetaData struct {
	Pagination interface{}
}

// type Pagination struct {
// 	Total int `json:"total"`
// 	Pages int `json:"pages"`
// 	Page int `json:"page"`
// 	Limit int `json:"limit"`
// 	Links inter
// }
type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}
