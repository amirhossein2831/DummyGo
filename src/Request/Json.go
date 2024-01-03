package Request

import (
	"encoding/json"
	"github.com/amirhossein2831/DummyGo/src/Response"
	"net/http"
)

func ToStruct(w http.ResponseWriter, r *http.Request, data interface{}) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&data)
	Response.JsonError(w, err)
}
