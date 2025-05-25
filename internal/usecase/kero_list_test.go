package usecase_test

import (
	"context"
	"errors"
	"github.com/kaerubo/kaeruashi/internal/entity"
	"github.com/kaerubo/kaeruashi/internal/repository/mock"
	"github.com/kaerubo/kaeruashi/internal/usecase"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestKeroLister_List(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mock.MockKeroFinder)
		wantErr   bool
	}{
		{
			name: "success",
			setupMock: func(m *mock.MockKeroFinder) {
				m.EXPECT().
					FindAll(gomock.Any()).
					Return([]*entity.Kero{
						{ID: "1", Title: "K1"},
						{ID: "2", Title: "K2"},
					}, nil)
			},
			wantErr: false,
		},
		{
			name: "repo error",
			setupMock: func(m *mock.MockKeroFinder) {
				m.EXPECT().
					FindAll(gomock.Any()).
					Return(nil, errors.New("DB error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFinder := mock.NewMockKeroFinder(ctrl)
			tt.setupMock(mockFinder)

			lister := usecase.NewKeroLister(mockFinder)
			_, err := lister.List(context.Background())

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
