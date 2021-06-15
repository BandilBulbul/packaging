package main

import (
	"fmt"

	"github.com/BandilBulbul/packaging/calculation"
)

func main() {
	fmt.Println("Hello User")
	fmt.Print(calculation.Add(1, 2, 3, 4, 5))
}

// open powershell as a adminstration :Start-Process powershell -Verb runAs
