package Controller

import (
	"github.com/amirhossein2831/DummyGo/src/Response"
	"net/http"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	homeDetail := struct {
		Status      int       `json:"status"`
		WebSiteName string    `json:"webSiteName"`
		Version     string    `json:"version"`
		Time        time.Time `json:"time"`
	}{
		Status:      200,
		WebSiteName: "DummyGo.com",
		Version:     "v1",
		Time:        time.Now(),
	}
	Response.ToJson(w, 200, homeDetail)
}
