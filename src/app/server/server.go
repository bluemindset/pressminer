package server

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
