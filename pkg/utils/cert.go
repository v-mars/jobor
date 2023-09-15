package utils

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// CertInfo "test.pem"
func CertInfo(certPath string, certBytes []byte) (*x509.Certificate, error) {
	var err error
	if len(certPath) > 0 {
		certBytes, err = ioutil.ReadFile(certPath)
		if err != nil {
			//fmt.Println(err.Error())
			return nil, err
		}
	}

	pemBlock, _ := pem.Decode(certBytes)
	if pemBlock == nil {
		//fmt.Println("decode error")
		return nil, fmt.Errorf("decode pem/cert error")
	}

	cert, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		//fmt.Println(err.Error())
		return cert, err
	}

	//fmt.Printf("Name %s\n", cert.Subject.CommonName)       // 域名
	//fmt.Printf("Not before %s\n", cert.NotBefore.String()) // 有效期开始时间
	//fmt.Printf("Not after %s\n", cert.NotAfter.String())   // 有效期结束时间
	return cert, nil
}
