package newsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type Client struct {
	ApiKey			string
}

type jsonPayload struct {
	q 			string
	sources			string
	language 		string
	domains			string
	excludeDomains 		string
	from			string
	to			string
	sortBy			string
	country			string
	category		string
	pageSize		string
	page			string
}

func (c Client) GetTopHeadlines(query string, sources string, language string, country string, category string, pageSize int, page int) []Article {

	var payload jsonPayload

	// Carry out any checks and add data to the payload
	if query != "" {
		payload.q = query
	}

	if (sources != "") && ((country != "") || (category != "")) {
		log.Fatal("cannot mix country/category param with sources param")
	}

	if sources != "" {
		payload.sources = sources
	}

	if language != "" {
		if checkStringInArray(language, Languages) {
			payload.language = language
		} else {
			log.Fatal("invalid language")
		}
	}

	if country != "" {
		if checkStringInArray(country, Countries) {
			payload.country = country
		} else {
			log.Fatal("invalid country")
		}
	}

	if category != "" {
		if checkStringInArray(category, Categories) {
			payload.category = category
		} else {
			log.Fatal("invalid category")
		}
	}

	if pageSize != 0 {
		if pageSize > 0 && pageSize <= 100 {
			payload.pageSize = string(pageSize)
		} else {
			log.Fatal("page size param should be an int between 1 and 100")
		}
	}

	if page != 0 {
		if page > 0 {
			payload.page = string(page)
		} else {
			log.Fatal("page param should be an int greater than zero")
		}
	}

	resp, err := makeRequest(TOP_HEADLINES_URL, payload, c.ApiKey)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Could not read the response")
	}

	var data articlesResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("error unmarshalling json")
	}

	return data.Articles
}

func (c Client) GetEverything(query string, sources string, domains string, excludeDomains string, from string, to string, language string, sortBy string, pageSize int, page int) []Article{

	var payload jsonPayload

	// Carry out any checks and add data to the payload
	if query != "" {
		payload.q = query
	}

	if sources != "" {
		payload.sources = sources
	}

	if domains != "" {
		payload.domains = domains
	}

	if excludeDomains != "" {
		payload.excludeDomains = excludeDomains
	}

	if language != "" {
		if checkStringInArray(language, Languages) {
			payload.language = language
		} else {
			log.Fatal("invalid language")
		}
	}

	if from != "" {
		if len(from) >= 10 {
			for i:=0; i < len(from); i++ {
				if (i == 4 && string(from[i]) != "-") || (i == 7 && string(from[i]) != "-") {
					log.Fatal("from param should be in the format YYYY-MM-DD or YYYY-MM-DDThh:mm:ss")
					break
				} else {
					payload.from = from
					break
				}

			}
		} else {
			log.Fatal("from param should be in the format YYYY-MM-DD or YYYY-MM-DDThh:mm:ss")
		}
	}

	if to != "" {
		if len(to) >= 10 {
			for i:=0; i < len(to); i++ {
				if (i == 4 && string(to[i]) != "-") || (i == 7 && string(to[i]) != "-") {
					log.Fatal("to param should be in the format YYYY-MM-DD or YYYY-MM-DDThh:mm:ss")
					break
				} else {
					payload.to = to
					break
				}
			}
		} else {
			log.Fatal("to param should be in the format YYYY-MM-DD or YYYY-MM-DDThh:mm:ss")
		}
	}

	// check that from is before to
	if from != "" && to != "" {

		var fromTime time.Time
		var toTime time.Time

		fromTime, err := time.Parse(layoutWithoutTime, from)
		if err != nil {
			fromTime, err = time.Parse(layoutWithTime, from)
			if err != nil {
				log.Fatal("from param should be in the format YYYY-MM-DD or YYYY-MM-DDThh:mm:ss")
			}
		}

		toTime, err = time.Parse(layoutWithoutTime, to)
		if err != nil {
			toTime, err = time.Parse(layoutWithTime, to)
			if err != nil {
				log.Fatal("to param should be in the format YYYY-MM-DD or YYYY-MM-DDThh:mm:ss")
			}
		}

		if !fromTime.Before(toTime) {
			log.Fatal("from param should be a date before the to param")
		}
	}

	if sortBy != "" {
		if checkStringInArray(sortBy, SortMethod) {
			payload.sortBy = sortBy
		}
	}

	if pageSize != 0 {
		if pageSize > 0 && pageSize <= 100 {
			payload.pageSize = string(pageSize)
		} else {
			log.Fatal("page size param should be an int between 1 and 100")
		}
	}

	if page != 0 {
		if page > 0 {
			payload.page = string(page)
		} else {
			log.Fatal("page param should be an int greater than zero")
		}
	}

	resp, err := makeRequest(EVERYTHING_URL, payload, c.ApiKey)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Could not read the response")
	}

	var data articlesResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("error unmarshalling json")
	}

	return data.Articles
}

func (c Client) GetSources(category string, language string, country string) []Source {
	var payload jsonPayload

	if language != "" {
		if checkStringInArray(language, Languages) {
			payload.language = language
		} else {
			log.Fatal("invalid language")
		}
	}

	if country != "" {
		if checkStringInArray(country, Countries) {
			payload.country = country
		} else {
			log.Fatal("invalid country")
		}
	}

	if category != "" {
		if checkStringInArray(category, Categories) {
			payload.category = category
		} else {
			log.Fatal("invalid category")
		}
	}

	resp, err := makeRequest(SOURCES_URL, payload, c.ApiKey)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Could not read the response")
	}

	var data sourcesResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("error unmarshalling json")
	}
	return data.Sources
}
