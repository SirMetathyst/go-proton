package modelprovider

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/SirMetathyst/proton/configuration"
	"github.com/SirMetathyst/proton/model"
)

// jsonContext ...
type jsonContext struct {
	ID        string `json:"id"`
	IsDefault bool   `json:"is_default"`
}

// jsonComponent ...
type jsonComponent struct {
	ID           string       `json:"id"`
	Prefix       string       `json:"prefix"`
	Unique       bool         `json:"unique"`
	EventBinding string       `json:"event_binding"`
	EventType    string       `json:"event_type"`
	Context      []string     `json:"context"`
	Member       []jsonMember `json:"member"`
}

// jsonMember ...
type jsonMember struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// jsonModel ...
type jsonModel struct {
	Context   []jsonContext   `json:"context"`
	Component []jsonComponent `json:"component"`
}

func getEventBinding(b string) int {
	if b == "self" {
		return 1
	} else if b == "global" {
		return 2
	}
	return 0
}

func getEventType(e string) int {
	if e == "removed" {
		return 1
	}
	return 0
}

// Json ...
func Json(c *configuration.C) (*model.M, error) {
	m := model.NewModel()
	jm := jsonModel{}

	raw, _ := ioutil.ReadFile(suffix(*c.GetStringP("File"), ".json"))

	json.Unmarshal(raw, &jm)

	m = model.NewModel()

	for _, ctx := range jm.Context {
		log.Println(ctx)
		m.CreateContext(false, ctx.ID).SetDefault(ctx.IsDefault)
	}

	for _, c := range jm.Component {
		cc := m.CreateComponent(false, c.ID).
			SetPrefix(c.Prefix).
			SetUnique(c.Unique).
			SetEventBinding(getEventBinding(c.EventBinding)).
			SetEventType(getEventType(c.EventType))

		for _, ctx := range c.Context {

			cc.AddContext(false, m.GetContextWithID(ctx))

		}

		for _, m := range c.Member {
			cc.CreateMember(true, m.ID).SetValue(m.Value)
		}
	}

	m.Refresh()

	return m, nil
}
