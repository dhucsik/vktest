package actor

import (
	"encoding/json"
	"net/http"

	"github.com/dhucsik/vktest/internal/entity"
	"github.com/dhucsik/vktest/internal/enum"
)

// createActorHandler godoc
// @Summary Create actor
// @Description create actor
// @Tags actors
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param request body createActorRequest true "request"
// @Success 200 {object} createActorResponse "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 403 {object} errorResponse "Forbidden"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /actors [post]
func (c *Controller) createActorHandler(w http.ResponseWriter, r *http.Request) {
	session, ok := entity.GetSession(r.Context())
	if !ok {
		sendErrorResponse(w, "undefined session", http.StatusUnauthorized)
		return
	}

	if session.Role != enum.UserRoleAdmin {
		sendErrorResponse(w, "permission denied", http.StatusForbidden)
		return
	}

	var req createActorRequest
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

	id, err := c.actorService.CreateActor(r.Context(), model)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createActorResponse{ID: id})
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
