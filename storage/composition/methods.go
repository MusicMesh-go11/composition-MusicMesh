package compositions

import (
	pb "MusicMesh/composition-MusicMesh/generate/composition"
	"context"
)

func (c *CompositionRepo) Create(ctx context.Context, in *pb.Composition) (*pb.Void, error) {
	_, err := c.DB.Exec(`INSERT INTO compositions (user_id, title, description, status)
	VALUES ($1, $2, $3, $4)`, in.UserId, in.Title, in.Description, in.Status)
	return &pb.Void{}, err
}

func (c *CompositionRepo) Get(ctx context.Context, in *pb.FilterRequest) (*pb.CompositionsRes, error) {
	query := "SELECT composition_id, user_id, title, description, status, created_at, updated_at FROM compositions WHERE deleted_at = 0"
	if len(in.Query) > 0 {
		query += " AND (" + in.Query + ")"
	}
	rows, err := c.DB.Query(query, in.Arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	compositions := &pb.CompositionsRes{}
	for rows.Next() {
		comp := &pb.CompositionRes{}
		err = rows.Scan(&comp.CompositionID, &comp.UserId, &comp.Title, &comp.Description, &comp.Status, &comp.CreateAt)
		if err != nil {
			return nil, err
		}
		compositions.Compositions = append(compositions.Compositions, comp)
	}
	return compositions, nil
}

func (c *CompositionRepo) GetByID(ctx context.Context, in *pb.CompositionId) (*pb.CompositionRes, error) {
	row := c.DB.QueryRow("SELECT composition_id, user_id, title, description, status, created_at, updated_at FROM compositions WHERE composition_id = $1 AND deleted_at = 0", in.CompositionID)
	comp := &pb.CompositionRes{}
	err := row.Scan(&comp.CompositionID, &comp.UserId, &comp.Title, &comp.Description, &comp.Status, &comp.CreateAt)
	if err != nil {
		return nil, err
	}
	return comp, nil
}

func (c *CompositionRepo) GetByUserID(ctx context.Context, in *pb.UserId) (*pb.CompositionsRes, error) {
	rows, err := c.DB.Query("SELECT composition_id, user_id, title, description, status, created_at, updated_at FROM compositions WHERE user_id = $1 AND deleted_at = 0", in.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	compositions := &pb.CompositionsRes{}
	for rows.Next() {
		comp := &pb.CompositionRes{}
		err = rows.Scan(&comp.CompositionID, &comp.UserId, &comp.Title, &comp.Description, &comp.Status, &comp.CreateAt)
		if err != nil {
			return nil, err
		}
		compositions.Compositions = append(compositions.Compositions, comp)
	}
	return compositions, nil
}

func (c *CompositionRepo) Update(ctx context.Context, in *pb.Composition) (*pb.Void, error) {
	_, err := c.DB.Exec(`UPDATE compositions SET title = $1, description = $2, status = $3, updated_at = current_timestamp WHERE composition_id = $4 AND deleted_at = 0`,
		in.Title, in.Description, in.Status, in.CompositionID)
	return &pb.Void{}, err
}

func (c *CompositionRepo) Delete(ctx context.Context, in *pb.CompositionId) (*pb.Void, error) {
	_, err := c.DB.Exec(`UPDATE compositions SET deleted_at = date_part('epoch', current_timestamp)::INT WHERE composition_id = $1 AND deleted_at = 0`, in.CompositionID)
	return &pb.Void{}, err
}
