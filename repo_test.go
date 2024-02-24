package core_test

import (
	"context"
	"sqlc-tutorial/core"
	"sqlc-tutorial/internal/database"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockQueryRepo mocks the QueryRepo interface for testing.
type MockQueryRepo struct {
	mock.Mock
}

type MockQueries struct {
	mock.Mock
}

func (m *MockQueryRepo) WithTx(tx pgx.Tx) *db.Queries {
	// Previously returned nil or an improperly initialized mock
	mockQueries := &MockQueries{}
	m.On("CreateAuthor", mock.Anything, mock.Anything).Return(db.Author{}, nil) // Default return
	return mockQueries
}

func (m *MockQueries) CreateAuthor(ctx context.Context, arg db.CreateAuthorParams) (db.Author, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(db.Author), args.Error(1)
}

func (m *MockQueryRepo) GetAuthor(ctx context.Context, id int32) (db.Author, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(db.Author), args.Error(1)
}

// MockdbInterface mocks the dbInterface interface for testing.
type MockdbInterface struct {
	mock.Mock
}

func (m *MockdbInterface) Begin(ctx context.Context) (pgx.Tx, error) {
	args := m.Called(ctx)
	return nil, args.Error(1) // Return nil for Tx for simplicity, adjust as needed
}

// TestRepo_Create tests the Create method of the Repo.
func TestRepo_Create(t *testing.T) {
	ctx := context.Background()
	mockQueryRepo := new(MockQueryRepo)
	mockDbInterface := new(MockdbInterface)
	repo := core.NewRepo(mockQueryRepo, mockDbInterface)

	// Setup mock expectations
	expectedAuthor := db.Author{
		ID:   1,
		Name: "John Doe",
		Bio:  pgtype.Text{String: "Bio"},
	}
	mockQueryRepo.On("WithTx", mock.Anything).Return(new(db.Queries)) // Adjust to return a mock *db.Queries
	mockQueryRepo.On("CreateAuthor", ctx, mock.AnythingOfType("db.CreateAuthorParams")).Return(expectedAuthor, nil)
	mockQueryRepo.On("GetAuthor", ctx, int32(1)).Return(expectedAuthor, nil)
	mockDbInterface.On("Begin", ctx).Return(nil, nil) // Simulate successful transaction start

	// Call the method under test
	auth, err := repo.Create(ctx, core.Author{Name: "John Doe", Bio: pgtype.Text{String: "Bio"}})

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", auth.Name)
	mockQueryRepo.AssertExpectations(t) // Verify all expectations were met
	mockDbInterface.AssertExpectations(t)
}
