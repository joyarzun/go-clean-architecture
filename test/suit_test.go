package test

import (
	"testing"

	."github.com/onsi/ginkgo/v2"
	."github.com/onsi/gomega"
)

func TestTasks(t *testing.T) {
	RegisterFailHandler(Fail) // connect ginkgo and gomega
	RunSpecs(t, "Tasks Suite") // tell ginkgo to start test suit
}