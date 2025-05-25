package handler_test

import (
	"errors"
	"github.com/kaerubo/kaeruashi/internal/handler"
	"github.com/kaerubo/kaeruashi/internal/usecase/mock"
	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteKero(t *testing.T) {
	e := echo.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDeleter := mock.NewMockKeroDeleter(ctrl)

	h := handler.NewKeroHandler(nil, nil, nil, nil, mockDeleter)

	tests := []struct {
		name       string
		keroID     string
		mockSetup  func()
		wantStatus int
		wantErrMsg string
	}{
		{
			name:       "missing id",
			keroID:     "",
			mockSetup:  func() {},
			wantStatus: http.StatusBadRequest,
			wantErrMsg: "id is required",
		},
		{
			name:   "use case returns error",
			keroID: "some-id",
			mockSetup: func() {
				mockDeleter.
					EXPECT().
					Delete(gomock.Any(), "some-id").
					Return(errors.New("delete error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantErrMsg: "delete error",
		},
		{
			name:   "success",
			keroID: "valid-id",
			mockSetup: func() {
				mockDeleter.
					EXPECT().
					Delete(gomock.Any(), "valid-id").
					Return(nil)
			},
			wantStatus: http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/keros/"+tt.keroID, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.keroID)

			tt.mockSetup()

			err := h.DeleteKero(c)
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
