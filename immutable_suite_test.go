package immutable_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "immutable test")
}
