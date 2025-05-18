package usecase_test

import (
	"context"
	"errors"
	"github.com/kaerubo/pond/internal/entity"
	"github.com/kaerubo/pond/internal/test/mock"
	"github.com/kaerubo/pond/internal/usecase"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestKeroReader_GetByID(t *testing.T) {
	tests := []struct {
		name      string
		inputID   string
		setupMock func(m *mock.MockKeroByIDFinder)
		wantErr   bool
	}{
		{
			name:    "empty id",
			inputID: "",
			setupMock: func(m *mock.MockKeroByIDFinder) {
				// No mock setup needed for this case
			},
			wantErr: true,
		},
		{
			name:    "found",
			inputID: "abc-123",
			setupMock: func(m *mock.MockKeroByIDFinder) {
				m.EXPECT().
					FindByID(gomock.Any(), "abc-123").
					Return(&entity.Kero{ID: "abc-123"}, nil)
			},
			wantErr: false,
		},
		{
			name:    "not found",
			inputID: "not-found",
			setupMock: func(m *mock.MockKeroByIDFinder) {
				m.EXPECT().
					FindByID(gomock.Any(), "not-found").
					Return(nil, errors.New("not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFinder := mock.NewMockKeroByIDFinder(ctrl)
			tt.setupMock(mockFinder)

			reader := usecase.NewKeroByIDReader(mockFinder)
			_, err := reader.GetByID(context.Background(), tt.inputID)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
