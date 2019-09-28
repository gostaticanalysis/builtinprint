package builtinprint_test

import (
	"testing"

	"github.com/gostaticanalysis/builtinprint"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, builtinprint.Analyzer, "a")
}