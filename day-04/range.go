package main

import (
	"fmt"
	"strconv"
	"strings"
)

type SectionRange struct {
	Start int
	End   int
}

func NewRange(section string) (*SectionRange, error) {
	var (
		err   error = nil
		sr          = new(SectionRange)
		parts       = strings.Split(strings.TrimSpace(section), "-")
	)
	if len(parts) != 2 {
		return nil, fmt.Errorf("incorrect number of parts for range")
	}
	sr.Start, err = strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("bad start of range: %s (%w)", parts[0], err)
	}
	sr.End, err = strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("bad end of range: %s (%w)", parts[1], err)
	}
	return sr, err
}

func NewPair(pairDefinition string) ([]*SectionRange, error) {
	var (
		err   error = nil
		pair        = make([]*SectionRange, 2)
		parts       = strings.Split(strings.TrimSpace(pairDefinition), ",")
	)
	if len(parts) != 2 {
		return nil, fmt.Errorf("incorrect number of parts for pair")
	}
	for k, v := range parts {
		pair[k], err = NewRange(v)
		if err != nil {
			return nil, fmt.Errorf("bad definition for part %d of pair: %w", k+1, err)
		}
	}
	return pair, err
}

func (sr SectionRange) Contains(osr *SectionRange) bool {
	return (sr.Start <= osr.Start) && (sr.End >= osr.End)
}

func (sr SectionRange) isWithin(point int) bool {
	return (sr.Start <= point) && (point <= sr.End)
}

func (sr SectionRange) Overlaps(osr *SectionRange) bool {
	return sr.isWithin(osr.Start) || // osr's Start is somewhere within sr
		sr.isWithin(osr.End) || // osr's End is somewhere within sr
		osr.isWithin(sr.Start) || // or vice-versa
		osr.isWithin(sr.End)
}
