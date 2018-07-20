package modelprovider

import (
	"github.com/SirMetathyst/go-blackboard"
	"github.com/SirMetathyst/proton/model/builder"

	//protonlang "github.com/SirMetathyst/proton-lang"

	"github.com/SirMetathyst/proton/model"
)

// ProtonLang ...
func ProtonLang(bb *blackboard.BB) (*model.MD, error) {
	mdb := modelbuilder.NewModelBuilder()

	//file, err := os.Open(suffix(*c.StringP("File"), ".proton"))
	//defer file.Close()

	//bytes, err := ioutil.ReadAll(file)
	//e.Must(err)

	//m, err = protonlang.Parse(string(bytes))
	//e.Must(err)
	return mdb.Build()
}
