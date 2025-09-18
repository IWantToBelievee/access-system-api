package service

import (
	"context"
	"testing"

	"access-system-api/internal/domain"
	mocks "access-system-api/internal/mocks/repository"

	"github.com/golang/mock/gomock"
	"github.com/pgvector/pgvector-go"
	"github.com/stretchr/testify/assert"
)

func TestEmbeddingService_AddEmbedding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEmbeddingRepository(ctrl)
	service := NewEmbeddingService(repo)

	ctx := context.Background()
	name := "test"
	vector := make([]float32, 512)
	for i := range vector {
		vector[i] = float32(i)
	}

	embedding := &domain.Embedding{
		Name:   name,
		Vector: pgvector.NewVector(vector),
	}

	repo.EXPECT().CreateEmbedding(ctx, embedding).Return(nil)

	err := service.AddEmbedding(ctx, name, vector)
	assert.NoError(t, err)
}

func TestEmbeddingService_AddEmbedding_InvalidVectorSize(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEmbeddingRepository(ctrl)
	service := NewEmbeddingService(repo)

	ctx := context.Background()
	name := "test"
	vector := make([]float32, 100) // Invalid size

	err := service.AddEmbedding(ctx, name, vector)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "vector size must be 512")
}

func TestEmbeddingService_ValidateEmbedding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEmbeddingRepository(ctrl)
	service := NewEmbeddingService(repo)

	ctx := context.Background()
	vector := make([]float32, 512)
	for i := range vector {
		vector[i] = float32(i)
	}

	repo.EXPECT().GetSimilarEmbeddingByVector(ctx, pgvector.NewVector(vector)).Return(&domain.Embedding{}, nil)

	err := service.ValidateEmbedding(ctx, vector)
	assert.NoError(t, err)
}

func TestEmbeddingService_ValidateEmbedding_InvalidVectorSize(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEmbeddingRepository(ctrl)
	service := NewEmbeddingService(repo)

	ctx := context.Background()
	vector := make([]float32, 100) // Invalid size

	err := service.ValidateEmbedding(ctx, vector)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "vector size must be 512")
}

func TestEmbeddingService_DeleteEmbedding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEmbeddingRepository(ctrl)
	service := NewEmbeddingService(repo)

	ctx := context.Background()
	id := int64(123)

	repo.EXPECT().DeleteEmbeddingById(ctx, id).Return(nil)

	err := service.DeleteEmbedding(ctx, id)
	assert.NoError(t, err)
}
