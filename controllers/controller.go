package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shivendutyagi/newapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"gopkg.in/mgo.v2/bson"
)

const connectionstring = "mongodb+srv://shivendutyagi:Shivendutyagi@cluster0.ke4hfmm.mongodb.net/test"
const dbname = "Netflix"
const colname = "watchlist"

var collection *mongo.Collection

func init() {
	clientoption := options.Client().ApplyURI(connectionstring)

	client, err := mongo.Connect(context.TODO(), clientoption)

	if err != nil {
		log.Fatal(err)
	}

	collection = (*mongo.Collection)(client.Database(dbname).Collection(colname))

}

func insertone(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Inserted vallue", inserted.InsertedID)

}

func updateonemovie(movieid string) {
	id, _ := primitive.ObjectIDFromHex(movieid)

	filter := bson.M{"_id": id}
	updated := bson.M{"$set": bson.M{"watched": true}}
	result, _ := collection.UpdateOne(context.Background(), filter, updated)

	fmt.Println("Updated", result.ModifiedCount)

}

func deleteonemovie(movieid string) {
	id, _ := primitive.ObjectIDFromHex(movieid)

	filter := bson.M{"_id": id}
	deletec, _ := collection.DeleteOne(context.Background(), filter)
	fmt.Print(deletec)

}

func getallmovies() []primitive.M {
	curr, _ := collection.Find(context.Background(), bson.D{{}})

	var movies []primitive.M

	for curr.Next(context.Background()) {
		var movie bson.M

		curr.Decode(&movie)
		movies = append(movies, movie)
	}
	defer curr.Close(context.Background())

	return movies

}

func Getallmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allmovies := getallmovies()
	json.NewEncoder(w).Encode(allmovies)

}

func Createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	json.NewDecoder(r.Body).Decode(&movie)

	insertone(movie)
	json.NewEncoder(w).Encode(movie)

}

func Marksaswatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateonemovie(params["id"])
	json.NewEncoder(w).Encode(params)

}
func Deleteamovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteonemovie(params["id"])
	json.NewEncoder(w).Encode(params)
}
