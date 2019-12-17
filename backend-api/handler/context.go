package handler

// Context - context interface
type Context interface {
	Param(name string) string
	Bind(i interface{}) error
	//Status(int)
	JSON(code int, i interface{}) error
}
