package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"faroukelkholy/survey/internal/service/models"
	"faroukelkholy/survey/internal/service/survey"
)

func PSurveyHandler(s survey.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var sv models.Survey
		if err = c.Bind(&sv); err != nil {
			log.Println("err bind survey ", err)
			return c.JSON(http.StatusBadRequest, HTTPResponse{
				Data: nil,
				Err: HTTPError{
					Title:       "survey data is not valid",
					Description: "",
				},
			})
		}

		err = s.SubmitSurvey(c.Request().Context(), &sv)
		if err != nil {
			log.Println("err execute service ", err)
			return c.JSON(http.StatusInternalServerError, HTTPResponse{
				Data: nil,
				Err: HTTPError{
					Title:       "internal error",
					Description: "",
				},
			})
		}

		return c.JSON(http.StatusCreated, HTTPResponse{
			Data: "created",
			Err:  HTTPError{},
		})
	}
}
