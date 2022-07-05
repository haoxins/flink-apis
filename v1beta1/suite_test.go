package v1beta1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLibs(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Flink APIs Suite")
}
