package logout

import (
	"fmt"
	"io"

	"github.com/pivotal-cf/pivnet-cli/v3/errorhandler"
	"github.com/pivotal-cf/pivnet-cli/v3/printer"
	"github.com/pivotal-cf/pivnet-cli/v3/ui"
)

//go:generate counterfeiter . RCHandler
type RCHandler interface {
	RemoveProfileWithName(profileName string) error
}

type LogoutClient struct {
	rcHandler    RCHandler
	eh           errorhandler.ErrorHandler
	format       string
	outputWriter io.Writer
	printer      printer.Printer
}

func NewLogoutClient(
	rcHandler RCHandler,
	eh errorhandler.ErrorHandler,
	format string,
	outputWriter io.Writer,
	printer printer.Printer,
) *LogoutClient {
	return &LogoutClient{
		rcHandler:    rcHandler,
		eh:           eh,
		format:       format,
		outputWriter: outputWriter,
		printer:      printer,
	}
}

func (c *LogoutClient) Logout(profileName string) error {
	err := c.rcHandler.RemoveProfileWithName(profileName)
	if err != nil {
		return c.eh.HandleError(err)
	}

	return c.printLogout()
}

func (c *LogoutClient) printLogout() error {
	switch c.format {

	case printer.PrintAsTable:
		message := "Logged-out successfully"
		coloredMessage := ui.SuccessColor.SprintFunc()(message)

		_, err := fmt.Fprintln(c.outputWriter, coloredMessage)

		return err
	}

	return nil
}
