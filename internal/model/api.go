package model

type ApiStoreInput struct {
	Menu   string
	Name   string
	Key    string
	Path   string
	Method int
	Type   int
	Pid    int
}

type ApiUpdateInput struct {
	Id     int
	Menu   string
	Name   string
	Key    string
	Path   string
	Method int
	Type   int
	Pid    int
}
