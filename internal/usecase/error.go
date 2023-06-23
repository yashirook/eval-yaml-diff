package usecase

import "errors"

var (
	DifferentDocumentNumberError = errors.New("different number of documents are not supported")
	DeniedDiffExistError         = errors.New("there are some differences that are not allowed by policy")
)
