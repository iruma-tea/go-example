package main

import "fmt"

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusError struct {
	Status  Status
	Message string
}

func (se StatusError) Error() string {
	return se.Message
}

func GenerateError(flag bool) error {
	var genErr StatusError
	if flag {
		genErr = StatusError{
			Status: NotFound,
		}
	}
	return genErr
}

func main() {
	err := GenerateError(true)
	fmt.Println(err != nil)
	err = GenerateError(false)
	fmt.Println(err != nil)
}
