package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pie-fire-dire/internal/api/beef/services"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBeefSummaryHandlers(t *testing.T) {
	t.Run("Beef summary handlers success", func(t *testing.T) {
		service := services.NewServices()

		str := "Fatback t-bone t-bone, pastrami t-bone. pork, meatloaf jowl enim. Bresaola t-bone."
		body := bytes.NewBuffer([]byte(str))

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", body)
		req.Header.Set(echo.HeaderContentType, "text/plain")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := NewHandlers(service)

		// Assertions
		if assert.NoError(t, h.PostSummary(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var response map[string]map[string]int
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)

			beefCounts, exists := response["beef"]
			if !exists {
				t.Fatal("Expected 'beef' key in response")
			}

			tboneCount, exists := beefCounts["t-bone"]
			if !exists {
				t.Fatal("Expected 't-bone' key in 'beef' counts")
			}

			assert.Equal(t, 4, tboneCount, "Expected 't-bone' count to be 4")
		}
	})
}
