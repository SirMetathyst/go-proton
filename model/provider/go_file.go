package modelprovider

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/SirMetathyst/go-blackboard"
	"github.com/SirMetathyst/go-entitas"
	"github.com/SirMetathyst/go-entitas/builder"
)

// functionData [DATA]...
type functionData struct {
	Func string
	Data []string
}

// newFunctionData ...
func newFunctionData(id string) *functionData {
	return &functionData{id, make([]string, 0)}
}

// AddData ...
func (fd *functionData) AddData(v ...string) {
	fd.Data = append(fd.Data, v...)
}

// fM ...
type fM struct {
	CurrentFunctionDataList []*functionData

	FunctionData map[string][][]*functionData
}

// newFileModel ...
func newFileModel() *fM {
	return &fM{make([]*functionData, 0), make(map[string][][]*functionData, 0)}
}

// AddFuncData ...
func (d *fM) AddFuncData(v *functionData) {
	d.CurrentFunctionDataList = append(d.CurrentFunctionDataList, v)
}

// Collapse ...
func (d *fM) Collapse(v string) {
	d.FunctionData[v] = append(d.FunctionData[v], d.CurrentFunctionDataList)
	d.CurrentFunctionDataList = make([]*functionData, 0)
}

var (
	fileModel = newFileModel()
)

// print ...
func print() {
	for k, v := range fileModel.FunctionData {
		fmt.Printf("%s\n\n", k)
		for i, s := range v {
			fmt.Printf("%d:(", i)
			for _, f := range s {
				fmt.Printf("%s(", f.Func)
				for _, v := range f.Data {
					fmt.Printf("[%s]", v)
				}
				fmt.Printf("), ")
			}
			fmt.Printf(")\n")
		}
		fmt.Printf("\n")
	}
}

type walker struct{}

func (w *walker) Visit(node ast.Node) ast.Visitor {
	switch t := node.(type) {
	case *ast.CallExpr:
		if selector, ok := t.Fun.(*ast.SelectorExpr); ok {

			f := newFunctionData(selector.Sel.Name)

			for _, a := range t.Args {
				if lit, ok := a.(*ast.BasicLit); ok {
					f.AddData(strings.Trim(lit.Value, "\""))
				}
			}

			fileModel.AddFuncData(f)

			if ident, ok := selector.X.(*ast.Ident); ok {
				fileModel.Collapse(ident.Name)
			}
		} else if ident, ok := t.Fun.(*ast.Ident); ok {
			fileModel.Collapse(ident.Name)
		}
	}
	return w
}

// GoFile ...
func GoFile(bb *blackboard.BB) (*entitas.MD, error) {
	mdb := entitasbuilder.NewModelBuilder()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, suffix(*bb.StringP("File"), ".go"), nil, 0)
	if err != nil {
		return nil, err
	}
	ast.Walk(new(walker), file)
	print()

	return mdb.Build()
}
