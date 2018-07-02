package modelprovider

import (
	"os"

	//protonlang "github.com/SirMetathyst/proton-lang"
	"github.com/SirMetathyst/proton/configuration"
	"github.com/SirMetathyst/proton/model"
	"github.com/mpvl/errd"
)

// ProtonLang ...
func ProtonLang(c *configuration.C) (*model.M, error) {
	var m *model.M
	return m, errd.Run(func(e *errd.E) {
		file, err := os.Open(suffix(*c.GetStringP("File"), ".proton"))
		e.Defer(file.Close)
		e.Must(err)

		//bytes, err := ioutil.ReadAll(file)
		//e.Must(err)

		//m, err = protonlang.Parse(string(bytes))
		//e.Must(err)
	})
}
