package main

import "fmt"

func main() {
	pinPan()
}

func divisivelPorTres() {
	for i := 3; i < 100; i += 3 {
		fmt.Println(i)
	}

}

func pinPan() {
	for i := 3; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Pin,Pan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		}
	}

}
