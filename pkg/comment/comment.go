package comment

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyser reports commented code.
var Analyzer = &analysis.Analyzer{
	Name:     "comment",
	Doc:      "Reports commented code.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

const statementsWrapper = `
package comment
func commentVeryUniqueFunctionName() {
%s
}
`

const functionsWrapper = `
package comment
%s
`

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	fileFilter := []ast.Node{
		(*ast.File)(nil),
	}

	inspector.Preorder(fileFilter, func(node ast.Node) {
		f := node.(*ast.File)
		for _, cg := range f.Comments {
			cg.Pos()
			cg.Text()
			t := cg.Text()
			text := strings.TrimSpace(t)
			if len(text) == 0 {
				continue
			}

			if len(text) > 4 && text[0:4] == "func" {
				src := fmt.Sprintf(functionsWrapper, t)
				_, err := parser.ParseFile(token.NewFileSet(), "", src, parser.DeclarationErrors)
				if err == nil {
					pass.Reportf(cg.Pos(), "commented code")
				}
				continue
			}
			if len(text) > 7 && text[0:7] == "package" {
				_, err := parser.ParseFile(token.NewFileSet(), "", text, parser.DeclarationErrors)
				if err == nil {
					pass.Reportf(cg.Pos(), "commented code")
				}
				continue
			}

			src := fmt.Sprintf(statementsWrapper, t)
			astFile, err := parser.ParseFile(token.NewFileSet(), "", src, parser.DeclarationErrors)
			if err != nil {
				continue
			}

			stmts := astFile.Decls[0].(*ast.FuncDecl).Body.List
			if len(stmts) == 0 {
				pass.Reportf(cg.Pos(), "empty double comment")
				continue
			}
			if len(stmts) == 1 {
				_, ok := stmts[0].(*ast.LabeledStmt)
				if ok {
					continue
				}
				_, ok = stmts[0].(*ast.ExprStmt)
				if ok {
					continue
				}
				_, ok = stmts[0].(*ast.ReturnStmt)
				if ok {
					continue
				}
			}

			pass.Reportf(cg.Pos(), "commented code")
		}
	})

	return nil, nil
}
