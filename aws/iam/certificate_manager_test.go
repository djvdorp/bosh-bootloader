package iam_test

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/pivotal-cf-experimental/bosh-bootloader/aws/iam"
	"github.com/pivotal-cf-experimental/bosh-bootloader/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CertificateManager", func() {
	var (
		certificateUploader  *fakes.CertificateUploader
		certificateDescriber *fakes.CertificateDescriber
		certificateDeleter   *fakes.CertificateDeleter
		manager              iam.CertificateManager
		certificateFile      *os.File
		privateKeyFile       *os.File
		chainFile            *os.File
	)

	BeforeEach(func() {
		var err error
		certificateUploader = &fakes.CertificateUploader{}
		certificateDescriber = &fakes.CertificateDescriber{}
		certificateDeleter = &fakes.CertificateDeleter{}
		manager = iam.NewCertificateManager(certificateUploader, certificateDescriber, certificateDeleter)

		certificateFile, err = ioutil.TempFile("", "")
		Expect(err).NotTo(HaveOccurred())

		privateKeyFile, err = ioutil.TempFile("", "")
		Expect(err).NotTo(HaveOccurred())

		chainFile, err = ioutil.TempFile("", "")
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Create", func() {
		It("creates the given certificate", func() {
			certificateUploader.UploadCall.Returns.CertificateName = "some-new-certificate"

			certificateName, err := manager.Create(certificateFile.Name(), privateKeyFile.Name(), chainFile.Name())
			Expect(err).NotTo(HaveOccurred())

			Expect(certificateUploader.UploadCall.CallCount).To(Equal(1))
			Expect(certificateUploader.UploadCall.Receives.CertificatePath).To(Equal(certificateFile.Name()))
			Expect(certificateUploader.UploadCall.Receives.PrivateKeyPath).To(Equal(privateKeyFile.Name()))
			Expect(certificateUploader.UploadCall.Receives.ChainPath).To(Equal(chainFile.Name()))

			Expect(certificateName).To(Equal("some-new-certificate"))
		})

		Context("failure cases", func() {
			Context("when certificate uploader fails to upload", func() {
				It("returns an error", func() {
					certificateUploader.UploadCall.Returns.Error = errors.New("upload failed")

					_, err := manager.Create(certificateFile.Name(), privateKeyFile.Name(), chainFile.Name())
					Expect(err).To(MatchError("upload failed"))
				})
			})
		})
	})

	Describe("Delete", func() {
		It("deletes the given certificate", func() {
			err := manager.Delete("some-certificate-name")
			Expect(err).NotTo(HaveOccurred())

			Expect(certificateDeleter.DeleteCall.Receives.CertificateName).To(Equal("some-certificate-name"))
		})

		Context("failure cases", func() {
			It("returns an error when certificate fails to delete", func() {
				certificateDeleter.DeleteCall.Returns.Error = errors.New("unknown certificate error")

				err := manager.Delete("some-non-existant-certificate")
				Expect(err).To(MatchError("unknown certificate error"))
			})
		})
	})

	Describe("Describe", func() {
		It("returns a certificate", func() {
			certificateDescriber.DescribeCall.Returns.Certificate = iam.Certificate{
				Name: "some-certificate-name",
				ARN:  "some-certificate-arn",
				Body: "some-certificate-body",
			}

			certificate, err := manager.Describe("some-certificate-name")
			Expect(err).NotTo(HaveOccurred())

			Expect(certificate).To(Equal(iam.Certificate{
				Name: "some-certificate-name",
				ARN:  "some-certificate-arn",
				Body: "some-certificate-body",
			}))
			Expect(certificateDescriber.DescribeCall.Receives.CertificateName).To(Equal("some-certificate-name"))
		})

		Context("failure cases", func() {
			It("returns an error when the describe fails", func() {
				certificateDescriber.DescribeCall.Returns.Error = errors.New("unknown certificate error")

				_, err := manager.Describe("some-non-existant-certificate")
				Expect(err).To(MatchError("unknown certificate error"))
			})
		})
	})
})
