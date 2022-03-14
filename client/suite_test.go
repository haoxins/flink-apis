package client

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFlinkAPIs(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Flink APIs Suite")
}
