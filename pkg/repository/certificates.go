package repository

import (
	"github.com/unerror/waffy/pkg/data"
	"github.com/unerror/waffy/pkg/services/protos/certificates"
)

const (
	// CertificateBucket is the Bucket Store that cerificates are stored in
	CertificateBucket = "certificates"
)

// CreateCertificate creates a Certificate in the data store
func CreateCertificate(d data.Store, c *certificates.Certificate) error {
	b, err := d.Bucket(CertificateBucket)
	if err != nil {
		return err
	}

	return Create(b, []byte(c.SerialNumber), c)
}
