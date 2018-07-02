package configurationprovider

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/SirMetathyst/proton/configuration"
)

type stringSlice []string

func (s *stringSlice) String() string {
	return fmt.Sprint(*s)
}

func (s *stringSlice) Set(value string) error {
	if len(*s) > 0 {
		return errors.New("string slice flag already set")
	}
	for _, v := range strings.Split(value, ",") {
		*s = append(*s, v)
	}
	return nil
}

func newStringSlice(value []string, p *[]string) *stringSlice {
	*p = value
	return (*stringSlice)(p)
}

// usage ...
func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "[Flags]\n")
	c := 0
	flag.VisitAll(func(f *flag.Flag) {
		l := len(fmt.Sprintf("  --%s", f.Name))
		if l > c {
			c = l
		}
	})
	flag.VisitAll(func(f *flag.Flag) {
		s := fmt.Sprintf("  --%s", f.Name)
		name, _ := flag.UnquoteUsage(f)
		if len(name) > 0 {
			s += " " + name
		}
		diff := c - len(s)
		if diff < 0 {
			diff = 0
		}
		s += fmt.Sprintf(strings.Repeat(" ", diff)+"     (default %v)", f.DefValue)

		fmt.Fprint(flag.CommandLine.Output(), s, "\n")
	})
}

// Flag ...
func Flag(c *configuration.C) error {

	for _, opt := range c.GetAllBool() {
		flag.BoolVar(opt.Value, opt.Key, *opt.Value, "")
	}

	for _, opt := range c.GetAllStringSlice() {
		flag.Var(newStringSlice(*opt.Value, opt.Value), opt.Key, "")
	}

	for _, opt := range c.GetAllString() {
		flag.StringVar(opt.Value, opt.Key, *opt.Value, "")
	}

	flag.Usage = usage
	flag.Parse()
	return nil
}
