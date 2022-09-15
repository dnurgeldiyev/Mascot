package v1

import (
	"dovran/mascot/internal/usecase"
	"dovran/mascot/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, uc usecase.UseCase, l logger.Interface) {

	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/v1")
	{
		methods(h, uc, l)
	}

}
