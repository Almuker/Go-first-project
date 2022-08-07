package services

import "fmt"

func ErrorsHandler(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}