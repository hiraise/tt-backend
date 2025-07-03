package v1

import (
	"net/http"
	"task-trail/internal/controller/http/v1/request"
	"task-trail/internal/controller/http/v1/response"
	"task-trail/internal/customerrors"
	"task-trail/internal/pkg/contextmanager"
	"task-trail/internal/pkg/storage"
	"task-trail/internal/usecase"
	"task-trail/internal/utils"

	"github.com/gin-gonic/gin"
)

type usersRoutes struct {
	contextmanager contextmanager.Gin
	errHandler     customerrors.ErrorHandler
	storage        storage.Service
	u              usecase.User
}

// @Summary 	return user by id
// @Description ...
// @Security BearerAuth
// @Tags 		/v1/users
// @Accept 		json
// @Produce 	json
// @Param 		id path int true "user id"
// @Success 	200
// @Router 		/v1/users/{id} [get]
func (r *usersRoutes) getUser(c *gin.Context) {
	// id := c.Param("id")
	c.JSON(http.StatusOK, c.Keys["userID"])
}

// @Summary 	return current user
// @Description ...
// @Security BearerAuth
// @Tags 		/v1/users
// @Accept 		json
// @Produce 	json
// @Success 	200
// @Router 		/v1/users/me [get]
func (r *usersRoutes) getMe(c *gin.Context) {
	userID := utils.Must(r.contextmanager.GetUserID(c))

	u, err := r.u.GetByID(c, userID)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, response.CurrentUserFromDTO(u))
}

// @Summary 	upload new avatar
// @Description ...
// @Security BearerAuth
// @Tags 		/v1/users
// @Accept 		json
// @Produce 	json
// @Param 		file formData file true "new file"
// @Success 	200
// @Failure		401 {object} response.ErrAPI "authentication required"
// @Router 		/v1/users/me/avatar [patch]
func (r *usersRoutes) updateAvatar(c *gin.Context) {
	// extract file
	file, err := c.FormFile("file")
	if err != nil {
		_ = c.Error(r.errHandler.Validation(err))
		return
	}
	// map file to FileReq
	f, err := request.FileFromAPI(file)
	if err != nil {
		_ = c.Error(r.errHandler.InternalTrouble(err, "file mapping failure"))
		return
	}

	userID := utils.Must(r.contextmanager.GetUserID(c))

	avatarID, err := r.u.UpdateAvatar(c, f.ToDTO(userID))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.AvatarToAPI(avatarID, r.storage))

}

// @Summary		update current user
// @Description ...
// @Security BearerAuth
// @Tags 		/v1/users
// @Accept 		json
// @Produce 	json
// @Param 		body body request.UpdateReq true "user data"
// @Success 	200
// @Failure		400 {object} response.ErrAPI "invalid request body"
// @Failure		404 {object} response.ErrAPI "user not found"
// @Failure		401 {object} response.ErrAPI "authentication required"
// @Router 		/v1/users/me [patch]
func (r *usersRoutes) updateMe(c *gin.Context) {
	// parse body
	var body request.UpdateReq
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		_ = c.Error(r.errHandler.Validation(err))
		return
	}
	// cast parsed body to entity
	userID := utils.Must(r.contextmanager.GetUserID(c))
	data := body.ToDTO(userID)

	u, err := r.u.UpdateByID(c, data)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, response.CurrentUserFromDTO(u))
}
func NewUserRouter(
	router *gin.RouterGroup,
	u usecase.User,
	authMW gin.HandlerFunc,
	errHandler customerrors.ErrorHandler,
	contextmanager contextmanager.Gin,
	storage storage.Service,
) {
	r := &usersRoutes{u: u, contextmanager: contextmanager, errHandler: errHandler, storage: storage}
	g := router.Group("/users")
	g.PATCH("me/avatar", authMW, r.updateAvatar)
	g.GET(":id", authMW, r.getUser)
	g.GET("me", authMW, r.getMe)
	g.PATCH("me", authMW, r.updateMe)
}
