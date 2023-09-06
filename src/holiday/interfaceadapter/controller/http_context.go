package controller

type httpContext interface {
	Param(string) string
	Bind(interface{}) error
	JSON(code int, i interface{}) error
}
