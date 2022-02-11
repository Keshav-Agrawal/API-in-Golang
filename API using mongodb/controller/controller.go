package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Keshav-Agrawal/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://keshav:keshav@cluster0.sjkrk.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "homework"
const colName = "task"

var collection *mongo.Collection

func init() {

	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}

func insertOneTask(work model.Homework) {
	inserted, err := collection.InsertOne(context.Background(), work)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 task in db with id: ", inserted.InsertedID)
}

func updateOneTask(workId string) {
	id, _ := primitive.ObjectIDFromHex(workId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"done": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

func deleteOneTask(workId string) {
	id, _ := primitive.ObjectIDFromHex(workId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("task got delete with delete count: ", deleteCount)
}

func deleteAllTask() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("NUmber of task delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func getAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var worklist []primitive.M

	for cur.Next(context.Background()) {
		var work bson.M
		err := cur.Decode(&work)
		if err != nil {
			log.Fatal(err)
		}
		worklist = append(worklist, work)
	}

	defer cur.Close(context.Background())
	return worklist
}

func GetMyAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allTask := getAllTask()
	json.NewEncoder(w).Encode(allTask)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var work model.Homework
	_ = json.NewDecoder(r.Body).Decode(&work)
	insertOneTask(work)
	json.NewEncoder(w).Encode(work)

}

func MarkAsDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteATask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllTask()
	json.NewEncoder(w).Encode(count)
}
