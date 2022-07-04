package shortener_test

import (
	"hypefast-technical/internal/service/link"
	"testing"
)

func TestLink(t *testing.T) {
	url := "https://github.com/hermansetiawan77777"
	newLink := link.AddNewLink(url)
	if newLink == nil {
		t.Errorf("Link not returned")
		// add return to eliminate linting on line 16 newLink.Id (newLink nil possibility)
		return
	}

	if newLink.Url == "" {
		t.Errorf("Please fill url")
		// add return to eliminate linting on line 16 newLink.Id (newLink nil possibility)
		return
	}

	existingLink := link.GetLinkByID(newLink.Id)
	if existingLink == nil {
		t.Errorf("Link not exist even after inserted")
	}
}
