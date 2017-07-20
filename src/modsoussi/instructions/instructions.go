package main

import (
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var (
	DB         = "test"
	collection = "instructions"
	IsDrop     = true
)

type Info struct {
	Instructions map[string][]string `bson:"instructions" json:"instructions"`
}

func main() {
	fmt.Println("-------- Starting MongoDB games --------")

	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	// drop databsae
	if IsDrop {
		err = sess.DB(DB).DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	// use peeps collection
	c := sess.DB(DB).C(collection)

	// insert instructions
	m := make(map[string][]string)
	a := make([]string, 0)
	a = append(a, "First, ")
	a = append(a, "Second, ")
	a = append(a, "Finally, ")
	m["checkin"] = a

	err = c.Insert(&Info{Instructions: m})
	if err != nil {
		fmt.Println(err)
	}

	var s Info
	err = c.Find(bson.M{}).One(&s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
}
