package layout_parser

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func (parser *LayoutParserImpl) getStarRating(countryDbArr *[]CountryDB, packageName string) *[]Rating {

	if countryDbArr == nil || len(*countryDbArr) == 0 {
		return &[]Rating{}
	}

	var ratingArr []Rating
	countryDbChan := make(chan CountryDB)

	go func() {
		for _, item := range *countryDbArr {
			countryDbChan <- item
		}
		close(countryDbChan)
	}()

	ratingChan := parser.fanOut(countryDbChan, len(*countryDbArr), packageName)

	for result := range ratingChan {
		ratingArr = append(ratingArr, result)
	}

	return &ratingArr
}

func (parser *LayoutParserImpl) fanOut(
	countryDbChan <-chan CountryDB,
	countryDbLen int,
	packageName string,
) <-chan Rating {
	ratingChan := make(chan Rating, 60)
	var wg sync.WaitGroup

	for i := 0; i <= countryDbLen; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			parser.ratingWorker(countryDbChan, ratingChan, packageName)
		}()
	}

	go func() {
		wg.Wait()
		close(ratingChan)
	}()

	return ratingChan
}

func (parser *LayoutParserImpl) ratingWorker(
	countryDbChan <-chan CountryDB,
	ratingChan chan<- Rating,
	packageName string,
) {
	for countryDb := range countryDbChan {
		var rating = Rating{
			CountryName: countryDb.CountryName,
			Url:         getUrlByPackage(packageName, &countryDb),
		}
		fmt.Println(rating.Url)
		var htmlResponse = parser.network.MakeGetRequest(rating.Url)
		time.Sleep(time.Millisecond * 500)
		if htmlResponse.Error != nil || htmlResponse.StatusCode != http.StatusOK || len(htmlResponse.Body) == 0 {
			fmt.Println(htmlResponse)
			continue
		}
		rating.Rating = getRating(string(htmlResponse.Body))
		if rating.Rating != "" {
			ratingChan <- rating
		}
	}
}
