package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMusicXmlGoParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MusicXmlGoParser Suite")
}
