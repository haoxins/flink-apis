package v1beta1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FlinkSessionJobSpec_Equals(t *testing.T) {
	j1 := &FlinkSessionJobSpec{
		DeploymentName: "test",
		Job: JobSpec{
			JarURI:      "test",
			Parallelism: 1,
		},
	}
	j2 := &FlinkSessionJobSpec{
		DeploymentName: "test",
		Job: JobSpec{
			JarURI:      "test",
			Parallelism: 1,
		},
	}

	v := j1.Equals(j2)
	assert.True(t, v)

	j2.DeploymentName = "test2"
	v = j1.Equals(j2)
	assert.False(t, v)

	j2.DeploymentName = "test"
	j2.Job.Parallelism = 2
	v = j1.Equals(j2)
	assert.False(t, v)

	j2.Job.Parallelism = 1
	j1.Job.Args = []string{""}
	j2.Job.Args = []string{""}
	v = j1.Equals(j2)
	assert.True(t, v)

	j1.Job.Args = []string{"--name", "a", "--topic", "b"}
	j2.Job.Args = []string{"--name", "a", "--topic", "b"}
	v = j1.Equals(j2)
	assert.True(t, v)

	j1.Job.Args = []string{"--name", "a", "--topic", "b"}
	j2.Job.Args = []string{"--topic", "b", "--name", "a"}
	v = j1.Equals(j2)
	assert.True(t, v)

	j1.Job.Args = []string{"--name", "a", "--topic", "b"}
	j2.Job.Args = []string{"--name", "b", "--topic", "a"}
	v = j1.Equals(j2)
	assert.False(t, v)
}
