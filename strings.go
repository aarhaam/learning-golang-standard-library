package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("Mahardika", "dika"))
	fmt.Println(strings.Split("Eko Kurniawan", " "))
	fmt.Println(strings.ToLower("MAHARDIKA"))
	fmt.Println(strings.Trim("		Mahardika		", "		"))

}
