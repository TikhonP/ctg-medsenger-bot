package handler

import (
	"github.com/TikhonP/ctg-medsenger-bot/util"
	"github.com/TikhonP/ctg-medsenger-bot/view"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ConnectDeviceHandler struct {
}

func (h *ConnectDeviceHandler) Handle(c echo.Context) error {
	contractId := util.QueryParamInt(c, "contract_id")
	if contractId == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "contract_id is required")
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return view.ConnectDevice("лол").Render(c.Request().Context(), c.Response().Writer)
}
