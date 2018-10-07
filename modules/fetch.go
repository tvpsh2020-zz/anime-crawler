package modules

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tvpsh2020/anime-crawler/config"

	"github.com/PuerkitoBio/goquery"
)

type AnimeQueueContent struct {
	Count int      `json:"count"`
	Title []string `json:"title"`
}

var AnimeQueueList = map[string]*AnimeQueueContent{}
var AnimeList map[string][]string

func readAnimeFromConfig() {
	AnimeList = make(map[string][]string)

	for _, v := range config.Anime.Setting {
		AnimeList[v.QueryString] = v.Keywords
	}
}

func animeDidExist(titleList []string, title string) bool {
	for _, val := range titleList {
		if val == title {
			return true
		}
	}

	return false
}

func insertTitleToQueue(queue map[string]*AnimeQueueContent, queryString, title string) {
	if _, ok := queue[queryString]; !ok {
		tmpAnimeQueueContent := &AnimeQueueContent{
			Count: 0,
			Title: []string{},
		}

		queue[queryString] = tmpAnimeQueueContent
	}

	tmpTitle := queue[queryString].Title
	tmpCount := queue[queryString].Count + 1
	tmpTitle = append(tmpTitle, title)

	queue[queryString] = &AnimeQueueContent{
		Count: tmpCount,
		Title: tmpTitle,
	}
}

func fetchWebsite(needNotify bool, queryString string, keywords []string) error {
	doc, err := goquery.NewDocument(config.Server.Setting.DmhyWebsiteURL + queryString)

	if err != nil {
		fmt.Println(err)
		return err
	}

	doc.Find("#topic_list tbody tr").Each(func(i int, s *goquery.Selection) {
		title := s.Find("td.title a").Text()
		title = strings.Replace(title, "\t", "", -1)
		title = strings.Replace(title, "\n", "", -1)
		title = strings.Replace(title, " ", "", -1)

		for _, val := range keywords {
			if strings.Contains(title, val) {

				if !animeDidExist(AnimeQueueList[queryString].Title, title) {
					log.Printf("DOM record: %d, title: %s", i, title)
					insertTitleToQueue(AnimeQueueList, queryString, title)
					if needNotify {
						mailToDest(title)
					}
					log.Printf("Array total count: %d", AnimeQueueList[queryString].Count)
					break
				}
			}
		}
	})

	log.Printf("Job done.")
	return nil
}

func initAnimeQueueListByConfig(config map[string][]string, queue map[string]*AnimeQueueContent) {
	for key := range config {
		if _, ok := queue[key]; !ok {
			tmpAnimeQueueContent := &AnimeQueueContent{
				Count: 0,
				Title: []string{},
			}

			queue[key] = tmpAnimeQueueContent
		}
	}
}

func fetchWebsiteFromConfig(args map[string][]string, needNotify bool) error {
	for queryString, keywords := range args {
		if err := fetchWebsite(needNotify, queryString, keywords); err != nil {
			return err
		}
	}

	return nil
}

func FetchDmhy() {
	readAnimeFromConfig()
	initAnimeQueueListByConfig(AnimeList, AnimeQueueList)
	checkServerStatus := time.Tick(config.Server.Setting.CheckServerStatusTimeInMinute * time.Minute)
	fetch := time.Tick(config.Server.Setting.FetchTimeInMinute * time.Minute)

	if err := fetchWebsiteFromConfig(AnimeList, false); err != nil {
		panic("Initializing error.")
	}

	for {
		select {
		case <-checkServerStatus:
			log.Printf("Server still alive.")
		case <-fetch:
			log.Printf("Time to get anime!")
			if err := fetchWebsiteFromConfig(AnimeList, true); err != nil {
				log.Printf("Job unfinished.")

				break
			}
		}
	}
}
