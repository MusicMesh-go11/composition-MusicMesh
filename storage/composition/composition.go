package compositions

import (
	pb "MusicMesh/composition-MusicMesh/generate/composition"
	"database/sql"
)

type CompositionRepo struct {
	pb.UnimplementedCompositionServiceServer
	DB *sql.DB
}

func NewCompositionRepo(db *sql.DB) *CompositionRepo {
	return &CompositionRepo{DB: db}
}
