package main

import (
	"fmt"
	"github.com/okatotomato/gobot/models/user"
)

func main() {
	fmt.Println("GoBot Dziolcha v.01")
	u := user.User{}
	fmt.Println(u.GetTableName())
}
