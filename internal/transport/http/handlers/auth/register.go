package auth

import (
	"encoding/json"
	"net/http"
)

// signupHandler godoc
// @Summary register
// @Description register
// @Tags auth
// @Accept json
// @Produce json
// @Param request body registerRequest true "request"
// @Success 200 {object} nil "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /auth/register [post]
func (c *Controller) signupHandler(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	model, err := req.convert()
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.usersService.CreateUser(r.Context(), model)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
