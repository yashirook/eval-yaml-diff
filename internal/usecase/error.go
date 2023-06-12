package usecase

import "errors"

var (
	DifferentDocumentNumberError = errors.New("DifferentDocumentNumberError: ")
	DeniedDiffExistError         = errors.New("DeniedDiffExist: ")
)
