package helper

import "fmt"

func ErrorPrint(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}