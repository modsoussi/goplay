package main

import "gopkg.in/mgo.v2/bson"
import "fmt"

func main() {
	id := bson.NewObjectId()
	fmt.Println(id)
	fmt.Println(id.Hex())
	fmt.Println(bson.ObjectIdHex(id.Hex()))

	id = bson.NewObjectId()
	fmt.Println(id)
	fmt.Println(id.Hex())
	fmt.Println(bson.ObjectIdHex(id.Hex()))
}
