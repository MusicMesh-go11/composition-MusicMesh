package compositions

import (
	pb "MusicMesh/composition-MusicMesh/generate/composition"
	"MusicMesh/composition-MusicMesh/generate/user"
	"database/sql"
	"google.golang.org/grpc"
)

type CompositionRepo struct {
	User user.UserServiceClient
	pb.UnimplementedCompositionServiceServer
	DB *sql.DB
}

func NewCompositionRepo(db *sql.DB, conn *grpc.ClientConn) *CompositionRepo {
	userClient := user.NewUserServiceClient(conn)
	return &CompositionRepo{DB: db, User: userClient}
}
