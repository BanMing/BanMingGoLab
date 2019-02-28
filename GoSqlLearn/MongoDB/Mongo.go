package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Per struct {
	Name string
	Age  int
}

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")
	err = c.Insert(&Per{"sdf", 34}, &Per{"Cla", 90})
	if err != nil {
		fmt.Println(err)
	}
	result := Per{}
	err = c.Find(bson.M{"name": "sdf"}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("age:", result.Age)
}
