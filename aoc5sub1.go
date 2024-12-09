package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var PAGE_ORDERING_REGEX = regexp.MustCompile("[0-9+]\\|[0-9+]")

type AoC5Sub1Processor struct {
	// maps values to what must come after them
	afterRequirements map[string][]string
	updates           []string
}

func (p *AoC5Sub1Processor) ProcessLine(line string) error {
	if len(line) == 0 {
		return nil
	}

	// we are tracking when a number must come after another, because
	// then when we run into it
	if PAGE_ORDERING_REGEX.MatchString(line) {
		orderInstructions := strings.Split(line, "|")
		before, after := orderInstructions[0], orderInstructions[1]
		if _, found := p.afterRequirements[before]; !found {
			p.afterRequirements[before] = []string{after}
		} else {
			p.afterRequirements[before] = append(p.afterRequirements[before], after)
		}
	} else {
		p.updates = append(p.updates, line)
	}
	return nil
}

func (p *AoC5Sub1Processor) Compute() int {
	runningTotal := 0
	for _, update := range p.updates {
		pages := strings.Split(update, ",")
		if p.isUpdateValid(pages) {
			middlePage, _ := strconv.Atoi(pages[(len(pages) / 2)])
			runningTotal += middlePage
		}
	}
	return runningTotal
}

// given a slice of page numbers, ensure that they're all in the right order
func (p *AoC5Sub1Processor) isUpdateValid(pages []string) bool {
	pageToLocation := make(map[string]int)
	for idx, page := range pages {
		if _, found := pageToLocation[page]; !found {
			pageToLocation[page] = idx
		} else {
			// in case there are dupes in the input, which isn't specified
			fmt.Printf("Error: page %d is duplicated in update\n", idx)
		}
	}

	// now, given the locations of each page in the update,
	// go through each page and determine if all the pages that
	// should come after it are either after it
	for _, page := range pages {
		pageLoc := pageToLocation[page]

		if afterPages, foundInstructions := p.afterRequirements[page]; foundInstructions {
			for _, afterPage := range afterPages {
				if afterLoc, foundAfterPage := pageToLocation[afterPage]; foundAfterPage && afterLoc < pageLoc {
					// we found a page that needs to be after, but it's got a lower page number
					return false
				}
			}

		} else {
			// no after instructions for this page, so it's good
			continue
		}
	}
	return true
}

func aoc5sub1() {
	processor := &AoC5Sub1Processor{
		make(map[string][]string),
		make([]string, 0),
	}
	generateSolution(processor)

}
