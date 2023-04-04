package defer_statement

import "fmt"

func A() {
	fmt.Println("a fonksiyonu çalıştı")
}

func B() {
	defer A() // b çalıştıktan sonra çalışmaktadır.
	defer C()
	defer D()
	fmt.Println("b fonksiyonu çalıştı")
}

func C() {
	fmt.Println("c fonksiyonu çalıştı")
}

func D() {
	fmt.Println("d fonksiyonu çalıştı")
}
