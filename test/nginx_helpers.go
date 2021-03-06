package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func runNginxTest(t *testing.T) {
	tag := "terratest/nginx"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	docker.Build(t, "../Dockerfile/nginx/", buildOptions)

	opts := &docker.RunOptions{Command: []string{"cat", "/usr/share/nginx/html/index.html"}}
	output := docker.Run(t, tag, opts)
	assert.Equal(t, "Hello, Terratest!", output)

}
