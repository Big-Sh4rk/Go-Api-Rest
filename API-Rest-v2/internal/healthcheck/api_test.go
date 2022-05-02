package healthcheck

import (
	"github.com/Matias-Guevara/API-Rest-v2/internal/test"
	"github.com/Matias-Guevara/API-Rest-v2/pkg/log"
	"net/http"
	"testing"
)

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)
	RegisterHandlers(router, "0.9.0")
	test.Endpoint(t, router, test.APITestCase{
		"ok", "GET", "/healthcheck", "", nil, http.StatusOK, `"OK 0.9.0"`,
	})
}
