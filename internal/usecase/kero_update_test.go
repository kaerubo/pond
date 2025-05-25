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

func TestKeroUpdater_Update(t *testing.T) {
	data := &entity.Kero{
		ID:      "1",
		Title:   "ok",
		Content: "ok",
	}
	tests := []struct {
		name      string
		input     *entity.Kero
		setupMock func(m *mock.MockKeroUpdater)
		wantErr   bool
	}{
		{
			name: "missing id",
			input: &entity.Kero{
				ID:      "",
				Title:   "valid",
				Content: "valid",
			},
			setupMock: func(m *mock.MockKeroUpdater) {},
			wantErr:   true,
		},
		{
			name: "missing title",
			input: &entity.Kero{
				ID:      "1",
				Title:   "",
				Content: "valid",
			},
			setupMock: func(m *mock.MockKeroUpdater) {},
			wantErr:   true,
		},
		{
			name: "missing content",
			input: &entity.Kero{
				ID:      "1",
				Title:   "valid",
				Content: "",
			},
			setupMock: func(m *mock.MockKeroUpdater) {},
			wantErr:   true,
		},
		{
			name:  "success",
			input: data,
			setupMock: func(m *mock.MockKeroUpdater) {
				m.EXPECT().
					Update(gomock.Any(), data).
					DoAndReturn(func(ctx context.Context, k *entity.Kero) error {
						if k.UpdatedAt.IsZero() {
							return errors.New("expected UpdatedAt to be set")
						}
						return nil
					})
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mock.NewMockKeroUpdater(ctrl)
			tt.setupMock(mockRepo)

			updater := usecase.NewKeroUpdater(mockRepo)
			err := updater.Update(context.Background(), tt.input)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
