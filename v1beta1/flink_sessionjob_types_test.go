package v1beta1

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test FlinkSessionJobSpec", func() {
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

		v := j1.Equals(j2)
		Expect(v).To(BeTrue())

		j2.DeploymentName = "test2"
		v = j1.Equals(j2)
		Expect(v).To(BeFalse())

		j2.DeploymentName = "test"
		j2.Job.Parallelism = 2
		v = j1.Equals(j2)
		Expect(v).To(BeFalse())

		j2.Job.Parallelism = 1
		j1.Job.Args = []string{""}
		j2.Job.Args = []string{""}
		v = j1.Equals(j2)
		Expect(v).To(BeTrue())

		j1.Job.Args = []string{"--name", "a", "--topic", "b"}
		j2.Job.Args = []string{"--name", "a", "--topic", "b"}
		v = j1.Equals(j2)
		Expect(v).To(BeTrue())

		j1.Job.Args = []string{"--name", "a", "--topic", "b"}
		j2.Job.Args = []string{"--topic", "b", "--name", "a"}
		v = j1.Equals(j2)
		Expect(v).To(BeTrue())

		j1.Job.Args = []string{"--name", "a", "--topic", "b"}
		j2.Job.Args = []string{"--name", "b", "--topic", "a"}
		v = j1.Equals(j2)
		Expect(v).To(BeFalse())
	})
})
