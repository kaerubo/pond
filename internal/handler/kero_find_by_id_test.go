package handler_test

import (
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

func TestFindKeroByID(t *testing.T) {
	e := echo.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockFinder := mock.NewMockKeroByIDFinder(ctrl)

	h := handler.NewKeroHandler(nil, mockFinder, nil, nil, nil)

	tests := []struct {
		name       string
		keroID     string
		mockSetup  func()
		wantStatus int
		wantResp   *entity.Kero
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
			name:   "not found",
			keroID: "not-found-id",
			mockSetup: func() {
				mockFinder.
					EXPECT().
					FindByID(gomock.Any(), "not-found-id").
					Return(nil, nil)
			},
			wantStatus: http.StatusNotFound,
			wantErrMsg: "kero not found",
		},
		{
			name:   "use case error",
			keroID: "error-id",
			mockSetup: func() {
				mockFinder.
					EXPECT().
					FindByID(gomock.Any(), "error-id").
					Return(nil, errors.New("db error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantErrMsg: "db error",
		},
		{
			name:   "success",
			keroID: "valid-id",
			mockSetup: func() {
				mockFinder.
					EXPECT().
					FindByID(gomock.Any(), "valid-id").
					Return(&entity.Kero{
						ID:      "valid-id",
						Title:   "title",
						Content: "content",
					}, nil)
			},
			wantStatus: http.StatusOK,
			wantResp: &entity.Kero{
				ID:      "valid-id",
				Title:   "title",
				Content: "content",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/keros/"+tt.keroID, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.keroID)

			tt.mockSetup()

			err := h.FindKeroByID(c)
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

			if tt.wantResp != nil {
				var resp entity.Kero
				if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
					t.Fatalf("failed to unmarshal: %v", err)
				}
				if resp.ID != tt.wantResp.ID || resp.Title != tt.wantResp.Title || resp.Content != tt.wantResp.Content {
					t.Errorf("unexpected response: got %+v, want %+v", resp, tt.wantResp)
				}
			}
		})
	}
}
