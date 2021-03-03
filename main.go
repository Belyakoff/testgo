package main

import (
	"context"
	"log"
	"time"
	"fmt"
	//"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)


func connect(dbname string)(*mongo.Collection){
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
 	defer cancel()


 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://darius:"+pswd+"@cluster0.aghsw.mongodb.net/"+dbname+"?retryWrites=true&w=majority",
  	))
	if err != nil { log.Fatal(err) }

	err = client.Ping(context.TODO(), nil)
	if err != nil {
    	log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database(dbname).Collection("apartments")

	return collection
}

func insert_document(collection *mongo.Collection, data Apartment){

	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
    	log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}


func find_document(collection *mongo.Collection)(Apartment){

	// create a value into which the result can be decoded
	var result Apartment
	filter := bson.M{}

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
    	log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	return result
}

func find_documents(collection *mongo.Collection)([]*Apartment){

	
	var results []*Apartment
	filter := bson.M{}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
    	log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

    
    	var elem Apartment
    	err := cur.Decode(&elem)
    	if err != nil {
        	log.Fatal(err)
    	}

    	results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
    	log.Fatal(err)
	}

	cur.Close(context.TODO())

	
	return results
}




type Apartment struct {
	Title         string  `json:"title"`
	HREF 		  string  `json:"href" validate:"required"`
	Adress 		  string  `json:"adress"`
	Rayon 		  string  `json:"rayon"`
	Price 		  int	  `json:"price"`
	Subprice 	  string  `json:"Subprice"`
	Phone_number  string  `json:"phone_number"`
	Description_text   string  `json:"description_text"`

}

	
func main(){

	/* flat := Apartment{"1-к квартира, 30 м², 12/22 эт.",
					"/moskva/kvartiry/1-k_kvartira_30_m_2234_et._2066460573",
					"Москва, трикотажная ул., 38к2",
					"Tushino",
					 33000,
					 "залог 33 000 ₽",
					 "8 962 650-12-76",
					 "Сдаётся на длительный срок",
				} */

	var list = []*Apartment{} 

 
 	collection := connect("real_estate")

 	//insert_document(collection, flat)
 	//flat = find_document(collection)
	//fmt.Println(flat)


 	list = find_documents(collection)


	fmt.Printf("Found multiple documents (array of pointers): %+v\n", list[1])


}


