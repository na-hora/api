package validators

import (
	"github.com/klassmann/cpfcnpj"
)

type DocumentValidatorInterface interface {
	ValidateCNPJ(cnpj string) bool
}

type documentValidator struct{}

func GetDocumentValidator() DocumentValidatorInterface {
	return &documentValidator{}
}

func (dv *documentValidator) ValidateCNPJ(cnpj string) bool {
	return cpfcnpj.ValidateCNPJ(cnpj)
}
