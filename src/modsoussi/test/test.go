package test

import (
	"fmt"

	"encoding/json"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var (
	DB         = "test"
	collection = "2017-hotel-occs"
	IsDrop     = true
)

type DayOccupancy struct {
	Day       int `bson:"day,omitempty" json:"day,omitempty"`
	Occupancy int `bson:"occupancy,omitempty" json:"occupancy,omitempty"`
}

type DayOccupancyContainer struct {
	Days []DayOccupancy `bson:"days" json:"days"`
}

type MonthOccupancy struct {
	Month int            `bson:"month,omitempty" json:"month,omitempty"`
	Days  []DayOccupancy `bson:"days,omitempty" json:"days,omitempty"`
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

	// Index
	index := mgo.Index{
		Key:        []string{"month"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	// insert data
	dateOccupancy := &DayOccupancy{Day: 25, Occupancy: 5}
	err = c.Insert(&MonthOccupancy{Month: 3, Days: []DayOccupancy{*dateOccupancy}})
	if err != nil {
		panic(err)
	}

	// add another occupancy
	err = c.Update(bson.M{"month": 3}, bson.M{"$push": bson.M{"days": bson.M{"day": 26, "occupancy": 10}}})
	if err != nil {
		panic(err)
	}

	// update occupancy for day 25 by setting it
	err = c.Update(bson.M{"month": 3, "days.day": 25}, bson.M{"$set": bson.M{"days.$.occupancy": 8}})
	if err != nil {
		panic(err)
	}

	// update occupancy for day 25 by incrementing it by 3
	err = c.Update(bson.M{"month": 3, "days.day": 25}, bson.M{"$inc": bson.M{"days.$.occupancy": 3}})
	if err != nil {
		panic(err)
	}

	// querying month occupancy
	var m MonthOccupancy
	err = c.Find(bson.M{"month": 3}).One(&m)
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(m)
	fmt.Println(string(b[:]))

	// querying day occupancy
	var d DayOccupancyContainer
	err = c.Find(bson.M{"month": 3}).Select(bson.M{"days": bson.M{"$elemMatch": bson.M{"day": 25}}}).One(&d)
	if err != nil {
		panic(err)
	}
	b, err = json.Marshal(d.Days[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b[:]))

	// attempt to update occupancy for day 18 by incrementing it by 3
	err = c.Update(bson.M{"month": 3, "days.day": 18}, bson.M{"$inc": bson.M{"days.$.occupancy": 3}})
	if err != nil {
		//panic(err) // returns not found
	}

	// attempt to query non-existing month
	count, err := c.Find(bson.M{"month": 4}).Count()
	if err != nil {
		panic(err)
	} else if count == 0 {
		err = c.Insert(&MonthOccupancy{Month: 4, Days: []DayOccupancy{DayOccupancy{Day: 17, Occupancy: 9}}})
		if err != nil {
			panic(err)
		}
	}

	// non-existent month in place, update its occupancy for day 7 of that month
	err = c.Update(bson.M{"month": 4}, bson.M{"$push": bson.M{"days": bson.M{"day": 7, "occupancy": 7}}})
	if err != nil {
		panic(err)
	}

	// query newly added month to make sure fields are correct
	err = c.Find(bson.M{"month": 4}).One(&m)
	if err != nil {
		panic(err)
	}
	b, err = json.Marshal(m)
	fmt.Println(string(b[:]))

	//querying a non-existent date
	var e DayOccupancyContainer
	err = c.Find(bson.M{"month": 3}).Select(bson.M{"days": bson.M{"$elemMatch": bson.M{"day": 18}}}).One(&e)
	if err != nil {
		panic(err)
	}
	if len(e.Days) > 0 {
		b, err = json.Marshal(e.Days[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b[:]))
	} else {
		fmt.Println("No such date.")
	}

	fmt.Println("-------- Ending MongoDB games --------")
}
