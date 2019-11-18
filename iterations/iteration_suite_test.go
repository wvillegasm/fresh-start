package iterations_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIteration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Iteration Suite")
}
