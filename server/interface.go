package server

type Server interface{
	Start() error
SetHandler(func(int, int) int) int
}
