package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"faroukelkholy/survey/internal/service/survey"
)

//Server struct holds echo instance that holds all mounted routes
type Server struct {
	echo *echo.Echo
}

//New Return Server instance
func New() *Server {
	srv := new(Server)
	srv.echo = echo.New()
	srv.init()

	return srv
}

//Start bind echo to a port and run echo server
func (srv *Server) Start(port string) error {
	return srv.echo.Start(fmt.Sprintf(":%s", port))
}

//init holds all middleware that can be useful
func (srv *Server) init() {
	srv.echo.Use(middleware.Logger())
}

//AddSurveyFormRoutes mount routes related to survey service
func (srv *Server) AddSurveyFormRoutes(s survey.Service) {
	srv.echo.GET("/survey_forms/:id/surveys", GSurveysBySfIDHandler(s))
	srv.echo.POST("/survey_forms/:id/surveys", PSurveyHandler(s))
	srv.echo.POST("/survey_forms", PSurveyFormHandler(s))
}
