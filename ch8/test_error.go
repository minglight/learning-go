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

func GenerateErrorAsPointer(flag bool) error {
	var err *StatusErr
	if flag {
		return StatusErr{
			status: NotFound,
		}
	}
	return err
}

func GenerateStructError(flag bool) *StatusErr {
	var err *StatusErr
	if flag {
		return &StatusErr{
			status: NotFound,
		}
	}
	return err
}
func main() {
	fmt.Println("======= GenerateError =======")
	fmt.Println("Return type is error")
	err := GenerateError(true)
	fmt.Println(err != nil)
	err = GenerateError(false)
	fmt.Println(err != nil)
	err2 := GenerateError(false)
	fmt.Println(err2 != nil)
	fmt.Println("======= GenerateErrorAsPointer =======")
	fmt.Println("Return type is error")
	err3 := GenerateErrorAsPointer(true)
	fmt.Println(err3 != nil)
	err4 := GenerateErrorAsPointer(false)
	fmt.Println(err4 != nil)
	err5 := GenerateErrorAsPointer(false)
	fmt.Println(err5 != nil)
	fmt.Println("======= GenerateStructError =======")
	fmt.Println("Return type is StatusErr")
	err6 := GenerateStructError(true)
	fmt.Println(err6 != nil)
	err7 := GenerateStructError(false)
	fmt.Println(err7 != nil)
	err8 := GenerateStructError(false)
	fmt.Println(err8 != nil)
}
