package main

import (
	"encoding/xml"
)

type Coverage struct {
	XMLName         xml.Name   `xml:"coverage"`
	LineRate        float32    `xml:"line-rate,attr"`
	BranchRate      float32    `xml:"branch-rate,attr"`
	Version         string     `xml:"version,attr"`
	Timestamp       int64      `xml:"timestamp,attr"`
	LinesCovered    int64      `xml:"lines-covered,attr"`
	LinesValid      int64      `xml:"lines-valid,attr"`
	BranchesCovered int64      `xml:"branches-covered,attr"`
	BranchesValid   int64      `xml:"branches-valid,attr"`
	Complexity      float32    `xml:"complexity,attr"`
	Sources         []*Source  `xml:"sources>source"`
	Packages        []*Package `xml:"packages>package"`
}

type Source struct {
	Path string `xml:",chardata"`
}

type Package struct {
	Name       string   `xml:"name,attr"`
	LineRate   float32  `xml:"line-rate,attr"`
	BranchRate float32  `xml:"branch-rate,attr"`
	Complexity float32  `xml:"complexity,attr"`
	Classes    []*Class `xml:"classes>class"`
}

type Class struct {
	Name       string    `xml:"name,attr"`
	Filename   string    `xml:"filename,attr"`
	LineRate   float32   `xml:"line-rate,attr"`
	BranchRate float32   `xml:"branch-rate,attr"`
	Complexity float32   `xml:"complexity,attr"`
	Methods    []*Method `xml:"methods>method"`
	Lines      []*Line   `xml:"lines>line"`
}

type Method struct {
	Name       string  `xml:"name,attr"`
	Signature  string  `xml:"signature,attr"`
	LineRate   float32 `xml:"line-rate,attr"`
	BranchRate float32 `xml:"branch-rate,attr"`
	Complexity float32 `xml:"complexity,attr"`
	Lines      []*Line `xml:"lines>line"`
}

type Line struct {
	Number int   `xml:"number,attr"`
	Hits   int64 `xml:"hits,attr"`
}
