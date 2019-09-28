package main

import (
	"github.com/gostaticanalysis/builtinprint"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(builtinprint.Analyzer) }
