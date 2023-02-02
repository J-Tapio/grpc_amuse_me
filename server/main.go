package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	amuseme "amuse-me/gen/amuseme"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serviceJokeConn *grpc.ClientConn

func connectServiceJoke() *grpc.ClientConn {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to service: %v", err)
	}
	return conn
}

func serveJoke(w http.ResponseWriter, r *http.Request) {
	c := amuseme.NewAmuseMeClient(serviceJokeConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	joke, err := c.GetJoke(ctx, &amuseme.Empty{})
	if err != nil {
		log.Printf("Could not get a joke: %v", err)
		w.WriteHeader(500)
	}

	cancel()

	jokeJSON, _ := json.Marshal(joke)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	w.Write(jokeJSON)
}

func serveMeme(w http.ResponseWriter, r *http.Request) {
	c := amuseme.NewAmuseMeClient(serviceJokeConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	joke, err := c.GetMeme(ctx, &amuseme.Empty{})
	if err != nil {
		log.Printf("Could not get a meme: %v", err)
		w.WriteHeader(500)
	}

	cancel()

	jokeJSON, _ := json.Marshal(joke)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	w.Write(jokeJSON)
}

func startServer(port string) {
	http.HandleFunc("/", serveJoke)
	http.HandleFunc("/meme", serveMeme)
	err := http.ListenAndServe(":"+port, nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	// Connect with service (jokes, memes)
	serviceJokeConn = connectServiceJoke()
	defer serviceJokeConn.Close()

	startServer("5050")
}
