package movie

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dhucsik/vktest/internal/entity"
	"github.com/dhucsik/vktest/internal/enum"
)

// updateMovieHandler godoc
// @Summary Update movie
// @Description update movie
// @Tags movies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "movie ID"
// @Param request body updateMovieRequest true "request"
// @Success 200 {object} nil "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 403 {object} errorResponse "Forbidden"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /movies/{id} [put]
func (c *Controller) updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	session, ok := entity.GetSession(r.Context())
	if !ok {
		sendErrorResponse(w, "undefined session", http.StatusUnauthorized)
		return
	}

	if session.Role != enum.UserRoleAdmin {
		sendErrorResponse(w, "permission denied", http.StatusForbidden)
		return
	}

	var req updateMovieRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	params, err := req.convert()
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.moviesService.UpdateMovie(r.Context(), id, params); err != nil {
		sendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
