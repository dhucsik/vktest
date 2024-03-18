package movie

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dhucsik/vktest/internal/entity"
)

// searchMoviesHandler godoc
// @Summary search movies
// @Description search movies
// @Tags movies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param limit query int true "limit"
// @Param offset query int true "offset"
// @Param title query string false "title"
// @Param actor query string false "actor"
// @Success 200 {object} getMoviesResponse "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /movies/search [get]
func (c *Controller) searchMoviesHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	actor := r.URL.Query().Get("actor")

	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	offsetStr := r.URL.Query().Get("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	var movies []*entity.MovieWithActors
	if title != "" {
		movies, err = c.moviesService.SearchByTitle(r.Context(), title, limit, offset)
	} else if actor != "" {
		movies, err = c.moviesService.SearchByActor(r.Context(), actor, limit, offset)
	} else {
		sendErrorResponse(w, "title or actor must be defined", http.StatusBadRequest)
		return
	}

	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(convertMovies(movies))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
