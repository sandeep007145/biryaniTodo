package biryani

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../mongodbconnect"
	"../structs"
	"gopkg.in/mgo.v2"
)

func TakeCheckenBiryani(w http.ResponseWriter, r *http.Request) {
	db := getCon()
	query := r.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	page, _ := strconv.Atoi(query.Get("page"))
	fmt.Println("Limit:", limit, page)
	total, res, err := mongodbconnect.TakeBiryani(db, page, limit)
	if err != nil {
		log.Print(err)
	}
	response := make(map[string]interface{})
	response["data"] = res
	response["count"] = total

	resByte, _ := json.Marshal(response)
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

func RemakeSingleBiryani(w http.ResponseWriter, r *http.Request) {
	db := getCon()
	id := r.URL.Query().Get("id")
	data := structs.Biryani{}
	err := json.NewDecoder(r.Body).Decode(&data)
	fmt.Println(err, data)
	errr := mongodbconnect.RemakeBiryani(db, data, id)
	fmt.Println(errr)
}

func RemoveCheckenBiryani(w http.ResponseWriter, r *http.Request) {
	db := getCon()
	id := r.URL.Query().Get("id")
	errr := mongodbconnect.RemoveBiryani(db, id)
	fmt.Println(errr, "delete")
}

func TakeSingleBiryani(w http.ResponseWriter, r *http.Request) {
	db := getCon()
	id := r.URL.Query().Get("id")
	res, err := mongodbconnect.GetSinglePack(db, id)
	if err != nil {
		log.Panic(err)
	}
	resByte, _ := json.Marshal(res)
	w.Write(resByte)
}

func getCon() *mgo.Database {
	return mongodbconnect.Adapter()
}
