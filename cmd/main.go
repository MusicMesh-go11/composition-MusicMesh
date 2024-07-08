package main

import (
	"MusicMesh/composition-MusicMesh/generate/composition"
	"MusicMesh/composition-MusicMesh/generate/tracks"
	compositions "MusicMesh/composition-MusicMesh/storage/composition"
	"MusicMesh/composition-MusicMesh/storage/postgres"
	track "MusicMesh/composition-MusicMesh/storage/tracks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

func main() {

	conn, err := grpc.NewClient(":5051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

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
	c := compositions.NewCompositionRepo(db, conn)
	t := track.NewTrackRepo(db)

	composition.RegisterCompositionServiceServer(server, c)
	tracks.RegisterTrackServiceServer(server, t)

	log.Println("Listening on :5052")
	panic(server.Serve(listner))
}
