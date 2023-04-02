package forrange

import "fmt"

func Demo1() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir"}

	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}

	for i, sehir := range sehirler {
		fmt.Println(i, " ", sehir)
	}
}
