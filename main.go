package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	Name string
	Age  string
}

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	//拿到需要使用的集合
	collection := session.DB("test").C("person")

	//插入数据
	err = collection.Insert(&Person{"aaa", "1"},
		&Person{"bbb", "2"})
	if err != nil {
		log.Fatal(err)
	}

	//查询单个数据
	result := Person{}
	err = collection.Find(bson.M{"name": "aaa"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("age:", result.Age)
}
