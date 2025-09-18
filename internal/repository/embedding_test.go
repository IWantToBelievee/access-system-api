package repository

import (
	"context"
	"database/sql"
	"testing"

	"access-system-api/internal/cfg"
	"access-system-api/internal/client"
	"access-system-api/internal/domain"

	_ "github.com/lib/pq"
	"github.com/pgvector/pgvector-go"
)

func cleanEmbeddingsTable(db *sql.DB) {
	db.Exec("DELETE FROM embedding")
}

func TestEmbeddingRepository_CRUD(t *testing.T) {
	dbCfg, err := cfg.LoadTestDbCfg()
	if err != nil {
		t.Fatalf("failed to load test db config: %v", err)
	}

	db, err := client.ConnectDB(dbCfg)
	if err != nil {
		t.Fatalf("failed to connect to test db: %v", err)
	}
	defer db.Close()
	cleanEmbeddingsTable(db)
	ctx := context.Background()

	repo := NewEmbeddingsRepository(db)

	var vector []float32
	for i := 0; i < 512; i++ {
		vector = append(vector, 0.1)
	}

	// Test CreateEmbedding
	emb := &domain.Embedding{
		Name:   "test-embedding",
		Vector: pgvector.NewVector(vector),
	}
	err = repo.CreateEmbedding(ctx, emb)
	if err != nil {
		t.Fatalf("CreateEmbedding failed: %v", err)
	}

	// Test GetSimilarEmbeddingByVector (should find)
	found, err := repo.GetSimilarEmbeddingByVector(ctx, pgvector.NewVector(vector))
	if err != nil {
		t.Fatalf("GetSimilarEmbeddingByVector failed: %v", err)
	}
	if found == nil || found.Name != emb.Name {
		t.Errorf("GetSimilarEmbeddingByVector returned wrong embedding: got %+v", found)
	}

	// Test DeleteEmbeddingById
	if found != nil {
		err = repo.DeleteEmbeddingById(ctx, found.ID)
		if err != nil {
			t.Fatalf("DeleteEmbeddingById failed: %v", err)
		}
	}

	// Test DeleteEmbeddingById (non-existent)
	err = repo.DeleteEmbeddingById(ctx, 999999)
	if err != nil {
		t.Errorf("DeleteEmbeddingById for non-existent id should not fail, got %v", err)
	}

	cleanEmbeddingsTable(db)
}

func TestEmbeddingRepository_DBError(t *testing.T) {
	dbCfg, err := cfg.LoadTestDbCfg()
	if err != nil {
		t.Fatalf("failed to load test db config: %v", err)
	}

	db, err := client.ConnectDB(dbCfg)
	if err != nil {
		t.Fatalf("failed to connect to test db: %v", err)
	}
	defer db.Close()
	ctx := context.Background()

	repo := NewEmbeddingsRepository(db)
	_ = db.Close() // forcibly close to simulate error

	var vector []float32
	for i := 0; i < 512; i++ {
		vector = append(vector, 0.1)
	}

	emb := &domain.Embedding{Name: "fail", Vector: pgvector.NewVector(vector)}
	if err := repo.CreateEmbedding(ctx, emb); err == nil {
		t.Error("expected error on CreateEmbedding with closed db")
	}
	if _, err := repo.GetSimilarEmbeddingByVector(ctx, emb.Vector); err == nil {
		t.Error("expected error on GetSimilarEmbeddingByVector with closed db")
	}
	if err := repo.DeleteEmbeddingById(ctx, 1); err == nil {
		t.Error("expected error on DeleteEmbeddingById with closed db")
	}
}
