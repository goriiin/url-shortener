package main

import (
	"fmt"
	"github.com/goriiin/myapp/backend/internal/repository/postgres"
)

func main() {
	a := postgres.New()
	str, err := a.GetURL("123")
	if err != nil {
		fmt.Printf("err - %s", err)
	}
	fmt.Println("URL:", str)
}
