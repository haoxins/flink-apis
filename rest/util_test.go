package rest

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test util", func() {
	It("GenJobId should work", func() {
		Expect(GenJobId("Hello", "Flink")).To(Equal("12e124280630fa9691c0bdab95782493"))
	})
})
