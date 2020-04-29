package handlers

import (
	"github.com/labstack/echo"
	"ncrypt-api/helpers"
	"ncrypt-api/models"
	"ncrypt-api/processors"
	"net/http"
)

func (di DI) PostStoreSecretV1(c echo.Context) error {
	payload := models.SecureMessageRequest{}

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			helpers.BuildResponse(
				http.StatusUnprocessableEntity,
				"body data not accepted",
				nil,
				nil,
				nil,
			),
		)
	}

	err = c.Validate(payload)
	if err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			helpers.BuildResponse(
				http.StatusUnprocessableEntity,
				"validation failed",
				nil,
				helpers.FormatValidationErrorMessage(err),
				nil,
			),
		)
	}

	messageUuid, err := processors.StoreMessage(di.RedisClient, di.Key, payload)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helpers.BuildResponse(
				http.StatusInternalServerError,
				"internal error occurred",
				nil,
				nil,
				nil,
			),
		)
	}

	return c.JSON(
		http.StatusCreated,
		helpers.BuildResponse(
			http.StatusCreated,
			"Message stored.",
			messageUuid,
			nil,
			nil,
		),
	)
}