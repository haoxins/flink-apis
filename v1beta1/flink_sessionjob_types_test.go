package v1beta1

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test FlinkSessionJobSpec Equals", func() {
	It("FlinkSessionJobSpec Equals should work", func() {
		j1 := FlinkSessionJobSpec{
			DeploymentName: "test",
			Job: JobSpec{
				JarURI:      "test",
				Parallelism: 1,
			},
		}
		j2 := FlinkSessionJobSpec{
			DeploymentName: "test",
			Job: JobSpec{
				JarURI:      "test",
				Parallelism: 1,
			},
		}

		yes, err := j1.Equals(j2)
		Expect(err).To(BeNil())
		Expect(yes).To(BeTrue())

		j2.DeploymentName = "test2"
		yes, err = j1.Equals(j2)
		Expect(err).To(BeNil())
		Expect(yes).To(BeFalse())

		j2.DeploymentName = "test"
		j2.Job.Parallelism = 2
		yes, err = j1.Equals(j2)
		Expect(err).To(BeNil())
		Expect(yes).To(BeFalse())

		j2.Job.Parallelism = 1
		j1.Job.Args = []string{""}
		j2.Job.Args = []string{""}
		yes, err = j1.Equals(j2)
		Expect(err.Error()).To(Equal("the args must be an even number"))
		Expect(yes).To(BeFalse())

		j1.Job.Args = []string{"--name", "a", "--topic", "b"}
		j2.Job.Args = []string{"--name", "a", "--topic", "b"}
		yes, err = j1.Equals(j2)
		Expect(err).To(BeNil())
		Expect(yes).To(BeTrue())

		j1.Job.Args = []string{"--name", "a", "--topic", "b"}
		j2.Job.Args = []string{"--topic", "b", "--name", "a"}
		yes, err = j1.Equals(j2)
		Expect(err).To(BeNil())
		Expect(yes).To(BeTrue())

		j1.Job.Args = []string{"--name", "a", "--topic", "b"}
		j2.Job.Args = []string{"--name", "b", "--topic", "a"}
		yes, err = j1.Equals(j2)
		Expect(err).To(BeNil())
		Expect(yes).To(BeFalse())
	})
})
