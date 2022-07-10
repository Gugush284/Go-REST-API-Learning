package model_user

import validation "github.com/go-ozzo/ozzo-validation"

// Validation ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Login, validation.Required, validation.Length(4, 25)),
		validation.Field(&u.DecryptedPassword, validation.By(requiredIf(u.Password == "")), validation.Length(8, 100)),
	)
}

func requiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}

		return nil
	}
}
