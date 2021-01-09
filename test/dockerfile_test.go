package test

import (
	"testing"
)

func TestImageAlpine(t *testing.T) {
	runAlpineTest(t)
}

func TestImageNginx(t *testing.T) {
	runNginxTest(t)
}
