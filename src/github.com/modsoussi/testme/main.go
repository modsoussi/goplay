package testme

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	a := bson.M{}
	a["1"] = "Hello"

	fmt.Println(a)
}
