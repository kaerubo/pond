package usecase_test

import (
	"context"
	"github.com/kaerubo/pond/internal/entity"
	"github.com/kaerubo/pond/internal/test/mock"
	"github.com/kaerubo/pond/internal/usecase"
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

			mockSaver := mock.NewMockKeroSaver(ctrl)
			c := context.Background()

			if !tt.wantErr {
				mockSaver.
					EXPECT().
					Save(c, tt.kero).
					Times(1)
			}

			creator := usecase.NewKeroCreator(mockSaver)
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
