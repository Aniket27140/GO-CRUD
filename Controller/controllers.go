package controller

import (
	model "CRUD/Model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "github.com/goccy/go-json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// db connection prerequisites
const connectionstring = "mongodb+srv://vishwakarmaaniket706:ddgd0uSXQyKsEoZL@cluster0.pf8rq.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const colname = "watchlist"

const dbname = "netflix"

// reference of mongodb which will be used for performing crud
// collections of type *mongo.Collection to store the reference to the MongoDB collection.
var collections *mongo.Collection

// The init() function is a special function in Go that runs automatically before the main() function
func init() {
	// 2) setup client options
	clientOptions := options.Client().ApplyURI(connectionstring)

	// 1) connection to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// 3) creating dbname & collection in db
	collections = client.Database(dbname).Collection(colname)

	fmt.Println("collection instance is ready")
}

func insertOneMovie(movie model.Netflix) {
	inserted, err := collections.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted one movie with id:", inserted.InsertedID)
}

func updateOneMovie(movieID string) {
	// The purpose of using the ObjectIDFromHex function in the provided code is to convert a hexadecimal string representation of an ObjectID into an ObjectID type.
	// This is necessary because MongoDB uses ObjectIDs as unique identifiers for its documents.
	id, _ := primitive.ObjectIDFromHex(movieID)

	update := bson.M{"$set": bson.M{"id": id}}

	filter := bson.M{"_id": id}

	result, err := collections.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("updated one movie", result.ModifiedCount)
}

func deleteOneMovie(movieID string) {
	// ObjectIDFromHex creates a new ObjectID from a hex string. It returns an error if the hex string is not a valid ObjectID.
	id, _ := primitive.ObjectIDFromHex(movieID)

	filter := bson.M{"_id": id}

	deleteresult, err := collections.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted one movie", deleteresult)
}

func deletemany() {
	// filter := bson.M{}
	delete, err := collections.DeleteMany(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleted all movies", delete.DeletedCount)
}

func getallmovies() []primitive.M {
	cur, err := collections.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

func Getallmymovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allmovies := getallmovies()
	// now we have a object but we deal in json in http, convert into json
	json.NewEncoder(w).Encode(allmovies)
}

func Createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	// created var bcz json is coming from fron-end, need to store it
	var movie model.Netflix
	// decoding json into object & storing the json into var movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)
	// returning json
	json.NewEncoder(w).Encode(movie)
}

func MarkWatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	// getting movie-id from params
	params := mux.Vars(r)

	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
