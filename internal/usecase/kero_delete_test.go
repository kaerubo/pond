package usecase_test

import (
	"context"
	"errors"
	"github.com/kaerubo/pond/internal/test/mock"
	"github.com/kaerubo/pond/internal/usecase"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestKeroDeleter_Delete(t *testing.T) {
	tests := []struct {
		name      string
		inputID   string
		setupMock func(m *mock.MockKeroDeleter)
		wantErr   bool
	}{
		{
			name:    "empty id",
			inputID: "",
			setupMock: func(m *mock.MockKeroDeleter) {
			},
			wantErr: true,
		},
		{
			name:    "successful delete",
			inputID: "abc-123",
			setupMock: func(m *mock.MockKeroDeleter) {
				m.EXPECT().
					Delete(gomock.Any(), "abc-123").
					Return(nil)
			},
			wantErr: false,
		},
		{
			name:    "repo returns error",
			inputID: "abc-123",
			setupMock: func(m *mock.MockKeroDeleter) {
				m.EXPECT().
					Delete(gomock.Any(), "abc-123").
					Return(errors.New("delete failed"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mock.NewMockKeroDeleter(ctrl)
			tt.setupMock(mockRepo)

			deleter := usecase.NewKeroDeleter(mockRepo)
			err := deleter.Delete(context.Background(), tt.inputID)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
