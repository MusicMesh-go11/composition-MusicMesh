package tracks

import (
	pb "MusicMesh/composition-MusicMesh/generate/tracks"
	"context"
)

func (t *TrackRepo) Create(ctx context.Context, in *pb.Track) (*pb.Void, error) {
	_, err := t.DB.Exec(`INSERT INTO tracks (composition_id, user_id, title, file_url)
	VALUES ($1, $2, $3, $4)`, in.CompositionID, in.UserID, in.Title, in.FileUrl)
	return &pb.Void{}, err
}

func (t *TrackRepo) Get(ctx context.Context, in *pb.FilterRequest) (*pb.TracksRes, error) {
	query := "SELECT track_id, composition_id, user_id, title, file_url, created_at, updated_at FROM tracks WHERE deleted_at = 0"
	if len(in.Query) > 0 {
		query += " AND (" + in.Query + ")"
	}
	rows, err := t.DB.Query(query, in.Arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tracks := &pb.TracksRes{}
	for rows.Next() {
		track := &pb.TrackRes{}
		err = rows.Scan(&track.TracId, &track.CompositionID, &track.UserID, &track.Title, &track.FileUrl, &track.CreateAt)
		if err != nil {
			return nil, err
		}
		tracks.Tracks = append(tracks.Tracks, track)
	}
	return tracks, nil
}

func (t *TrackRepo) GetByID(ctx context.Context, in *pb.TrackId) (*pb.TrackRes, error) {
	row := t.DB.QueryRow("SELECT track_id, composition_id, user_id, title, file_url, created_at, updated_at FROM tracks WHERE track_id = $1 AND deleted_at = 0", in.TrackId)
	track := &pb.TrackRes{}
	err := row.Scan(&track.TracId, &track.CompositionID, &track.UserID, &track.Title, &track.FileUrl, &track.CreateAt)
	if err != nil {
		return nil, err
	}
	return track, nil
}

func (t *TrackRepo) GetByCompositionID(ctx context.Context, in *pb.CompositionID) (*pb.TracksRes, error) {
	rows, err := t.DB.Query("SELECT track_id, composition_id, user_id, title, file_url, created_at, updated_at FROM tracks WHERE composition_id = $1 AND deleted_at = 0", in.CompositionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tracks := &pb.TracksRes{}
	for rows.Next() {
		track := &pb.TrackRes{}
		err = rows.Scan(&track.TracId, &track.CompositionID, &track.UserID, &track.Title, &track.FileUrl, &track.CreateAt)
		if err != nil {
			return nil, err
		}
		tracks.Tracks = append(tracks.Tracks, track)
	}
	return tracks, nil
}

func (t *TrackRepo) Update(ctx context.Context, in *pb.TrackRes) (*pb.Void, error) {
	_, err := t.DB.Exec(`UPDATE tracks SET title = $1, file_url = $2, updated_at = current_timestamp WHERE track_id = $3 AND deleted_at = 0`,
		in.Title, in.FileUrl, in.TracId)
	return &pb.Void{}, err
}

func (t *TrackRepo) Delete(ctx context.Context, in *pb.TrackId) (*pb.Void, error) {
	_, err := t.DB.Exec(`UPDATE tracks SET deleted_at = date_part('epoch', current_timestamp)::INT WHERE track_id = $1 AND deleted_at = 0`, in.TrackId)
	return &pb.Void{}, err
}
