package service

type Student struct {
	Id      int64
	Name    string
	Subject []string
	Address Address
	Point   int `Maximum cant over 99`
}

type Address struct {
	Location string
}
