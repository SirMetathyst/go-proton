package proton_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TestProton ...
func TestProton(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Proton Suite")
}
