package mongodbconnect

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

func Adapter() *mgo.Database {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	fmt.Println("connected with database")
	return session.DB("makeBiryani")
}

func InsertBiryani(db *mgo.Database, data interface{}) error {
	fmt.Println("Inside insert")
	return db.C("biryani").Insert(data)

}

func TakeBiryani(db *mgo.Database) ([]bson.M, error) {
	fmt.Println("Inside find")
	var query interface{}
	var res []bson.M
	err := db.C("biryani").Find(query).All(&res)
	return res, err
}
