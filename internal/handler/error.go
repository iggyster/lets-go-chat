package handler

type (
	Errors struct {
		Errors []*Error `json:"errors"`
	}

	Error struct {
		Field   string `json:"field"`
		Message string `json:"message"`
		Detail  string `json:"detail"`
	}
)

func (errors *Errors) AddError(field, message, detail string) {
	errors.Errors = append(errors.Errors, &Error{Field: field, Message: message, Detail: detail})
}

func (errors *Errors) Count() int {
	return len(errors.Errors)
}
