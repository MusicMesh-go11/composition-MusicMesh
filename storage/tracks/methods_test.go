package track

//func TestTrackRepo_Create(t *testing.T) {
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	repo := &TrackRepo{DB: db}
//
//	compositionID := "037aac5c-fae7-4369-a35f-2570b8661c48"
//	userID := "36be4369-6df1-444b-9b8c-1e52c36fbcf7"
//
//	testTrack := &pb.Track{
//		CompositionID: compositionID,
//		UserID:        userID,
//		Title:         "Test Track",
//		FileUrl:       "https://example.com/test.mp3",
//	}
//
//	_, err = repo.Create(context.Background(), testTrack)
//	assert.NoError(t, err)
//
//	var count int
//	err = db.QueryRow("SELECT COUNT(*) FROM tracks WHERE composition_id = $1 AND user_id = $2 AND title = $3 AND file_url = $4",
//		testTrack.CompositionID, testTrack.UserID, testTrack.Title, testTrack.FileUrl).Scan(&count)
//	assert.NoError(t, err)
//	assert.Equal(t, 1, count)
//}

//func TestTrackRepo_GetByID(t *testing.T) {
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	repo := &TrackRepo{DB: db}
//
//	// Insert a test track
//	trackID := "6d521f8e-d094-4a66-9f84-5cc6a200d8d0"
//	compositionID := "037aac5c-fae7-4369-a35f-2570b8661c48"
//
//	_, err = db.Exec(`INSERT INTO tracks (track_id, composition_id, user_id, title, file_url)
//      VALUES ($1, $2, $3, 'Test Track', 'https://example.com/test.mp3')
//      ON CONFLICT (track_id) DO NOTHING`, trackID, compositionID, compositionID)
//	assert.NoError(t, err)
//
//	trackIdReq := &pb.TrackId{TrackId: trackID}
//	track, err := repo.GetByID(context.Background(), trackIdReq)
//
//	assert.NoError(t, err)
//	assert.Equal(t, trackID, track.TracId)
//	assert.Equal(t, "Test Track", track.Title)
//	assert.Equal(t, "https://example.com/test.mp3", track.FileUrl)
//}

//func TestTrackRepo_GetByCompositionID(t *testing.T) {
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	repo := &TrackRepo{DB: db}
//
//	// Insert a test track
//	trackID := "6d521f8e-d094-4a66-9f84-5cc6a200d8d0"
//	compositionID := "037aac5c-fae7-4369-a35f-2570b8661c48"
//
//	_, err = db.Exec(`INSERT INTO tracks (track_id, composition_id, user_id, title, file_url)
//	      VALUES ($1, $2, $3, 'Test Track', 'https://example.com/test.mp3')
//	      ON CONFLICT (track_id) DO NOTHING`, trackID, compositionID, compositionID)
//	assert.NoError(t, err)
//
//	compositionIDReq := &pb.CompositionID{CompositionID: compositionID}
//	tracks, err := repo.GetByCompositionID(context.Background(), compositionIDReq)
//
//	assert.NoError(t, err)
//	assert.Len(t, tracks.Tracks, 2)
//	assert.Equal(t, "Track 1", tracks.Tracks[0].Title)
//	assert.Equal(t, "Track 2", tracks.Tracks[1].Title)
//}

//func TestTrackRepo_Update(t *testing.T) {
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	repo := &TrackRepo{DB: db}
//
//	// Insert a test track
//	trackID := "36be4369-6df1-444b-9b8c-1e52c36fbcf8"
//	_, err = db.Exec(`INSERT INTO tracks (track_id, composition_id, user_id, title, file_url, created_at, updated_at, deleted_at)
//       VALUES ($1, 1, 1, 'Original Track', 'https://example.com/original.mp3', now(), now(), 0)
//       ON CONFLICT (track_id) DO NOTHING`, trackID)
//	assert.NoError(t, err)
//
//	updatedTrack := &pb.TrackRes{
//		TracId:  trackID,
//		Title:   "Updated Track",
//		FileUrl: "https://example.com/updated.mp3",
//	}
//
//	_, err = repo.Update(context.Background(), updatedTrack)
//	assert.NoError(t, err)
//
//	var title, fileUrl string
//	err = db.QueryRow("SELECT title, file_url FROM tracks WHERE track_id = $1", trackID).Scan(&title, &fileUrl)
//	assert.NoError(t, err)
//	assert.Equal(t, "Updated Track", title)
//	assert.Equal(t, "https://example.com/updated.mp3", fileUrl)
//}

//
//func TestTrackRepo_Delete(t *testing.T) {
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	repo := &TrackRepo{DB: db}
//
//	// Insert a test track
//	trackID := "36be4369-6df1-444b-9b8c-1e52c36fbcf8"
//	_, err = db.Exec(`INSERT INTO tracks (track_id, composition_id, user_id, title, file_url, created_at, updated_at, deleted_at)
//        VALUES ($1, 1, 1, 'Test Track', 'https://example.com/test.mp3', now(), now(), 0)
//        ON CONFLICT (track_id) DO NOTHING`, trackID)
//	assert.NoError(t, err)
//
//	trackIdReq := &pb.TrackId{TrackId: trackID}
//	_, err = repo.Delete(context.Background(), trackIdReq)
//	assert.NoError(t, err)
//
//	var deletedAt int64
//	err = db.QueryRow("SELECT deleted_at FROM tracks WHERE track_id = $1", trackID).Scan(&deletedAt)
//	assert.NoError(t, err)
//	assert.NotEqual(t, 0, deletedAt)
//}
