package utils

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// Извлечь dnQualifier из x509 сертификата.
func ExtractDnQualifier(cert []byte) (string, error) {
	// Декодировать PEM блок.
	pemBlock, _ := pem.Decode(cert)
	if pemBlock == nil { // Если PEM блока нет, вернуть ошибку.
		return "", errors.New("pem block is nil")
	}
	// Парсить PEM блок.
	x509Cert, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return "", err
	}
	// Пройти циклом по именам субъекта и сверить OID.
	var oid = []int{2, 5, 4, 46} // OID dnQualifier.
	for _, attr := range x509Cert.Subject.Names {
		if attr.Type.Equal(oid) { // Вернуть значение dnQualifier.
			return attr.Value.(string), nil
		}
	}
	return "", errors.New("dnQualifier not found")
}
