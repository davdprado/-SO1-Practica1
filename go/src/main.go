package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go/src/db"
	"go/src/entorno"
	"go/src/modelos"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var numMongo int
var clienteMongo *mongo.Client
var tiempoMongo float64

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my api")
}

func nuevaOp(w http.ResponseWriter, r *http.Request) {
	var operacion modelos.Operacion
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(false)
		return
	}
	json.Unmarshal(reqBody, &operacion)
	fmt.Println(operacion)
	err = insertar(operacion)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(false)
		return
	}
	json.NewEncoder(w).Encode(true)
}

func insertar(o modelos.Operacion) error {
	fmt.Print(o)
	if clienteMongo != nil {
		tInicio := time.Now()
		collection := clienteMongo.Database("practica1").Collection("calculadora")
		insertResult, err := collection.InsertOne(context.TODO(), o)
		if err != nil {
			fmt.Println("the collection does not exist in MongoDB ")
			return err
		}
		fmt.Println("Publicacion had been inserted: ", insertResult.InsertedID)
		numMongo++
		tFinal := time.Now()
		tiempoMongo += (float64(tFinal.Hour())-float64(tInicio.Hour()))*3600 + (float64(tFinal.Minute())-float64(tInicio.Minute()))*60 + (float64(tFinal.Second()) - float64(tInicio.Second())) + (float64(tFinal.UnixMilli())-float64(tInicio.UnixMilli()))*0.001
		return nil
	} else {
		fmt.Println("MongoDB is not connected")
		return errors.New("unavailable")
	}
}

func ConnectMongo() error {
	var err error
	clienteMongo, err = mongo.Connect(context.TODO(), db.GetClient())
	if err != nil {
		log.Println(err)
		return err
	}
	err = clienteMongo.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println("MongoDB is connected to: " + os.Getenv("MONGO_NAME"))
	return nil
}

func obtenerdatos(w http.ResponseWriter, r *http.Request) {
	collectionO := clienteMongo.Database("practica1").Collection("calculadora")
	cursor, err := collectionO.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var ope []bson.M
	if err = cursor.All(context.TODO(), &ope); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(ope)

	data, _ := json.Marshal(ope)
	fmt.Println(string(data))

	json.NewEncoder(w).Encode(string(data))

}

func main() {
	//Iniciamos variables de entorno
	entorno.LoadEnv()
	ConnectMongo()
	port := os.Getenv("PORT")
	//Creamos el router
	router := mux.NewRouter().StrictSlash(true)
	//Ruta principal
	router.HandleFunc("/", indexRoute)
	//Escuchamos al puerto
	router.HandleFunc("/insertar", nuevaOp).Methods("POST")
	router.HandleFunc("/obtener", obtenerdatos).Methods("GET")

	fmt.Println("Server on port:" + port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
