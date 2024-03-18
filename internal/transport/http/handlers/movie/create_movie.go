package movie

import (
	"encoding/json"
	"net/http"

	"github.com/dhucsik/vktest/internal/entity"
	"github.com/dhucsik/vktest/internal/enum"
)

// createMovieHandler godoc
// @Summary Create movie
// @Description create movie
// @Tags movies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param request body createMovieRequest true "request"
// @Success 200 {object} createMovieResponse "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 403 {object} errorResponse "Forbidden"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /movies [post]
func (c *Controller) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	session, ok := entity.GetSession(r.Context())
	if !ok {
		sendErrorResponse(w, "undefined session", http.StatusUnauthorized)
		return
	}

	if session.Role != enum.UserRoleAdmin {
		sendErrorResponse(w, "permission denied", http.StatusForbidden)
		return
	}

	var req createMovieRequest
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

	id, err := c.moviesService.CreateMovie(r.Context(), model, req.ActorIDs)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createMovieResponse{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func sendErrorResponse(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(errorResponse{Error: msg})
	if err != nil {
		http.Error(w, err.Error(), status)
	}
}
