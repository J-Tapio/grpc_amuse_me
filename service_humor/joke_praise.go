package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"amuse-me/gen/amuseme"
)

type RedditJoke struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

func randomJoke(fileName string) RedditJoke {
	// Get joke by random between 1-250
	rand.Seed(time.Now().UnixNano())
	nthJoke := rand.Intn(250)

	var joke RedditJoke
	file, _ := os.Open("./reddit_jokes.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	// Read the array open bracket
	decoder.Token()

	for i := 1; ; i++ {
		decoder.Decode(&joke)
		if i == nthJoke {
			break
		}
		if decoder.More() {
			// Carry on until nth joke if possible
			continue
		} else {
			break
		}
	}

	return joke
}

func Joke() (*amuseme.JokeReply, error) {
	// Random joke from dataset (json-file)
	//var RedditJokes []*RedditJoke
	redditJoke := randomJoke("./reddit_jokes.json")

	return &amuseme.JokeReply{Title: redditJoke.Title, Body: redditJoke.Body}, nil
}

func Meme() (*amuseme.MemeReply, error) {
	err := godotenv.Load()
  if err != nil {
    log.Println("Error loading .env file")
		return &amuseme.MemeReply{}, err
  }

	// Random meme
	url := "https://humor-jokes-and-memes.p.rapidapi.com/memes/random?media-type=image"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return &amuseme.MemeReply{}, err
	}
	memeAPIKey := os.Getenv("MEME_API_KEY")
	req.Header.Add("X-RapidAPI-Key", memeAPIKey)
	req.Header.Add("X-RapidAPI-Host", "humor-jokes-and-memes.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println(err)
		return &amuseme.MemeReply{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &amuseme.MemeReply{}, err
	}

	return &amuseme.MemeReply{Meme: string(body)}, nil
}
