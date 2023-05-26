package main

import "golang/project"

func main() {
	//variables.Demo1()
	//fmt.Print()
	//conditionals.Demo3()
	//loops.Demo5()
	//slices.Demo2()
	//functions.SelamVer("Enes")
	//functions.SelamVer("Derin")
	//var sonuc = functions.Topla(2, 6)
	//fmt.Println(sonuc)

	// sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(10, 6)

	// fmt.Println("Toplam :", sonuc1)
	// fmt.Println("Cikarim :", sonuc2)
	// fmt.Println("Carpim", sonuc3)
	// //fmt.Println("Bolum", sonuc4)

	// fmt.Println(functions.ToplaVariadic(1, 4, 6, 3, 10))
	// fmt.Println(functions.ToplaVariadic(1, 4))
	// fmt.Println(functions.ToplaVariadic())

	// sayilar := []int{4, 6, 7, 0, 11}
	// fmt.Println(functions.ToplaVariadic(sayilar...))

	// sayi := 20
	// pointers.Demo1(&sayi)
	// fmt.Println("Mainde ki sayi: ", sayi)

	// sayilar := []int{1, 2, 3}
	// pointers.Demo2(sayilar)
	// fmt.Println("Mainde ki sayilar :", sayilar[0])

	//structs.Demo2()
	// go goroutines.CiftSayilar()
	// go goroutines.TekSayilar()

	// ciftSayiCn := make(chan int)
	// tekSayiCn := make(chan int)
	// go channels.CiftSayilar(ciftSayiCn)
	// go channels.TekSayilar(tekSayiCn)

	// ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn

	// carpim := ciftSayiToplam * tekSayiToplam
	// fmt.Println("Çarpım : ", carpim)

	//interfaces.Demo2()
	//error_handling.Demo1()
	//interfaces.Demo3()

	//fmt.Println(error_handling.TahminEt2(101))

	//string_functions.Demo2()

	//restful.Demo2()

	//project.GetAllProducts()

	project.AddProduct()
	project.GetAllProducts()
}
