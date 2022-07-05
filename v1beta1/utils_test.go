package v1beta1

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Utils", func() {
	It("IsArgsEqual should work", func() {
		yes, err := IsArgsEqual(nil, nil)
		Expect(err).To(BeNil())
		Expect(yes).To(BeTrue())
		yes, err = IsArgsEqual([]string{}, []string{})
		Expect(err).To(BeNil())
		Expect(yes).To(BeTrue())
		yes, err = IsArgsEqual([]string{"a"}, []string{"a"})
		Expect(err.Error()).To(Equal("The args must be an even number"))
		Expect(yes).To(BeFalse())
		yes, err = IsArgsEqual([]string{}, []string{"a"})
		Expect(err).To(BeNil())
		Expect(yes).To(BeFalse())
		yes, err = IsArgsEqual([]string{"a", "b"}, []string{"a", "b"})
		Expect(err).To(BeNil())
		Expect(yes).To(BeTrue())
		yes, err = IsArgsEqual([]string{"--topic", "a", "--name", "xin"}, []string{"--name", "xin", "--topic", "a"})
		Expect(err).To(BeNil())
		Expect(yes).To(BeTrue())
		yes, err = IsArgsEqual([]string{"--name", "xin", "--topic", "a"}, []string{"--name", "xin", "--topic", "a"})
		Expect(err).To(BeNil())
		Expect(yes).To(BeTrue())
		yes, err = IsArgsEqual([]string{"--name", "xin", "--topic", "a"}, []string{"--name", "xin", "--topic", "b"})
		Expect(err).To(BeNil())
		Expect(yes).To(BeFalse())
	})
})
