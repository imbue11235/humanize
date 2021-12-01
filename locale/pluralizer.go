package locale

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	boundaryRegex   = "^\\[(\\d+)-?(\\d+|\\*)?\\](.*$)"
	pluralSeparator = "|"
)

type boundary struct {
	min, max int
	text     string
}

type pluralizer struct {
	source     string
	boundaries []*boundary
}

func (p *pluralizer) getText(count int) string {
	for _, b := range p.boundaries {
		// if min is equal to max and count exactly matches min
		// return text, as this would be a singular boundary e.g. [1] some text
		if count == b.min && b.max == b.min {
			return b.text
		}

		// if count is greater than min and less than max or
		// max is equal to * (-1, infinity), return text
		if count >= b.min && (b.max == -1 || count <= b.max) {
			return b.text
		}
	}

	// if no boundaries match, return the last boundary
	return p.boundaries[len(p.boundaries)-1].text
}

func (p *pluralizer) apply(count int) string {
	return p.applyCountToText(p.getText(count), count)
}

func (p *pluralizer) applyCountToText(text string, count int) string {
	if strings.Contains(text, "%d") {
		return fmt.Sprintf(text, count)
	}

	return text
}

func createPluralizer(source string) (*pluralizer, error) {
	boundaries, err := getBoundariesFromSource(source)

	if err != nil {
		return nil, err
	}

	return &pluralizer{source, boundaries}, nil
}

func createBoundaryFromSubmatch(group []string) (*boundary, error) {
	// captured groups length must be at least 2 and maximum 4,
	// to contain full-capture, min, max? and text
	if len(group) != 4 {
		return nil, fmt.Errorf("invalid boundary format: %s", group)
	}

	minStringValue := group[1]
	maxStringValue := group[2]
	text := strings.TrimSpace(group[3])

	// convert minimum bound value to int
	// e.g. [1] some text
	//
	// we could void this error check as
	// we already know that the regex matched
	min, _ := strconv.Atoi(minStringValue)
	bound := &boundary{min: min, max: -1, text: text}

	// if both groups are captured
	// e.g. [1-3] some text
	// or [1-*] some text
	if maxStringValue != "" {
		if maxStringValue == "*" {
			bound.max = -1
			return bound, nil
		}

		// we can void this error check as
		// we already know that the regex matched
		max, _ := strconv.Atoi(maxStringValue)

		bound.max = max

		if max < min {
			return nil, fmt.Errorf("invalid boundary format, max must be higher than min: %s", group)
		}

		return bound, nil
	}

	// if max is not set
	// max will be equal to min
	bound.max = bound.min

	return bound, nil
}

func getBoundariesFromSource(source string) ([]*boundary, error) {
	var boundaries []*boundary
	parts := strings.Split(source, pluralSeparator)
	if len(parts) == 1 {
		return nil, fmt.Errorf("invalid plural source: %s", source)
	}

	// keeping track of last boundary to
	// make sure we don't have overlapping boundaries
	lastBoundary := &boundary{min: -1, max: -1}

	r := regexp.MustCompile(boundaryRegex)

	for _, part := range parts {
		groups := r.FindStringSubmatch(part)
		bound, err := createBoundaryFromSubmatch(groups)
		if err != nil {
			return nil, err
		}

		// if the new boundary is overlapping with the last boundary
		// we need to merge them
		// unless the current boundary max is equal to * (-1, infinity)
		if bound.min <= lastBoundary.max && bound.max != -1 {
			bound.min = lastBoundary.max + 1

			// if min is higher than max, set max to min
			if bound.min >= bound.max {
				bound.max = bound.min
			}

		}

		boundaries = append(boundaries, bound)

		// set the last boundary to the current boundary
		lastBoundary = bound

		// if bound max is infinity,
		// there should not be any other boundaries
		// as they would be redundant
		if bound.max == -1 {
			break
		}
	}

	return boundaries, nil
}
