package service

import (
	"fmt"
	"log"
	"strconv"
)

func ReadWebPage(url string) (ResponseDto, error) {
	var responseDto ResponseDto

	htmlVerion, err := getHTMLVersion(url)
	if err != nil {
		log.Println("This is an error")
		responseDto.HtmlVersion = "Please check the input URL"
		return responseDto, nil
	}

	htmlTitle, err := getPageTitle(url)
	if err != nil {
		log.Println("This is an error")
		responseDto.Title = "Please check the input URL"
		return responseDto, nil
	}

	htmlHeadings, err := countHeadings(url)
	if err != nil {
		log.Println("This is an error")
		responseDto.Title = "Please check the input URL"
		return responseDto, nil
	}

	internalCount, externalCount, err := countLinks(url)
	if err != nil {
		log.Println("Error: %v\n", err)
	}
	fmt.Printf("Number of Internal Links in %s: %d\n", url, internalCount)
	fmt.Printf("Number of External Links in %s: %d\n", url, externalCount)

	// response
	responseDto.HtmlVersion = htmlVerion
	responseDto.Title = htmlTitle
	responseDto.HeadLines = strconv.Itoa(htmlHeadings)
	responseDto.NumberOfLinks = internalCount + externalCount

	return responseDto, nil
}
