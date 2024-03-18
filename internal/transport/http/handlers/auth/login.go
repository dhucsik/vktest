package auth

import (
	"encoding/json"
	"net/http"
)

// loginHandler godoc
// @Summary login
// @Description login
// @Tags auth
// @Accept json
// @Produce json
// @Param request body authRequest true "request"
// @Success 200 {object} authResponse "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /auth/login [post]
func (c *Controller) loginHandler(w http.ResponseWriter, r *http.Request) {
	var req authRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	access, refresh, err := c.usersService.Auth(r.Context(), req.Username, req.Password)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(authResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	})
}

func sendErrorResponse(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(errorResponse{Error: msg})
	if err != nil {
		http.Error(w, err.Error(), status)
	}
}
