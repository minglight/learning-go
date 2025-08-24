package main

import "fmt"

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	status  Status
	Message string
}

func (s StatusErr) Error() string {
	return s.Message
}

func GenerateError(flag bool) error {

	if flag {
		return StatusErr{
			status: NotFound,
		}
	}
	return nil
}

func main() {
	err := GenerateError(true)
	fmt.Println(err != nil)
	err = GenerateError(false)
	fmt.Println(err != nil)
	err2 := GenerateError(false)
	fmt.Println(err2 != nil)
}
