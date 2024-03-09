package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ctgNotificationModel struct {
	Id        uuid.UUID `json:"uuid" validate:"required"`
	PatientId string    `json:"id"`
	ClinicId  string    `json:"idClinic"`
}

type CtgNotificationHandler struct {
}

func (h *CtgNotificationHandler) Handle(c echo.Context) error {
	m := new(ctgNotificationModel)
	if err := c.Bind(m); err != nil {
		return err
	}
	if err := c.Validate(m); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
