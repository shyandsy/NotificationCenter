package server

type Server interface {
	Start(addr string)
	Notify()
}
