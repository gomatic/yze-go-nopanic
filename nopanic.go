// Package nopanic provides a go/analysis analyzer that forbids calls to the
// built-in panic outside test files, per the gomatic Go standard that errors are
// returned rather than raised by panicking.
package nopanic

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	goyze "github.com/gomatic/go-yze"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const message = "panic is not permitted outside tests; return an error instead"

// Analyzer reports calls to the built-in panic in non-test files.
var Analyzer = &analysis.Analyzer{
	Name:     "nopanic",
	Doc:      "reports calls to the built-in panic outside test files, which the gomatic Go standard forbids in favor of returning errors",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

// Registration declares this analyzer to the yze framework.
var Registration = goyze.Registration{
	Name:       "nopanic",
	Categories: []goyze.Category{"patterns"},
	URL:        "https://docs.gomatic.dev/yze/nopanic",
	Analyzer:   Analyzer,
}

// run reports every built-in panic call in a non-test file of the analyzed package.
func run(pass *analysis.Pass) (any, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	insp.Preorder([]ast.Node{(*ast.CallExpr)(nil)}, func(n ast.Node) {
		call := n.(*ast.CallExpr)
		if isBuiltinPanic(pass, call) && !inTestFile(pass, call.Pos()) {
			pass.Reportf(call.Pos(), message)
		}
	})
	return nil, nil
}

// isBuiltinPanic reports whether call invokes the predeclared panic built-in
// (not a shadowing local or variable named panic).
func isBuiltinPanic(pass *analysis.Pass, call *ast.CallExpr) bool {
	fun, ok := call.Fun.(*ast.Ident)
	if !ok || fun.Name != "panic" {
		return false
	}
	_, builtin := pass.TypesInfo.ObjectOf(fun).(*types.Builtin)
	return builtin
}

// inTestFile reports whether pos lies in a _test.go file.
func inTestFile(pass *analysis.Pass, pos token.Pos) bool {
	return strings.HasSuffix(pass.Fset.File(pos).Name(), "_test.go")
}
