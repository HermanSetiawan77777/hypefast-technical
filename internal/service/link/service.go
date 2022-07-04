package link

import (
	"fmt"
	"math/rand"
	"time"
)

var Links []*Link

const idLength = 6

//Ideally we should create a service struct

func AddNewLink(url string) *Link {
	newLink := &Link{
		Id:            generateUniqueId(),
		Url:           url,
		CreatedAt:     time.Now(),
		RedirectCount: 0,
	}
	Links = append(Links, newLink)
	return newLink
}

func GetLinkByID(id string) *Link {
	for _, l := range Links {
		if l.Id == id {
			return l
		}
	}

	return nil
}

func UpdateRedirectCount(id string, newRedirectCount int) error {
	for _, l := range Links {
		if l.Id == id {
			l.RedirectCount = newRedirectCount
			return nil
		}
	}

	return fmt.Errorf("Link not found")
}

func generateUniqueId() string {
	newId := generateRandomString(idLength)
	existingLink := GetLinkByID(newId)
	for existingLink != nil {
		newId = generateRandomString(idLength)
		existingLink = GetLinkByID(newId)
	}

	return newId
}

func generateRandomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
