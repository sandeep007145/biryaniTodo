package biryani

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../mongodbconnect"
	"../structs"
	"gopkg.in/mgo.v2"
)

func TakeCheckenBiryani(w http.ResponseWriter, r *http.Request) {
	db := getCon()
	res, err := mongodbconnect.TakeBiryani(db)
	if err != nil {
		log.Print(err)
	}
	resByte, _ := json.Marshal(res)
	w.Write(resByte)
}

func MakeCheckenBiryani(w http.ResponseWriter, r *http.Request) {
	db := getCon()
	data := structs.Biryani{}
	err := json.NewDecoder(r.Body).Decode(&data)
	fmt.Println(err, data)
	errr := mongodbconnect.InsertBiryani(db, data)
	fmt.Println("Error:", errr)
}

func getCon() *mgo.Database {
	return mongodbconnect.Adapter()
}
