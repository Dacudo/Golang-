package main

import (
	"fmt"
	"os"
	"strconv"

	"Visual/github.com/Dacudo/cmp"
)

const (
	errMsg1 = "First value inputed is invalid"
	errMsg2 = "Second value inputed is invalid"
	errMsg3 = "Third value inputed is invalid"
)

func main() {

	if len(os.Args) != 4 {

		fmt.Println("Error! There should be three utility values : ")

		return
	}

	v1, err1 := strconv.ParseFloat(os.Args[1], 8)
	v2, err2 := strconv.ParseFloat(os.Args[2], 8)
	v3, err3 := strconv.ParseFloat(os.Args[3], 8)

	if err1 != nil {

		fmt.Println(errMsg1)

		return
	}

	if err2 != nil {

		fmt.Println(errMsg2)

		return
	}

	if err3 != nil {

		fmt.Println(errMsg3)

		return
	}

	fmt.Println("The variables are as follows:", v1, v2, v3)

	fmt.Println("Now then, lets get started shall we ?")

	vO := cmp.Utilities(v1, v2, v3)

	fmt.Printf("The amounts owed per person is : %.2f\n", vO)

}
