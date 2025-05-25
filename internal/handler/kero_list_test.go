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

func TestListKeros(t *testing.T) {
	e := echo.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLister := mock.NewMockKeroLister(ctrl)

	h := handler.NewKeroHandler(nil, nil, mockLister, nil, nil)

	tests := []struct {
		name       string
		mockSetup  func()
		wantStatus int
		wantResp   []*entity.Kero
		wantErrMsg string
	}{
		{
			name: "use case returns error",
			mockSetup: func() {
				mockLister.
					EXPECT().
					List(gomock.Any()).
					Return(nil, errors.New("list error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantErrMsg: "list error",
		},
		{
			name: "success",
			mockSetup: func() {
				mockLister.
					EXPECT().
					List(gomock.Any()).
					Return([]*entity.Kero{
						{
							ID:      "id1",
							Title:   "title1",
							Content: "content1",
						},
						{
							ID:      "id2",
							Title:   "title2",
							Content: "content2",
						},
					}, nil)
			},
			wantStatus: http.StatusOK,
			wantResp: []*entity.Kero{
				{
					ID:      "id1",
					Title:   "title1",
					Content: "content1",
				},
				{
					ID:      "id2",
					Title:   "title2",
					Content: "content2",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/keros", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.mockSetup()

			err := h.ListKeros(c)
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
				var resp []*entity.Kero
				if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
					t.Fatalf("failed to unmarshal: %v", err)
				}
				if len(resp) != len(tt.wantResp) {
					t.Fatalf("unexpected response length: got %d, want %d", len(resp), len(tt.wantResp))
				}
				for i := range resp {
					if resp[i].ID != tt.wantResp[i].ID || resp[i].Title != tt.wantResp[i].Title || resp[i].Content != tt.wantResp[i].Content {
						t.Errorf("unexpected response[%d]: got %+v, want %+v", i, resp[i], tt.wantResp[i])
					}
				}
			}
		})
	}
}
