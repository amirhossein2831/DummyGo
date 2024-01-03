package Controller

import (
	"github.com/amirhossein2831/DummyGo/src/App"
	"github.com/amirhossein2831/DummyGo/src/Model"
	"github.com/amirhossein2831/DummyGo/src/Request"
	"github.com/amirhossein2831/DummyGo/src/Response"
	"net/http"
)

func index() {

}

func show() {

}

func StoreUser(w http.ResponseWriter, r *http.Request) {
	user := Model.User{}
	Request.ToStruct(w, r, &user)

	res := App.App.DB.Create(&user)
	if res.Error != nil {
		Response.ToJson(w, 400, res.Error)
	}
	Response.ToJson(w, 200, user)
}
