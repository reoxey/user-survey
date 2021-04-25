package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"survey/core"
	"survey/logger"
)

// Handler interface called gin router for GET & POST calls
type Handler interface {
	CreateSurvey(c *gin.Context)
	GetAllSurveys(c *gin.Context)
	GetSurvey(c *gin.Context)

	AddResult(c *gin.Context)
	GetAllResults(c *gin.Context)
	GetResult(c *gin.Context)
}

// handle initialised with custon logger and
// service which will process the request and forward data to repo
type handle struct {
	service core.Service
	log *logger.Logger
}

func (h handle) CreateSurvey(c *gin.Context) {
	var s *core.Survey

	if err := c.Bind(&s); err != nil {
		h.log.Println("ERROR:handler.CreateSurvey", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	id, err := h.service.CreateSurvey(c, s)
	if err != nil {
		h.log.Println("ERROR:handler.CreateSurvey", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	c.Header("Location",
		fmt.Sprintf("/api/surveys/%s", id),
	)
	c.JSON(http.StatusCreated, nil)
}

func (h handle) GetAllSurveys(c *gin.Context) {
	surveys, err := h.service.GetAllSurveys(c)
	if err != nil {
		h.log.Println("ERROR:handler.GetAllSurvey", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, surveys)
}

func (h handle) GetSurvey(c *gin.Context) {
	survey, err := h.service.GetSurvey(c, c.Param("sid"))
	if err != nil {
		h.log.Println("ERROR:handler.GetSurvey", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, survey)
}

func (h handle) AddResult(c *gin.Context) {
	var r *core.Result

	if err := c.Bind(&r); err != nil {
		h.log.Println("ERROR:handler.AddResult", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	id, err := h.service.AddResult(c, r, c.Param("sid"))
	if err != nil {
		h.log.Println("ERROR:handler.AddResult", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	c.Header("Location",
		fmt.Sprintf("/api/surveys/%s/results/%s", r.SurveyId.Hex(), id),
	)
	c.JSON(http.StatusCreated, nil)
}

func (h handle) GetAllResults(c *gin.Context) {
	results, err := h.service.GetAllResultsBySurvey(c, c.Param("sid"))
	if err != nil {
		h.log.Println("ERROR:handler.GetAllResults", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, results)
}

func (h handle) GetResult(c *gin.Context) {
	result, err := h.service.GetResultById(c, c.Param("rid"))
	if err != nil {
		h.log.Println("ERROR:handler.GetResult", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, result)
}

func NewHandler(s core.Service, log *logger.Logger) Handler {
	return &handle{s, log}
}
