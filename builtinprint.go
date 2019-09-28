package builtinprint

import (
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "builtinprint",
	Doc:  Doc,
	Run:  run,
}

const Doc = "builtinprint find builtin print and println calling"

func run(pass *analysis.Pass) (interface{}, error) {

	for ident, obj := range pass.TypesInfo.Uses {
		switch ident.Name {
		case "print", "println":
			if obj != nil && obj.Parent() == types.Universe {
				pass.Reportf(ident.Pos(), "must not use builtin %s", ident.Name)
			}
		}
	}

	return nil, nil
}
