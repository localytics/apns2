package certificate_test

import (
	"crypto/tls"
	"errors"
	"github.com/sideshow/apns2/certificate"
	"testing"
)

// p12

func TestValidCertificateFromP12File(t *testing.T) {
	cer, err := certificate.FromP12File("_fixtures/certificate-valid.p12", "")
	if err != nil {
		t.Fatal(err)
	}
	if e := verifyHostname(cer); e != nil {
		t.Fatal(e)
	}
}

func TestEncryptedValidCertificateFromP12File(t *testing.T) {
	cer, err := certificate.FromP12File("_fixtures/certificate-valid-encrypted.p12", "password")
	if err != nil {
		t.Fatal(err)
	}
	if e := verifyHostname(cer); e != nil {
		t.Fatal(e)
	}
}

func TestNoSuchFileP12File(t *testing.T) {
	_, err := certificate.FromP12File("", "")
	if err.Error() != errors.New("open : no such file or directory").Error() {
		t.Fatal("expected error", "open : no such file or directory")
	}
}

func TestBadPasswordP12File(t *testing.T) {
	_, err := certificate.FromP12File("_fixtures/certificate-valid-encrypted.p12", "")
	if err.Error() != errors.New("pkcs12: decryption password incorrect").Error() {
		t.Fatal("expected", "pkcs12: decryption password incorrect")
	}
}

// pem

func TestValidCertificateFromPemFile(t *testing.T) {
	cer, err := certificate.FromPemFile("_fixtures/certificate-valid.pem", "")
	if err != nil {
		t.Fatal(err)
	}
	if e := verifyHostname(cer); e != nil {
		t.Fatal(e)
	}
}

func TestEncryptedValidCertificateFromPemFile(t *testing.T) {
	cer, err := certificate.FromPemFile("_fixtures/certificate-valid-encrypted.pem", "password")
	if err != nil {
		t.Fatal(err)
	}
	if e := verifyHostname(cer); e != nil {
		t.Fatal(e)
	}
}

func TestNoSuchFilePemFile(t *testing.T) {
	_, err := certificate.FromPemFile("", "")
	if err.Error() != errors.New("open : no such file or directory").Error() {
		t.Fatal("expected error", "open : no such file or directory")
	}
}

func TestBadPasswordPemFile(t *testing.T) {
	_, err := certificate.FromPemFile("_fixtures/certificate-valid-encrypted.pem", "badpassword")
	if err != certificate.ErrFailedToDecryptKey {
		t.Fatal("expected error", certificate.ErrFailedToDecryptKey)
	}
}

func verifyHostname(cert tls.Certificate) error {
	if cert.Leaf == nil {
		return errors.New("expected leaf cert")
	}
	return cert.Leaf.VerifyHostname("APNS/2 Development IOS Push Services: com.sideshow.Apns2")
}