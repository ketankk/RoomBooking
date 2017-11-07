package main

import (
	"fmt"
	"database/sql/driver"
)

func main() {
	fmt.Println("Hello GoLang")
	var stmt=driver.Stmt("")
	fmt.Print(stmt)

}
