package application

import (
	"github.com/aprilnurf/grocerystore-oauth-api/clients/cassandra"
	"github.com/aprilnurf/grocerystore-oauth-api/src/domain/access_token"
	"github.com/aprilnurf/grocerystore-oauth-api/src/http"
	"github.com/aprilnurf/grocerystore-oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)
var (
	router = gin.Default()
)
func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()
	atHandler := http.NewHandler(access_token.NewService(db.New()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":9998")
}
