package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"faroukelkholy/survey/internal/service/survey"
)

func GSurveysBySfIDHandler(s survey.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		log.Println("id", c.Param("id"))
		result, err := s.GetSurveysBySurveyFormID(c.Request().Context(), c.Param("id"))
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

		if result == nil {
			return c.JSON(http.StatusNotFound, HTTPResponse{
				Data: nil,
				Err: HTTPError{
					Title:       "surveys not found",
					Description: "",
				},
			})
		}

		return c.JSON(http.StatusOK, HTTPResponse{
			Data: result,
			Err:  HTTPError{},
		})
	}
}
