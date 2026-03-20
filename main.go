package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("some realization")
	}
	fmt.Println("start git/branch/dev1")
	for n := 20; n < 30; n++ {
		fmt.Println("some real for dev1")
		fmt.Println("there is some new realization")
		fmt.Println("start branch origin/dev2")
		for n := 10; n < 20; n++ {
			fmt.Println("new feature for dev2")
		}
	}
	fmt.Println("hello")
}
