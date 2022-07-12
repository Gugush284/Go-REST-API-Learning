package ModelImage

import validation "github.com/go-ozzo/ozzo-validation"

// Validation ...
func (i *Image) Validate() error {
	return validation.ValidateStruct(
		i,
		validation.Field(&i.ImageName, validation.Required),
		validation.Field(&i.ImageType, validation.Required),
		validation.Field(&i.Image, validation.Required),
	)
}
