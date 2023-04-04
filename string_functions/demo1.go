package string_functions

// alias
import (
	"fmt"
	s "strings"
)

// case sensitive
func Demo1() {
	isim := "Enes"
	fmt.Println(s.Count(isim, "e"))
	fmt.Println(s.Contains(isim, "e"))
	sonuc := s.Index(isim, "a")

	if sonuc != -1 {
		fmt.Println("a harfi var")
	} else {
		fmt.Println("a harfi yok")
	}

	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))
}
