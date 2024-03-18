package auth

import (
	"encoding/json"
	"net/http"
)

// refreshHandler godoc
// @Summary refresh
// @Description refresh
// @Tags auth
// @Accept json
// @Produce json
// @Param request body refreshRequest true "request"
// @Success 200 {object} authResponse "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /auth/refresh [post]
func (c *Controller) refreshHandler(w http.ResponseWriter, r *http.Request) {
	var req refreshRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	access, refresh, err := c.usersService.Refresh(r.Context(), req.Token)
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
