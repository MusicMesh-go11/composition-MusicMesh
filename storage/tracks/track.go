package tracks

import (
	pb "MusicMesh/composition-MusicMesh/generate/track"
	"database/sql"
)

type TrackRepo struct {
	pb.UnimplementedTrackServiceServer
	DB *sql.DB
}

func NewTrackRepo(db *sql.DB) *TrackRepo{
	return &TrackRepo{DB: db}
}