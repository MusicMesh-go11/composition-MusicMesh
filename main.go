package main

import (
	"MusicMesh/composition-MusicMesh/generate/composition"
	"MusicMesh/composition-MusicMesh/generate/tracks"
	compositions "MusicMesh/composition-MusicMesh/storage/composition"
	"MusicMesh/composition-MusicMesh/storage/postgres"
	track "MusicMesh/composition-MusicMesh/storage/tracks"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db, err := postgres.Conn()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	listner, err := net.Listen("tcp", ":5052")
	if err != nil {
		panic(err)
	}
	defer listner.Close()

	server := grpc.NewServer()
	c := compositions.NewCompositionRepo(db)
	t := track.NewTrackRepo(db)

	composition.RegisterCompositionServiceServer(server, c)
	tracks.RegisterTrackServiceServer(server, t)

	log.Println("Listening on :5052")
	panic(server.Serve(listner))
}
