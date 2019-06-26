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

func RemakeBiryani(db *mgo.Database, data interface{}, id string) error {
	fmt.Println("inside update")
	Id := bson.ObjectIdHex(id)
	return db.C("biryani").Update(bson.M{"_id": Id}, data)
}

func TakeBiryani(db *mgo.Database, page int, limit int) (int, []bson.M, error) {
	fmt.Println("Inside find")
	var query interface{}
	var res []bson.M
	total, _ := db.C("biryani").Count()
	err := db.C("biryani").Find(query).
		Skip((page - 1) * limit).Limit(limit).All(&res)
	return total, res, err
}

func GetSinglePack(db *mgo.Database, id string) ([]bson.M, error) {
	fmt.Println("inside single")
	Id := bson.ObjectIdHex(id)
	var res []bson.M
	err := db.C("biryani").Find(bson.M{"_id": Id}).All(&res)
	return res, err
}

func RemoveBiryani(db *mgo.Database, id string) error {
	Id := bson.ObjectIdHex(id)
	err := db.C("biryani").Remove(bson.M{"_id": Id})
	fmt.Println(Id)
	return err
}
