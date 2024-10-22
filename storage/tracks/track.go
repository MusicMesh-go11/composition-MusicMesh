package track

import (
	pb "MusicMesh/composition-MusicMesh/generate/tracks"
	"database/sql"
)

type TrackRepo struct {
	pb.UnimplementedTrackServiceServer
	DB *sql.DB
}

func NewTrackRepo(db *sql.DB) *TrackRepo {
	return &TrackRepo{DB: db}
}
