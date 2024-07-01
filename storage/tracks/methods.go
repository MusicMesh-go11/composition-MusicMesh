package tracks

import (
	pb "MusicMesh/composition-MusicMesh/generate/track"
	"context"
)

func(t *TrackRepo) Create(ctx context.Context, in *pb.Track) (pb.Void, error){
	
}


Create(ctx context.Context, in *Composition) (*Void, error)
Get(ctx context.Context, in *FilterRequest) (*CompositionsRes, error)
GetByID(ctx context.Context, in *CompositionId) (*CompositionRes, error)
GetByUserID(ctx context.Context, in *UserId) (*CompositionRes, error)
Update(ctx context.Context, in *Composition) (*Void, error)
Delete(ctx context.Context, in *CompositionId) (*Void, error)