package movie

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// listMoviesHandler godoc
// @Summary list movies
// @Description list movies
// @Tags movies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param limit query int true "limit"
// @Param offset query int true "offset"
// @Param order query string false "order"
// @Success 200 {object} getMoviesResponse "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /movies [get]
func (c *Controller) listMoviesHandler(w http.ResponseWriter, r *http.Request) {
	order := r.URL.Query().Get("order")

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

	movies, err := c.moviesService.OrderMovies(r.Context(), order, limit, offset)
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
