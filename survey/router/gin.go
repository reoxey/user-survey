package router

import (
	"github.com/gin-gonic/gin"
	"survey/api"
	"survey/core"
	"survey/logger"
)

type Gin struct {
	*gin.Engine
	log *logger.Logger
}

func New(l *logger.Logger, debug bool) *Gin {
	g := gin.New()
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	return &Gin{
		g,
		l,
	}
}

func (g *Gin) Handle(s core.Service) {

	h := api.NewHandler(s, g.log)

	group := g.Group("/api")

	group.GET("/surveys", h.GetAllSurveys)
	group.GET("/surveys/:sid", h.GetSurvey)

	group.POST("/surveys", h.CreateSurvey)

	group.GET("/surveys/:sid/results", h.GetAllResults)
	group.GET("/surveys/:sid/results/:rid", h.GetResult)

	group.POST("/surveys/:sid/results", h.AddResult)
}
