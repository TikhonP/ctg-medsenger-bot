package server

import (
	"github.com/TikhonP/ctg-medsenger-bot/appconfig"
	"github.com/TikhonP/ctg-medsenger-bot/db"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var medsengerAgentKey string

type ContractIdModel struct {
	*ApiKey
	ContractId int `json:"contract_id"`
}

func initView(c echo.Context) error {
	cid := new(ContractIdModel)
	if bindErr := c.Bind(cid); bindErr != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}
	if dbErr := db.UpsetContract(cid.ContractId); dbErr != nil {
		return dbErr
	}
	return c.NoContent(http.StatusCreated)
}

func Serve(cfg *appconfig.Server) {
	medsengerAgentKey = cfg.MedsengerAgentKey

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Купил мужик шляпу, а она ему как раз!")
	})
	e.POST("/init", initView)

	addr := cfg.Host + ":" + strconv.Itoa(int(cfg.Port))
	err := e.Start(addr)
	if err != nil {
		panic(err)
	}
}
