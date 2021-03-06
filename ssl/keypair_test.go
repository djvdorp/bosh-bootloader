package ssl_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf-experimental/bosh-bootloader/ssl"
)

var _ = Describe("KeyPair", func() {
	Describe("IsEmpty", func() {
		It("returns true if the keypair is empty", func() {
			keyPair := ssl.KeyPair{}

			Expect(keyPair.IsEmpty()).To(BeTrue())
		})

		It("returns false if the keypair is not empty", func() {
			keyPair := ssl.KeyPair{
				Certificate: []byte("some-cert"),
				PrivateKey:  []byte("some-key"),
			}

			Expect(keyPair.IsEmpty()).To(BeFalse())
		})
	})

	Describe("IsValidForIP", func() {
		It("returns false if the keypair is empty", func() {
			keyPair := ssl.KeyPair{}

			Expect(keyPair.IsValidForIP("127.0.0.1")).To(BeFalse())
		})

		It("returns false if the keypair is not empty", func() {
			keyPair := ssl.KeyPair{
				Certificate: []byte(certificate),
				PrivateKey:  []byte(privateKey),
			}

			Expect(keyPair.IsValidForIP("127.0.0.1")).To(BeFalse())
			Expect(keyPair.IsValidForIP("52.0.112.12")).To(BeTrue())
		})

		Context("failure cases", func() {
			Context("when the cert cannot be decoded", func() {
				It("returns false", func() {
					keyPair := ssl.KeyPair{
						Certificate: []byte(privateKey),
						PrivateKey:  []byte(certificate),
					}

					Expect(keyPair.IsValidForIP("52.0.112.12")).To(BeFalse())
				})
			})

			Context("when the cert is not PEM encoded", func() {
				It("returns false", func() {
					keyPair := ssl.KeyPair{
						Certificate: []byte("something"),
						PrivateKey:  []byte("something"),
					}

					Expect(keyPair.IsValidForIP("52.0.112.12")).To(BeFalse())
				})
			})
		})
	})
})
