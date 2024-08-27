package model

type Task struct {
	Id     string `json:"id"`
	UserID string `json:"userId"`
	Title  string `json:"title"`
}