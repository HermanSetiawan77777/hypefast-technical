package link

import (
	"fmt"
	"math/rand"
	"time"
)

//var Links []*Link

var m = make(map[string]*Link)

const idLength = 6

//Ideally we should create a service struct

func AddNewLink(url, optionShort string) *Link {
	var tempstr string
	//sharon.raissa@hypefast.id
	if optionShort == "" {
		tempstr = generateUniqueId()
	}
	if m[optionShort] != nil && optionShort != "" {
		fmt.Println(m[optionShort])
		tempstr = generateUniqueId()
	} else if optionShort != "" {
		tempstr = optionShort
	}
	newLink := &Link{
		Id:            tempstr,
		Url:           url,
		CreatedAt:     time.Now(),
		RedirectCount: 0,
	}
	m[tempstr] = newLink
	//Links = append(Links, newLink)
	return newLink
}

func GetLinkByID(id string) *Link {
	if m[id] != nil {
		return m[id]
	}

	return nil
}

func UpdateRedirectCount(id string, newRedirectCount int) error {

	if m[id] != nil {
		m[id].RedirectCount = newRedirectCount
		return nil
	}

	// for _, l := range m {
	// 	if l.Id == id {
	// 		l.RedirectCount = newRedirectCount
	// 		return nil
	// 	}
	// }

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
