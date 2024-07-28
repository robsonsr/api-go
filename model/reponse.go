package model

type ResponseError struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func NewInternalServerError(details string) ResponseError {
	reponseError := ResponseError{
		Message: "Erro no servidor.",
		Details: details,
	}
	return reponseError
}

func NewBadRequestError(details string) ResponseError {
	reponseError := ResponseError{
		Message: "Os dados enviados são inválidos.",
		Details: details,
	}
	return reponseError
}
