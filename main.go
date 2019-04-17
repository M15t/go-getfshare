package main

import (
	"go-getfshare/getfs"
	"log"
)

func main() {
	userEmail := `your registered email`
	password := `your password`
	fileURL := `fshare URL`

	s := getfs.NewService(userEmail, password)
	if err := s.Login(); err != nil {
		log.Fatal(err)
	}

	fileInfo, err := s.GetFileInfo(fileURL)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("=== %+v", fileInfo)

	link, err := s.GetLink(fileURL)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("=== %+v", link)
}
