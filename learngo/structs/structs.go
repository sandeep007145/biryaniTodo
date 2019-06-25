package structs

type Biryani struct {
	// ID      bson.ObjectId `bson:"_id" json:"id"`
	Rice    string `bson:"rice" json:"rice"`
	Chicken string `bson:"chicken" json:"chicken"`
	Masala  string `bson:"masala" json:"masala"`
}
