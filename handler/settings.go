package handler

import (
	"github.com/TikhonP/ctg-medsenger-bot/db"
	"github.com/TikhonP/ctg-medsenger-bot/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SettingsHandler struct {
}

func (h *SettingsHandler) Handle(c echo.Context) error {
	contractId := util.QueryParamInt(c, "contract_id")
	if contractId == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "contract_id is required")
	}
	contract, err := db.GetContractById(*contractId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, contract)
}
