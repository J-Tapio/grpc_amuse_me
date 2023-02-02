package main

import (
	"log"
	"context"
	"net"

	amuseme "amuse-me/gen/amuseme"
	"google.golang.org/grpc"
)

type server struct {
	amuseme.AmuseMeServer
}

func (s *server) GetJoke(ctx context.Context, in *amuseme.Empty) (*amuseme.JokeReply, error) {
	return Joke()
}

func (s *server) GetMeme(ctx context.Context, in *amuseme.Empty) (*amuseme.MemeReply, error) {
	return Meme()
}


func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	amuseme.RegisterAmuseMeServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}