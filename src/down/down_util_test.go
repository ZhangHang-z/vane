package down

import (
	"testing"
)

func TestRveContentsFromLink(t *testing.T) {
	_, err := RveContentsFromLink("https://nodejs.org")
	if err != nil {
		t.Error(err)
	}
}

func TestExtractTarArchive(t *testing.T) {

}
