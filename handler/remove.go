package handler

import (
	"github.com/TikhonP/ctg-medsenger-bot/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RemoveHandler struct {
}

func (h *RemoveHandler) Handle(c echo.Context) error {
	m := new(contractIdModel)
	if err := c.Bind(m); err != nil {
		return err
	}
	if err := c.Validate(m); err != nil {
		return err
	}
	if err := db.MarkInactiveContractWithId(m.ContractId); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
