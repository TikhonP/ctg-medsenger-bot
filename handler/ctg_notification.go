package handler

import (
	"github.com/TikhonP/ctg-medsenger-bot/util"
	"github.com/TikhonP/maigo"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type ctgNotificationModel struct {
	Id uuid.UUID `json:"uuid" validate:"required"`

	// PatientIdentifier is medsenger contract ID.
	// It's a value that user enters as identifier in CTG-app
	PatientIdentifier string `json:"id" validate:"required"`

	ClinicId string `json:"idClinic"`
}

type CtgNotificationHandler struct {
	MaigoClient *maigo.Client
	CtgClient   *util.CtgClient
}

func (h *CtgNotificationHandler) Handle(c echo.Context) error {
	m := new(ctgNotificationModel)
	if err := c.Bind(m); err != nil {
		return err
	}
	if err := c.Validate(m); err != nil {
		return err
	}
	contractId, err := strconv.Atoi(m.PatientIdentifier)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "`id` is invalid.")
	}
	log.Println("Got new notification: ", m)
	go h.SendPDFToMedsenger(m.Id, contractId)
	return c.NoContent(http.StatusOK)
}

func (h *CtgNotificationHandler) SendPDFToMedsenger(recordId uuid.UUID, contractId int) {
	pdfData, err := h.CtgClient.GetRecordPDF(recordId)
	if err != nil {
		sentry.CaptureException(err)
	}
	defer pdfData.Close()
	attachment, err := maigo.NewMessageAttachment("ctg_record.pdf", "application/pdf", pdfData)
	if err != nil {
		sentry.CaptureException(err)
	}
	_, err = h.MaigoClient.SendMessage(contractId, "", maigo.WithAttachment(attachment))
	if err != nil {
		sentry.CaptureException(err)
	}
}
