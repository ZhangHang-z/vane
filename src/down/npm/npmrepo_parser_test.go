package npm

import (
	"testing"
)

func TestNPMRegistryInit(t *testing.T) {
	url := "http://registry.npmjs.com/vue"
	npmRepo := NPMRegistryInit(url)
	npmRepo.PrintAvailableVersions()
}
