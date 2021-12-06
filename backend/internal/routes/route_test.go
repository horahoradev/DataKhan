package routes

import (
	"encoding/json"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
	dkmetrics "github.com/horahoradev/DataKhan/backend/internal/metrics"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestView(t *testing.T) {
	tests := []struct {
		IP        string `json:"ip"`
		Useragent string `json:"useragent"`
		URI       string `json:"uri"`
	}{
		{
			IP:        "127.0.0.1",
			Useragent: "whatever",
			URI:       "https://mywebsite.com/datakhan",
		},
	}

	for _, test := range tests {
		e := echo.New()
		json, err := json.Marshal(test)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/view/", strings.NewReader(string(json)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		r := NewRouteHandler(dkmetrics.MockCounter)

		assert.NoError(t, r.handleView(c))
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
