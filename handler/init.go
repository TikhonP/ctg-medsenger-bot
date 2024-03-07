package handler

import (
	"github.com/TikhonP/ctg-medsenger-bot/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

type contractIdModel struct {
	ContractId int `json:"contract_id" validate:"required"`
}

type InitHandler struct{}

func (h *InitHandler) Handle(c echo.Context) error {
	m := new(contractIdModel)
	if err := c.Bind(m); err != nil {
		return err
	}
	if err := c.Validate(m); err != nil {
		return err
	}
	if err := db.UpsetContract(m.ContractId); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}
