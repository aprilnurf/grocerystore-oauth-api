package http

import (
	"github.com/aprilnurf/grocerystore-oauth-api/src/domain/access_token"
	"github.com/aprilnurf/grocerystore-oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func (h *accessTokenHandler) Create(ctx *gin.Context) {
	var at access_token.AccessToken
	if err := ctx.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		ctx.JSON(restErr.Status, restErr)
	}

	if err := h.service.Create(at); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusCreated, at)
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := h.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
