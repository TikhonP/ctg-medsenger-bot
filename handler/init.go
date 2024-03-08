package handler

import (
	"github.com/TikhonP/ctg-medsenger-bot/db"
	"github.com/TikhonP/maigo"
	"github.com/labstack/echo/v4"
	"net/http"
)

type contractIdModel struct {
	ContractId int `json:"contract_id" validate:"required"`
}

type InitHandler struct {
	MAIClient *maigo.Client
}

func (h *InitHandler) Handle(c echo.Context) error {
	m := new(contractIdModel)
	if err := c.Bind(m); err != nil {
		return err
	}
	if err := c.Validate(m); err != nil {
		return err
	}
	ci, err := h.MAIClient.GetContractInfo(m.ContractId)
	if err != nil {
		return err
	}
	agentToken, err := h.MAIClient.GetAgentTokenForContractId(m.ContractId)
	if err != nil {
		return err
	}
	contract := &db.Contract{
		Id:           ci.Id,
		IsActive:     true,
		AgentToken:   &agentToken.Token,
		PatientName:  &ci.PatientName,
		PatientEmail: &ci.PatientEmail,
	}
	if err := contract.Save(); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}
