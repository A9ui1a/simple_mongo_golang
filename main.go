package main

import (
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

type Toy struct {
	Name  string
	Price int
}

var c_info string = "127.0.0.1:27017"
var index_gundame string = "Gundam" + "2"

func main() {
	fmt.Println("=====Start connect mongoDB=====")
	s, err := mgo.Dial(c_info)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	fmt.Println("success connect")
	s.SetSafe(&mgo.Safe{})
	c := s.DB("ToysProduct").C("Toy")
	total, err := c.Find(bson.M{}).Count()
	fmt.Printf("Data Total => %d\n", total)

	if total == 0 {
		rand.Seed(time.Now().UnixNano())
		for i := 1; i <= 6; i++ {
			tmp_price := i * 10
			tmp := fmt.Sprintf("%s%d", "Gundam", i)
			_ = c.Insert(&Toy{tmp, tmp_price})
		}
		fmt.Printf("%s", "success insert data")

	}

	fmt.Println("=====Get Toys Data=====")
	res := Toy{}
	data := c.Find(bson.M{}).Iter()
	for data.Next(&res) {
		fmt.Printf("Name=>%s,Price=>%d\n", res.Name, res.Price)
	}

	fmt.Println("=====Get Toys Data Price >30=====")
	data = c.Find(bson.M{"price": bson.M{"$gt": 30}}).Iter()
	for data.Next(&res) {
		fmt.Printf("Name=>%s,Price=>%d\n", res.Name, res.Price)
	}

	fmt.Println("=====Update Toys=====")
	index_gundame = "Gundam" + "2"
	c.Update(bson.M{"Name": index_gundame}, bson.M{"Price": "10"})

	fmt.Println("=====Delete Toys=====")
	index_gundame = "Gundam" + "3"
	c.Remove(bson.M{"Name": index_gundame})

	fmt.Println("=====Get left Toys Data=====")
	data = c.Find(bson.M{}).Iter()
	for data.Next(&res) {
		fmt.Printf("Name=>%s,Price=>%d\n", res.Name, res.Price)
	}

	fmt.Println("=====clear Toys Data=====")
	c.Remove(bson.M{})

	fmt.Println("=====Drop Toys Data=====")
	c.Database.DropDatabase()

	fmt.Println("=====Finish=====")
}
