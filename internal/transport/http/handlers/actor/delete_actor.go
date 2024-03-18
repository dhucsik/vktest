package actor

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/vktest/internal/entity"
	"github.com/dhucsik/vktest/internal/enum"
)

// deleteActorHandler godoc
// @Summary Delete actor
// @Description delete actor
// @Tags actors
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "actor ID"
// @Success 200 {object} nil "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 403 {object} errorResponse "Forbidden"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /actors/{id} [delete]
func (c *Controller) deleteActorHandler(w http.ResponseWriter, r *http.Request) {
	session, ok := entity.GetSession(r.Context())
	if !ok {
		sendErrorResponse(w, "undefined session", http.StatusUnauthorized)
		return
	}

	if session.Role != enum.UserRoleAdmin {
		sendErrorResponse(w, "permission denied", http.StatusForbidden)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.actorService.DeleteActor(r.Context(), id)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
