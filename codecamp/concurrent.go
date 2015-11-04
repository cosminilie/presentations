package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"sync"
	"time"
)

// NewsType Example
type RSS2 struct { // HL
	XMLName     xml.Name `xml:"rss"`
	Version     string   `xml:"version,attr"`
	Title       string   `xml:"channel>title"`
	Link        string   `xml:"channel>link"`        // OMIT
	Description string   `xml:"channel>description"` // OMIT
	PubDate     string   `xml:"channel>pubDate"`     // OMIT
	ItemList    []Item   `xml:"channel>item"`        // HL
}

type Item struct { // HL
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	Content     template.HTML `xml:"encoded"`  // OMIT
	PubDate     string        `xml:"pubDate"`  // OMIT
	Comments    string        `xml:"comments"` // OMIT
}

type NewsStory struct { // HL
	Source      string
	Title       string
	Description template.HTML
	Link        string
}

// ENDNEWSTYPE OMIT
var wg sync.WaitGroup

// FetchNews Example
func NewsFetch(feed string, newsChan chan NewsStory, errChan chan error) { // HL
	resp, err := http.Get(feed)
	defer wg.Done() //OMIT
	// OMIT
	if err != nil {
		errChan <- err // HL
		return
	}
	defer resp.Body.Close()

	news := RSS2{}
	err = xml.NewDecoder(resp.Body).Decode(&news)
	if err != nil { // OMIT
		errChan <- err // OMIT
		return         // OMIT
	} //OMIT

	for _, item := range news.ItemList {
		newsChan <- NewsStory{ // HL
			Source:      news.Title,       // HL
			Title:       item.Title,       // HL
			Description: item.Description, // HL
			Link:        item.Link,        // HL
		}
	}

	fmt.Println("Closing feed:", feed) // OMIT
	return                             // OMIT
}

// ENDEXAMPLE OMIT

func main() {

	// Example RSS
	feeds := []string{"http://www.space.com/home/feed/site.xml",
		"http://spaceflightnow.com/category/news/feed",
		"http://www.nasaspaceflight.com/feed",
		"http://aaa"}

	newsChan := make(chan NewsStory) // HL
	errChan := make(chan error)      // HL

	for _, feed := range feeds {
		fmt.Println("Feed:", feed)            // OMIT
		wg.Add(1)                             // OMIT
		go NewsFetch(feed, newsChan, errChan) // HL
	}

	go func() { // HL
		for {
			select { // HL
			case news := <-newsChan: // HL
				fmt.Printf("News Title: %s\nNews Content: %s\n News Source: %s Source Link: %s\n\n ", news.Title, news.Description, news.Source, news.Link)
				time.Sleep(3 * time.Second) // OMIT
			case err := <-errChan: // HL
				fmt.Println("ReceivedError", err)
			}
		}
	}()

	// ENDEXAMPLERSS OMIT
	wg.Wait()
}
