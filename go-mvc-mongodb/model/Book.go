package model

type Book struct {
	// add ID and tags if you need them
	// ID     bson.ObjectId // `json:"id" bson:"_id"`
	Isbn   string  // `json:"isbn" bson:"isbn"`
	Title  string  // `json:"title" bson:"title"`
	Author string  // `json:"author" bson:"author"`
	Price  float32 // `json:"price" bson:"price"`
}