package v1

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"task-trail/internal/controller/http/v1/request"
	"task-trail/internal/controller/http/v1/response"
	"task-trail/internal/customerrors"
	"task-trail/internal/pkg/contextmanager"
	"task-trail/internal/pkg/storage"
	"task-trail/internal/usecase"

	"github.com/gin-gonic/gin"
)

var allowedMimeTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/webp": true,
}

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
	userID := r.contextmanager.GetUserID(c).(int)
	u, err := r.u.GetByID(c, userID)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, response.NewCurrentUser(u, r.storage))
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
	file, err := c.FormFile("file")
	if err != nil {
		_ = c.Error(r.errHandler.Validation(err))
		return
	}
	mimeType := file.Header.Get("Content-Type")
	if !allowedMimeTypes[mimeType] {
		_ = c.Error(r.errHandler.Validation(fmt.Errorf("invalid mime type: %s", mimeType)))
		return
	}
	filename := filepath.Base(file.Filename)
	buf := bytes.NewBuffer(nil)
	f, err := file.Open()
	if err != nil {
		_ = c.Error(r.errHandler.Validation(fmt.Errorf("cant read file: %w", err)))
		return
	}
	defer f.Close()

	if _, err = io.Copy(buf, f); err != nil {
		_ = c.Error(r.errHandler.InternalTrouble(err, "cant copy file"))
		return
	}

	userID := r.contextmanager.GetUserID(c).(int)

	avatarID, err := r.u.UpdateAvatar(c, userID, buf.Bytes(), filename, mimeType)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.NewUserAvatar(avatarID, r.storage))

}

// @Summary		update current user
// @Description ...
// @Security BearerAuth
// @Tags 		/v1/users
// @Accept 		json
// @Produce 	json
// @Param 		body body request.UpdateUser true "user data"
// @Success 	200
// @Failure		400 {object} response.ErrAPI "invalid request body"
// @Failure		404 {object} response.ErrAPI "user not found"
// @Failure		401 {object} response.ErrAPI "authentication required"
// @Router 		/v1/users/me [patch]
func (r *usersRoutes) updateMe(c *gin.Context) {
	var body request.UpdateUser
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		_ = c.Error(r.errHandler.Validation(err))
		return
	}
	data := body.FromUpdateUser()
	data.ID = r.contextmanager.GetUserID(c).(int)
	u, err := r.u.UpdateByID(c, data)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, response.NewCurrentUser(u, r.storage))
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
