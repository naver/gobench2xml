package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/tools/benchmark/parse"
)

type BenchmarkValueMap map[string]interface{}

func (m BenchmarkValueMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	for k, v := range m {
		t := xml.StartElement{Name: xml.Name{Local: k}}
		if err := e.EncodeElement(v, t); err != nil {
			return err
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

type Benchmarks struct {
	XMLName           xml.Name          `xml:"Benchmarks"`
	NsPerOp           BenchmarkValueMap `xml:"NsPerOp"`
	AllocedBytesPerOp BenchmarkValueMap `xml:"AllocedBytesPerOp"`
	AllocsPerOp       BenchmarkValueMap `xml:"AllocsPerOp"`
	MBPerS            BenchmarkValueMap `xml:"MBPerS"`
}

func prettyName(org string) string {
	tmp := strings.Replace(org, "/", "-", -1)
	re := regexp.MustCompile("-[0-9]*$")
	return re.ReplaceAllString(tmp, "")
}

func main() {
	bs, err := parse.ParseSet(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	bm := &Benchmarks{
		NsPerOp:           make(BenchmarkValueMap),
		AllocedBytesPerOp: make(BenchmarkValueMap),
		AllocsPerOp:       make(BenchmarkValueMap),
		MBPerS:            make(BenchmarkValueMap),
	}

	for name, bb := range bs {
		if len(bb) > 1 {
			fmt.Fprintf(os.Stderr, "[WARN] multiple results for %q found .", name)
		}
		if len(bb) == 0 {
			continue
		}

		b := bb[0]
		prettyName := prettyName(name)
		bm.NsPerOp[prettyName] = b.NsPerOp
		bm.AllocedBytesPerOp[prettyName] = b.AllocedBytesPerOp
		bm.AllocsPerOp[prettyName] = b.AllocsPerOp
		bm.MBPerS[prettyName] = b.MBPerS
	}

	output, err := xml.MarshalIndent(bm, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Stdout.Write(output)
}
