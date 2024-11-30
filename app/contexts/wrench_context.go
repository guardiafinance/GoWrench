package contexts

import (
	"net/http"
)

type WrenchContext struct {
	ResponseWriter *http.ResponseWriter
}
