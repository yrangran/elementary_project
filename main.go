package main

import (
	"elementary_project/struct"
	"fmt"
)

func main(){
	f := Website.Website{
		URL: "https://example.com",
		Status: true,
		LastChecked: 1234567,
	}
	fmt.Println(f)
}