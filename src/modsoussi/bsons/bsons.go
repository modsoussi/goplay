package main

import "gopkg.in/mgo.v2/bson"
import "fmt"

func main() {
	a := map[string]string{"hotelName": "The Modsoussi Hotel", "toRemove": "HIIIII!"}
	fmt.Println(a)

	b := bson.M{}

	for k, v := range a {
		b[k] = v
	}
	fmt.Println(b)

	delete(b, "toRemove")
	fmt.Println(b)
}
