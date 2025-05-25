package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/kaerubo/kaeruashi/internal/entity"
	"github.com/kaerubo/kaeruashi/internal/handler"
	"github.com/kaerubo/kaeruashi/internal/usecase/mock"
	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateKero(t *testing.T) {
	e := echo.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCreator := mock.NewMockKeroCreator(ctrl)

	h := handler.NewKeroHandler(mockCreator, nil, nil, nil, nil)

	tests := []struct {
		name        string
		requestBody string
		mockSetup   func()
		wantStatus  int
		wantID      string
		wantErrMsg  string
	}{
		{
			name:        "invalid JSON",
			requestBody: `invalid-json`,
			mockSetup:   func() {},
			wantStatus:  http.StatusBadRequest,
			wantErrMsg:  "invalid request",
		},
		{
			name:        "missing title or content",
			requestBody: `{"title":"","content":""}`,
			mockSetup:   func() {},
			wantStatus:  http.StatusBadRequest,
			wantErrMsg:  "title and content are required",
		},
		{
			name:        "use case returns error",
			requestBody: `{"title":"test","content":"test content"}`,
			mockSetup: func() {
				mockCreator.
					EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(errors.New("db error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantErrMsg: "db error",
		},
		{
			name:        "success",
			requestBody: `{"title":"test","content":"test content"}`,
			mockSetup: func() {
				mockCreator.
					EXPECT().
					Create(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ interface{}, k *entity.Kero) error {
						k.ID = "generated-id"
						return nil
					})
			},
			wantStatus: http.StatusCreated,
			wantID:     "generated-id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/keros", bytes.NewBufferString(tt.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.mockSetup()

			err := h.CreateKero(c)
			if err != nil {
				var httpError *echo.HTTPError
				ok := errors.As(err, &httpError)
				if !ok {
					t.Fatalf("expected echo.HTTPError, got %T", err)
				}
				if httpError.Code != tt.wantStatus {
					t.Errorf("unexpected status: got %d, want %d", httpError.Code, tt.wantStatus)
				}
				if httpError.Message != tt.wantErrMsg {
					t.Errorf("unexpected error message: got %v, want %v", httpError.Message, tt.wantErrMsg)
				}
				return
			}

			if tt.wantID != "" {
				if rec.Code != tt.wantStatus {
					t.Errorf("unexpected status: got %d, want %d", rec.Code, tt.wantStatus)
				}
				var resp map[string]string
				if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
					t.Fatalf("failed to unmarshal: %v", err)
				}
				if resp["id"] != tt.wantID {
					t.Errorf("unexpected id: got %s, want %s", resp["id"], tt.wantID)
				}
			}
		})
	}
}
