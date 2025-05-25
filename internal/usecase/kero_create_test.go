package usecase_test

import (
	"context"
	"github.com/kaerubo/kaeruashi/internal/entity"
	"github.com/kaerubo/kaeruashi/internal/repository/mock"
	"github.com/kaerubo/kaeruashi/internal/usecase"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestKeroCreator_Create(t *testing.T) {
	tests := []struct {
		name    string
		kero    *entity.Kero
		wantErr bool
	}{
		{
			name: "missing title",
			kero: &entity.Kero{
				Title:   "",
				Content: "valid content",
			},
			wantErr: true,
		},
		{
			name: "missing content",
			kero: &entity.Kero{
				Title:   "valid title",
				Content: "",
			},
			wantErr: true,
		},
		{
			name: "valid kero",
			kero: &entity.Kero{
				Title:   "Hello",
				Content: "This is valid",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockInserter := mock.NewMockKeroInserter(ctrl)
			c := context.Background()

			if !tt.wantErr {
				mockInserter.
					EXPECT().
					Insert(c, tt.kero).
					Times(1)
			}

			creator := usecase.NewKeroCreator(mockInserter)
			err := creator.Create(c, tt.kero)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
