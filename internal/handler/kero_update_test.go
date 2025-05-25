package handler_test

import (
	"bytes"
	"errors"
	"github.com/kaerubo/kaeruashi/internal/handler"
	"github.com/kaerubo/kaeruashi/internal/usecase/mock"
	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateKero(t *testing.T) {
	e := echo.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUpdater := mock.NewMockKeroUpdater(ctrl)

	h := handler.NewKeroHandler(nil, nil, nil, mockUpdater, nil)

	tests := []struct {
		name        string
		keroID      string
		requestBody string
		mockSetup   func()
		wantStatus  int
		wantErrMsg  string
	}{
		{
			name:        "missing id",
			keroID:      "",
			requestBody: `{"title":"test","content":"test content"}`,
			mockSetup:   func() {},
			wantStatus:  http.StatusBadRequest,
			wantErrMsg:  "id is required",
		},
		{
			name:        "invalid JSON",
			keroID:      "valid-id",
			requestBody: `invalid-json`,
			mockSetup:   func() {},
			wantStatus:  http.StatusBadRequest,
			wantErrMsg:  "invalid request",
		},
		{
			name:        "missing title or content",
			keroID:      "valid-id",
			requestBody: `{"title":"","content":""}`,
			mockSetup:   func() {},
			wantStatus:  http.StatusBadRequest,
			wantErrMsg:  "title and content are required",
		},
		{
			name:        "usecase returns error",
			keroID:      "valid-id",
			requestBody: `{"title":"test","content":"test content"}`,
			mockSetup: func() {
				mockUpdater.
					EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(errors.New("update error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantErrMsg: "update error",
		},
		{
			name:        "success",
			keroID:      "valid-id",
			requestBody: `{"title":"updated","content":"updated content"}`,
			mockSetup: func() {
				mockUpdater.
					EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			wantStatus: http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPut, "/keros/"+tt.keroID, bytes.NewBufferString(tt.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.keroID)

			tt.mockSetup()

			err := h.UpdateKero(c)
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

			if rec.Code != tt.wantStatus {
				t.Errorf("unexpected status: got %d, want %d", rec.Code, tt.wantStatus)
			}

			if rec.Body.Len() != 0 && tt.wantStatus == http.StatusNoContent {
				t.Errorf("unexpected body for NoContent: %s", rec.Body.String())
			}
		})
	}
}
