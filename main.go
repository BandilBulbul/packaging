//$ cat hello.go

package main

import (
	"fmt"

	"github.com/BandilBulbul/packaging/calculation"
)

func main() {
	fmt.Println("Hello User")
	fmt.Print(calculation.Add(1, 2, 3, 4, 5))
}

//https://www.youtube.com/watch?v=S79DBpCq3HQ

// open powershell as a adminstration :Start-Process powershell -Verb runAs
//$ GOOS=windows GOARCH=386 go build -o hello.exe hello.go
