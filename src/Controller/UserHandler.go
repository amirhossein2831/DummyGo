package Controller

import (
	"github.com/amirhossein2831/DummyGo/src/App"
	"github.com/amirhossein2831/DummyGo/src/Model"
	"github.com/amirhossein2831/DummyGo/src/Request"
	"github.com/amirhossein2831/DummyGo/src/Response"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
	users, err := Model.GetAllUser()
	if err != nil {
		Response.ToJson(w, 400, err)
		return
	}
	Response.ToJson(w, 200, struct {
		Users []Model.User `json:"users"`
	}{Users: users})
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	user, err := Model.GetUser(userID)
	if err != nil {
		Response.ToJson(w, 400, err)
		return
	}
	Response.ToJson(w, 200, user)
}

func StoreUser(w http.ResponseWriter, r *http.Request) {
	user := Model.User{}
	Request.ToStruct(w, r, &user)

	res := App.App.DB.Create(&user)
	if res.Error != nil {
		Response.ToJson(w, 400, res.Error)
		return
	}
	Response.ToJson(w, 200, user)
}
