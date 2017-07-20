package main

import (
	"fmt"
	"labix.org/v2/mgo/bson"
)

func main() {
	s := ""
	fmt.Scanf("%s", s)
	fmt.Printf(s)
	fmt.Println(bson.IsObjectIdHex(s))
	return
}
