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
			fmt.Print(pErr)
		}
		fmt.Println("Dosya bulunamadi")
	} else {
		fmt.Println(f.Name())
	}
}
