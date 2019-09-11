package errorhandler

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/pivotal-cf/go-pivnet/v2"
	"github.com/pivotal-cf/pivnet-cli/printer"
	"github.com/pivotal-cf/pivnet-cli/ui"
)

var (
	ErrAlreadyHandled = errors.New("error already handled")
	RedFunc           = ui.ErrorColor.SprintFunc()
)

//go:generate counterfeiter . ErrorHandler

type ErrorHandler interface {
	HandleError(err error) error
}

type errorHandler struct {
	format       string
	outputWriter io.Writer
	logWriter    io.Writer
}

func NewErrorHandler(
	format string,
	outputWriter io.Writer,
	logWriter io.Writer,
) ErrorHandler {
	return &errorHandler{
		format:       format,
		outputWriter: outputWriter,
		logWriter:    logWriter,
	}
}

func (h errorHandler) HandleError(err error) error {
	if err == nil {
		return nil
	}

	var message string

	switch err.(type) {
	case pivnet.ErrUnauthorized:
		message = fmt.Sprintf("Failed to authenticate - please provide valid API token")
	case pivnet.ErrNotFound:
		message = fmt.Sprintf("Pivnet error: %s", err.Error())
	case pivnet.ErrPivnetOther:
		e := err.(pivnet.ErrPivnetOther)

		var errorMessages []string
		for _, pivErr := range e.Errors {
			errorMessages = append(errorMessages, fmt.Sprintln("- ", pivErr))
		}

		message = fmt.Sprintf(
			"Pivnet returned %d - %s.%s%s",
			e.ResponseCode,
			e.Message,
			fmt.Sprintln(),
			strings.Join(errorMessages, ""),
		)
	default:
		message = err.Error()
	}

	coloredMessage := fmt.Sprintf(RedFunc(message))

	switch h.format {
	case printer.PrintAsJSON:
		_ = h.printLogln(coloredMessage)

		return ErrAlreadyHandled

	case printer.PrintAsYAML:
		_ = h.printLogln(coloredMessage)

		return ErrAlreadyHandled

	default:
		h.println(coloredMessage)
		return ErrAlreadyHandled
	}
}

func (h errorHandler) println(message string) error {
	_, err := h.outputWriter.Write([]byte(fmt.Sprintln(message)))
	return err
}

func (h errorHandler) printLogln(message string) error {
	_, err := h.logWriter.Write([]byte(fmt.Sprintln(message)))
	return err
}
