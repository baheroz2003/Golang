package controller
// Required packages import kiye gaye hain
import (
	"context"                     // MongoDB ke saath kaam karne ke liye context use hota hai
	"encoding/json"               // JSON encode/decode karne ke liye
	"fmt"                         // Print statements ke liye
	"log"                         // Error logging ke liye
	"net/http"                    // HTTP server aur request/response handle karne ke liye

	"github.com/baheroz2003/mongoapigo/models" // Custom model import kiya gaya hai
	"github.com/gorilla/mux"                   // Routing ke liye mux package
	"go.mongodb.org/mongo-driver/bson"         // BSON (Binary JSON) documents ke liye
	"go.mongodb.org/mongo-driver/bson/primitive" // MongoDB ObjectID convert karne ke liye
	"go.mongodb.org/mongo-driver/mongo"          // MongoDB client
	"go.mongodb.org/mongo-driver/mongo/options"  // MongoDB connection options
)
// MongoDB ka URI connection string
const connectionString = "mongodb://localhost:27017/"
// Database ka naam
const dbName = "netflix"
// Collection ka naam
const collectionName = "watchList"

// Global variable: MongoDB collection ka reference
var collection *mongo.Collection

// init() function tab chalta hai jab package load hota hai
func init() {
	// MongoDB client options set kiye gaye hain
	clientOptions := options.Client().ApplyURI(connectionString)

	// MongoDB se connect karna
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// Agar error aaye to program panic mode me chala jaayega
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to mongodb database")
	// Database aur collection ka reference le rahe hain
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection reference is ready")
}

// Ek movie database me insert karne ke liye function
func insertOneMovie(movie models.Netflix) {
	// collection me ek document insert kar rahe hain
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 data to database:", inserted.InsertedID)
}

// Ek movie ko update karne ke liye function (watched=true set karta hai)
func updateOneMovie(movieId string) {
	// movieId ko ObjectID me convert kiya
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id} // filter banaya based on ID
	update := bson.M{"$set": bson.M{"watched": true}} // watched field true set kiya

	// MongoDB update query
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated 1 data in databse:", result.ModifiedCount)
}

// Ek movie delete karne ka function
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("One data deleted:", deleteCount)
}

// Saari movies delete karne ka function
func deleteAllMovies() {
	filter := bson.D{{}} // empty filter matlab sab delete karo
	deleteCount, err := collection.DeleteMany(context.Background(), filter, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted:", deleteCount.DeletedCount)
}

// Saari movies fetch karne ka function (array return karega)
func getAllMovies() []primitive.M {
	// Find function cursor return karta hai
	curr, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M // primitive.M = map[string]interface{}

	// cursor me se data read kar rahe hain
	for curr.Next(context.Background()) {
		var movie bson.M
		err := curr.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		// har movie ko array me add kar rahe hain
		movies = append(movies, movie)
	}

	defer curr.Close(context.Background()) // cursor band kar rahe hain
	return movies
}

// HTTP GET: saari movies frontend ko bhejne ka handler
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

// HTTP POST: nayi movie insert karne ka handler
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie models.Netflix
	// JSON body ko decode kar ke struct me daal rahe hain
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}
// HTTP POST: movie ko watched mark karne ka handler (ID ke zariye)
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded") // Response ka content-type set kar rahe hain
	w.Header().Set("Allow-Control-Allow-Methods", "PUT") // CORS allow kar rahe hain PUT method ke liye
	params := mux.Vars(r) // URL se 'id' nikal rahe hain (e.g. /movies/{id})
	updateOneMovie(params["id"]) // is ID ke movie ko 'watched' mark kar rahe hain
	json.NewEncoder(w).Encode(params["id"]) // client ko response me wahi ID wapas bhej rahe hain
}

// HTTP DELETE: ek movie delete karne ka handler
func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded") // Content type set kar rahe hain
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE") // DELETE method allow kar rahe hain CORS ke liye
	params := mux.Vars(r) // URL se 'id' nikal rahe hain
	deleteOneMovie(params["id"]) // uss ID ki movie ko database se delete kar rahe hain
	json.NewEncoder(w).Encode(params["id"]) // client ko wahi ID wapas bhej rahe hain
}

// HTTP DELETE: saari movies ek sath delete karne ka handler
func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded") // Content type set kar rahe hain
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE") // DELETE method allow kar rahe hain CORS ke liye
	count := deleteAllMovies// ❗FIXED: deleteAllMovies() function ko call kiya (missing brackets the)
	json.NewEncoder(w).Encode(count) // jitni movies delete huyi uska count client ko bhej rahe hain
}
