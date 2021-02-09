package commands

import "github.com/pivotal-cf/pivnet-cli/v3/commands/release"

type ReleasesCommand struct {
	ProductSlug string `long:"product-slug" short:"p" description:"Product slug e.g. p-mysql" required:"true"`
	Limit       string `long:"limit" short:"l" description:"Limit the number of returned releases to the most recent"`
}

type ReleaseCommand struct {
	ProductSlug    string `long:"product-slug" short:"p" description:"Product slug e.g. p-mysql" required:"true"`
	ReleaseVersion string `long:"release-version" short:"r" description:"Release version e.g. 0.1.2-rc1" required:"true"`
}

type UpdateReleaseCommand struct {
	ProductSlug    string  `long:"product-slug" short:"p" description:"Product slug e.g. p-mysql" required:"true"`
	ReleaseVersion string  `long:"release-version" short:"r" description:"Release version e.g. 0.1.2-rc1" required:"true"`
	Availability   *string `long:"availability" description:"Release availability. Optional." choice:"admins" choice:"selected-user-groups" choice:"all"`
	ReleaseType    *string `long:"release-type" description:"Release type. Optional." choice:"all-in-one" choice:"major" choice:"minor" choice:"service" choice:"maintenance" choice:"security" choice:"alpha" choice:"beta" choice:"edge"`
}

type DeleteReleaseCommand struct {
	ProductSlug    string `long:"product-slug" short:"p" description:"Product slug e.g. p-mysql" required:"true"`
	ReleaseVersion string `long:"release-version" short:"r" description:"Release version e.g. 0.1.2-rc1" required:"true"`
}

type CreateReleaseCommand struct {
	ProductSlug    string `long:"product-slug" short:"p" description:"Product slug e.g. p-mysql" required:"true"`
	ReleaseVersion string `long:"release-version" short:"r" description:"Release version e.g. 0.1.2-rc1" required:"true"`
	ReleaseType    string `long:"release-type" short:"t" description:"Release type e.g. 'Minor Release'" required:"true"`
	EULASlug       string `long:"eula-slug" short:"e" description:"EULA slug e.g. pivotal_software_eula" required:"true"`
}

//go:generate counterfeiter . ReleaseClient
type ReleaseClient interface {
	List(productSlug string) error
	ListWithLimit(productSlug string, limit string) error
	Get(productSlug string, releaseVersion string) error
	Create(productSlug string, releaseVersion string, releaseType string, eulaSlug string) error
	Update(productSlug string, releaseVersion string, availability *string, releaseType *string) error
	Delete(productSlug string, releaseVersion string) error
}

var NewReleaseClient = func(client release.PivnetClient) ReleaseClient {
	return release.NewReleaseClient(
		client,
		ErrorHandler,
		Pivnet.Format,
		OutputWriter,
		Printer,
	)
}

func (command *ReleasesCommand) Execute([]string) error {
	err := Init(true)
	if err != nil {
		return err
	}

	client := NewPivnetClient()
	err = Auth.AuthenticateClient(client)
	if err != nil {
		return err
	}

	if command.Limit != "" {
		return NewReleaseClient(client).ListWithLimit(command.ProductSlug, command.Limit)
	}
	return NewReleaseClient(client).List(command.ProductSlug)
}

func (command *ReleaseCommand) Execute([]string) error {
	err := Init(true)
	if err != nil {
		return err
	}

	client := NewPivnetClient()
	err = Auth.AuthenticateClient(client)
	if err != nil {
		return err
	}

	return NewReleaseClient(client).Get(command.ProductSlug, command.ReleaseVersion)
}

func (command *CreateReleaseCommand) Execute([]string) error {
	err := Init(true)
	if err != nil {
		return err
	}

	client := NewPivnetClient()
	err = Auth.AuthenticateClient(client)
	if err != nil {
		return err
	}

	return NewReleaseClient(client).Create(
		command.ProductSlug,
		command.ReleaseVersion,
		command.ReleaseType,
		command.EULASlug,
	)
}

func (command *UpdateReleaseCommand) Execute([]string) error {
	err := Init(true)
	if err != nil {
		return err
	}

	client := NewPivnetClient()
	err = Auth.AuthenticateClient(client)
	if err != nil {
		return err
	}

	return NewReleaseClient(client).Update(
		command.ProductSlug,
		command.ReleaseVersion,
		command.Availability,
		command.ReleaseType,
	)
}

func (command *DeleteReleaseCommand) Execute([]string) error {
	err := Init(true)
	if err != nil {
		return err
	}

	client := NewPivnetClient()
	err = Auth.AuthenticateClient(client)
	if err != nil {
		return err
	}

	return NewReleaseClient(client).Delete(command.ProductSlug, command.ReleaseVersion)
}
