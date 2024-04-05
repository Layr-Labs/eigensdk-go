package types

import (
	"github.com/Layr-Labs/eigensdk-go/utils"
)

// OperatorMetadata is the metadata operator uploads while registering
// itself to EigenLayer
type OperatorMetadata struct {

	// Name of the operator
	// It is a required field
	Name string `yaml:"name" json:"name"`

	// Website of the operator
	// It is a required field
	Website string `yaml:"website" json:"website"`

	// Description of the operator. There is a 200-character limit
	// It is a required field
	Description string `yaml:"description" json:"description"`

	// Logo of the operator. This should be a link to a image file
	// which is publicly accessible
	// It is a required field
	Logo string `yaml:"logo" json:"logo"`

	// Twitter handle of the operator
	// It is an optional field
	Twitter string `yaml:"twitter" json:"twitter"`
}

func (om *OperatorMetadata) Validate() error {
	err := utils.ValidateText(om.Name)
	if err != nil {
		return WrapError(ErrInvalidName, err)
	}

	err = utils.ValidateText(om.Description)
	if err != nil {
		return WrapError(ErrInvalidDescription, err)
	}

	if len(om.Logo) == 0 {
		return ErrLogoRequired
	}

	if err = utils.IsImageURL(om.Logo); err != nil {
		return err
	}

	if len(om.Website) != 0 {
		err = utils.CheckIfUrlIsValid(om.Website)
		if err != nil {
			return WrapError(ErrInvalidWebsiteUrl, err)
		}
	}

	if len(om.Twitter) != 0 {
		err := utils.CheckIfValidTwitterURL(om.Twitter)
		if err != nil {
			return WrapError(ErrInvalidTwitterUrl, err)
		}
	}

	return nil
}
