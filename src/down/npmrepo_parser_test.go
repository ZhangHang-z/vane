package down

import (
	"testing"
)

func TestGetNPMRegistry(t *testing.T) {
	url := "http://registry.npmjs.com/vue"
	GetNPMRegistry(url)
}
