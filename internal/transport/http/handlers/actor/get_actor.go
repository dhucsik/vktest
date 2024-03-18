package actor

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// getActorHandler godoc
// @Summary Get actor
// @Description get actor
// @Tags actors
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "actor ID"
// @Success 200 {object} getActorResponse "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 403 {object} errorResponse "Forbidden"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /actors/{id} [get]
func (c *Controller) getActorByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	actor, err := c.actorService.GetActor(r.Context(), id)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(convertActor(actor))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
