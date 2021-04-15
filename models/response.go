package models

type Response struct {
	Code    int
	Body    interface{}
	Title   string
	Message string
}
