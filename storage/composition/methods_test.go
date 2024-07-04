package compositions

import (
	pb "MusicMesh/composition-MusicMesh/generate/composition"
	"MusicMesh/composition-MusicMesh/storage/postgres"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//func TestCompositionRepo_Create(t *testing.T) {
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	repo := &CompositionRepo{DB: db}
//
//	userID := "36be4369-6df1-444b-9b8c-1e52c36fbcf7"
//
//	testComposition := &pb.Composition{
//		UserId:      userID,
//		Title:       "Test Composition",
//		Description: "This is a test composition",
//		Status:      "draft",
//	}
//
//	_, err = repo.Create(context.Background(), testComposition)
//
//	assert.NoError(t, err)
//
//	var count int
//	err = db.QueryRow("SELECT COUNT(*) FROM compositions WHERE user_id = $1 AND title = $2 AND description = $3 AND status = $4",
//		testComposition.UserId, testComposition.Title, testComposition.Description, testComposition.Status).Scan(&count)
//	if err != nil {
//		t.Fatalf("Failed to query compositions table, %v", err)
//	}
//
//	assert.Equal(t, 1, count)
//}

//func TestCompositionRepo_GetByID(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &CompositionRepo{DB: db}
//
//	// Insert a test composition into the compositions table if not already present
//	compositionID := "037aac5c-fae7-4369-a35f-2570b8661c48"
//	// _, err = db.Exec(`INSERT INTO compositions (composition_id, user_id, title, description, status, created_at, updated_at, deleted_at)
//	//     VALUES ($1, 1, 'Test Composition', 'This is a test composition', 'draft', now(), now(), 0)
//	//     ON CONFLICT (composition_id) DO NOTHING`, compositionID)
//	// if err != nil {
//	//     t.Fatalf("Failed to insert test data, %v", err)
//	// }
//
//	// Create a CompositionId request
//	compositionId := pb.CompositionId{CompositionID: compositionID}
//
//	// Call the GetByID method
//	compositionRes, err := repo.GetByID(context.Background(), &compositionId)
//
//	// Assert the GetByID method returned no error
//	assert.NoError(t, err)
//
//	// Verify the details of the composition
//	assert.Equal(t, "037aac5c-fae7-4369-a35f-2570b8661c48", compositionRes.CompositionID)
//	assert.Equal(t, "36be4369-6df1-444b-9b8c-1e52c36fbcf7", compositionRes.UserId)
//	assert.Equal(t, "Test Composition", compositionRes.Title)
//	assert.Equal(t, "This is a test composition", compositionRes.Description)
//	assert.Equal(t, "draft", compositionRes.Status)
//	assert.NotNil(t, compositionRes.CreateAt)
//}

//func TestCompositionRepo_GetByUserID(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository
//	repo := &CompositionRepo{DB: db}
//
//	// Insert a test composition
//	userID := "df004dcb-4207-43d6-a028-abb48991b613"
//	_, err = db.Exec(`INSERT INTO compositions (user_id, title, description, status, created_at, updated_at, deleted_at)
//       VALUES ($1, 'Test Composition', 'Test Description', 'draft', now(), now(), 0)`, userID)
//	if err != nil {
//		t.Fatalf("Failed to insert test data, %v", err)
//	}
//
//	// Call the GetByUserID method
//	userIdReq := &pb.UserId{UserID: userID}
//	compositionsRes, err := repo.GetByUserID(context.Background(), userIdReq)
//
//	// Check for errors
//	assert.NoError(t, err)
//
//	// Check if we got any compositions
//	assert.NotEmpty(t, compositionsRes.Compositions)
//
//	// Check the details of the composition
//	comp := compositionsRes.Compositions[0]
//	assert.Equal(t, userID, comp.UserId)
//	assert.Equal(t, "Test Composition", comp.Title)
//	assert.Equal(t, "Test Description", comp.Description)
//	assert.Equal(t, "draft", comp.Status)
//}

//func TestCompositionRepo_Update(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &CompositionRepo{DB: db}
//
//	// Insert a test composition into the compositions table if not already present
//	compositionID := "df004dcb-4207-43d6-a028-abb48991b614"
//	userID := "df004dcb-4207-43d6-a028-abb48991b614"
//	_, err = db.Exec(`INSERT INTO compositions (composition_id, user_id, title, description, status)
//       VALUES ($1, $2, 'Original Title', 'Original Description', 'draft')
//       ON CONFLICT (composition_id) DO NOTHING`, compositionID, userID)
//	if err != nil {
//		t.Fatalf("Failed to insert test data, %v", err)
//	}
//
//	// Define the updated composition data
//	updatedComposition := &pb.CompositionRes{
//		CompositionID: compositionID,
//		Title:         "Updated Title",
//		Description:   "Updated Description",
//		Status:        "published",
//	}
//
//	// Call the Update method
//	_, err = repo.Update(context.Background(), updatedComposition)
//
//	// Assert the Update method returned no error
//	assert.NoError(t, err)
//
//	// Verify the composition details were updated correctly by querying the database
//	var title, description, status string
//	var updatedAt time.Time
//	err = db.QueryRow("SELECT title, description, status, updated_at FROM compositions WHERE composition_id = $1", compositionID).Scan(&title, &description, &status, &updatedAt)
//	if err != nil {
//		t.Fatalf("Failed to query compositions table, %v", err)
//	}
//
//	// Assert that the composition details were updated correctly
//	assert.Equal(t, "Updated Title", title)
//	assert.Equal(t, "Updated Description", description)
//	assert.Equal(t, "published", status)
//	assert.True(t, updatedAt.After(time.Now().Add(-time.Minute)), "updated_at should be recent")
//}

func TestCompositionRepo_Delete(t *testing.T) {
	// Connect to the test database
	db, err := postgres.Conn()
	if err != nil {
		t.Fatalf("Failed to open sql db, %v", err)
	}
	defer db.Close()

	// Initialize the repository with the test database connection
	repo := &CompositionRepo{DB: db}

	// Insert a test composition into the compositions table if not already present
	compositionID := "df004dcb-4207-43d6-a028-abb48991b614"
	userID := "df004dcb-4207-43d6-a028-abb48991b614"
	_, err = db.Exec(`INSERT INTO compositions (composition_id, user_id, title, description, status)
	      VALUES ($1, $2, 'Original Title', 'Original Description', 'draft')
	      ON CONFLICT (composition_id) DO NOTHING`, compositionID, userID)
	if err != nil {
		t.Fatalf("Failed to insert test data, %v", err)
	}

	// Define the composition ID to delete
	compositionToDelete := &pb.CompositionId{CompositionID: compositionID}

	// Call the Delete method
	_, err = repo.Delete(context.Background(), compositionToDelete)

	// Assert the Delete method returned no error
	assert.NoError(t, err)

	// Verify the composition details were updated correctly by querying the database
	var deletedAt int64
	err = db.QueryRow("SELECT deleted_at FROM compositions WHERE composition_id = $1", compositionID).Scan(&deletedAt)
	if err != nil {
		t.Fatalf("Failed to query compositions table, %v", err)
	}

	// Assert that the composition was marked as deleted
	assert.NotEqual(t, 0, deletedAt, "Expected deleted_at to be set to a non-zero value")

	// Additional check: Ensure the deleted_at is close to the current time
	currentEpoch := time.Now().Unix()
	assert.InDelta(t, currentEpoch, deletedAt, 5, "deleted_at should be close to the current time")
}
