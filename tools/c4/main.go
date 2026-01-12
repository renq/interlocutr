package main

import (
	"fmt"
	"os"

	"github.com/krzysztofreczek/go-structurizr/pkg/scraper"
	"github.com/krzysztofreczek/go-structurizr/pkg/view"

	comments "github.com/renq/interlocutr/internal/comments/factory"
)

const (
	scraperConfig = "scraper.yml"
	viewConfig    = "view.yml"
	outputFile    = "out/view-%s.plantuml"
)

func main() {
	commentsApp := comments.BuildApp()
	scrape(commentsApp, "comments")
}

func scrape(app interface{}, name string) {
	s, err := scraper.NewScraperFromConfigFile(scraperConfig)
	if err != nil {
		panic(err)
	}

	structure := s.Scrape(app)

	outFileName := fmt.Sprintf(outputFile, name)
	outFile, err := os.Create(outFileName)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = outFile.Close()
	}()

	v, err := view.NewViewFromConfigFile(viewConfig)
	if err != nil {
		panic(err)
	}

	err = v.RenderStructureTo(structure, outFile)
	if err != nil {
		panic(err)
	}
}
