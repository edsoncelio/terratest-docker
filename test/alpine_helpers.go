package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func runAlpineTest(t *testing.T) {
	tag := "terratest/alpine"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	docker.Build(t, "../Dockerfile/alpine/", buildOptions)

	opts := &docker.RunOptions{Command: []string{"date", "+'%Z'"}}
	output := docker.Run(t, tag, opts)
	assert.Equal(t, "'UTC'", output)

}
