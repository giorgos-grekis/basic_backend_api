package user

import (
	"net/http"

	"github.com/giorgos-grekis/basic_backend_api/types"
	"github.com/giorgos-grekis/basic_backend_api/utils"
	"github.com/gorilla/mux"
)

// type Handler struct {
// 	store types.IUserStore
// }

//	func NewHandler(store types.IUserStore) *Handler {
//		return &Handler{store: store}
//	}
type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	if r.Body == nil {

	}
	var payload types.RegisterUserPayload
	// err := json.NewDecoder(r.Body).Decode(payload)
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// check if user exists

	// if not, create the new user
}
