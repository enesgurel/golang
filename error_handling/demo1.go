package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt")
	//nil
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Print("Dosya bulunamadi : ", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadi")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}
