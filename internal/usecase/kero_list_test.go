package usecase_test

import (
	"context"
	"errors"
	"github.com/kaerubo/kaeruashi/internal/entity"
	"github.com/kaerubo/kaeruashi/internal/test/mock"
	"github.com/kaerubo/kaeruashi/internal/usecase"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestKeroLister_List(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mock.MockKeroLister)
		wantErr   bool
	}{
		{
			name: "success",
			setupMock: func(m *mock.MockKeroLister) {
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
			setupMock: func(m *mock.MockKeroLister) {
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

			mockRepo := mock.NewMockKeroLister(ctrl)
			tt.setupMock(mockRepo)

			lister := usecase.NewKeroLister(mockRepo)
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
