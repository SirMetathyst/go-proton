package client

import (
	"encoding/json"
	"io/ioutil"

	"github.com/SirMetathyst/go-blackboard"
	"github.com/SirMetathyst/go-entitas"
	"github.com/SirMetathyst/go-entitas/builder"
)

// jsonContext ...
type jsonContext struct {
	ID string `json:"id"`
}

// jsonComponent ...
type jsonComponent struct {
	ID            string       `json:"id"`
	FlagPrefix    string       `json:"flag_prefix"`
	Unique        bool         `json:"unique"`
	EventTarget   string       `json:"event_target"`
	EventType     string       `json:"event_type"`
	EventPriority int          `json:"event_priority"`
	CleanupMode   string       `json:"cleanup_mode"`
	Contexts      []string     `json:"contexts"`
	Members       []jsonMember `json:"members"`
}

// jsonEntityIndex ...
type jsonEntityIndex struct {
	ID      string       `json:"id"`
	Primary bool         `json:"primary`
	Context string       `json:"context"`
	Methods []jsonMethod `json:"methods"`
}

// jsonMethod ...
type jsonMethod struct {
	ID      string       `json:"id"`
	Members []jsonMember `json:"members"`
}

// jsonMember ...
type jsonMember struct {
	ID          string `json:"id"`
	Value       string `json:"value"`
	EntityIndex string `json:"entity_index"`
}

// jsonModel ...
type jsonModel struct {
	Namespace      string            `json:"namespace"`
	Contexts       []jsonContext     `json:"contexts"`
	DefaultContext string            `json:"default_context"`
	Components     []jsonComponent   `json:"components"`
	EntityIndex    []jsonEntityIndex `json:"entity_index"`
}

func getEventTarget(b string) entitas.EventTarget {
	if b == "self" {
		return entitas.SelfTarget
	} else if b == "any" {
		return entitas.AnyTarget
	}
	return entitas.NoTarget
}

func getEventType(e string) entitas.EventType {
	if e == "removed" {
		return entitas.RemovedEvent
	}
	return entitas.AddedEvent
}

func getEntityIndex(e string) entitas.EntityIndex {
	if e == "multiple" {
		return entitas.MultipleEntityIndex
	} else if e == "single" {
		return entitas.SingleEntityIndex
	}
	return entitas.NoEntityIndex
}

func getCleanupMode(m string) entitas.CleanupMode {
	if m == "destroy_entity" {
		return entitas.DestroyEntity
	} else if m == "remove_component" {
		return entitas.RemoveComponent
	}
	return entitas.NoCleanup
}

// componentID ...
func componentID(c string, cp jsonComponent) string {
	var eventTypeSuffix = ""
	if getEventType(cp.EventType) == entitas.RemovedEvent {
		eventTypeSuffix = "Removed"
	}
	var optionalContextID = ""
	if len(cp.Contexts) > 1 {
		optionalContextID = c
	}
	componentID := optionalContextID + entitas.String(cp.ID).WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix + "Listener"
	return componentID
}

func createEventComponent(mdb *builder.MDB, default_context string, cp jsonComponent) error {
	g := func(c string, cp jsonComponent) error {
		cpb := mdb.NewComponent()
		cpb.SetID(componentID(c, cp) + "Component")
		cpb.SetListener(true)
		cpb.AddContext(c)

		err := cpb.NewMember().
			SetID("value").
			SetValue("System.Collections.Generic.List<I" + componentID(c, cp) + ">").
			Build()

		if err != nil {
			return err
		}
		err = cpb.Build()
		if err != nil {
			return err
		}
		return nil
	}

	if len(cp.Contexts) > 0 {
		for _, c := range cp.Contexts {
			err := g(c, cp)
			if err != nil {
				return err
			}
		}
	} else {
		err := g(default_context, cp)
		if err != nil {
			return err
		}
	}

	return nil
}

// JSON ...
func JSON(bb *blackboard.BB) (*entitas.MD, error) {
	jm := jsonModel{}

	raw, _ := ioutil.ReadFile(suffix(*bb.StringP("File"), ".json"))

	err := json.Unmarshal(raw, &jm)
	if err != nil {
		return nil, err
	}

	mdb := builder.NewModelBuilder()
	mdb.SetNamespace(jm.Namespace)

	for _, c := range jm.Contexts {
		err = mdb.NewContext().
			SetID(c.ID).
			Build()

		if err != nil {
			return nil, err
		}
	}

	mdb.SetDefaultContext(jm.DefaultContext)

	for _, cp := range jm.Components {
		cpb := mdb.NewComponent().
			SetID(cp.ID).
			SetFlagPrefix(cp.FlagPrefix).
			SetUnique(cp.Unique).
			SetEventTarget(getEventTarget(cp.EventTarget)).
			SetEventType(getEventType(cp.EventType)).
			SetEventPriority(cp.EventPriority).
			SetCleanupMode(getCleanupMode(cp.CleanupMode))
		for _, c := range cp.Contexts {
			cpb.AddContext(c)
		}
		for _, m := range cp.Members {
			err := cpb.NewMember().
				SetID(m.ID).
				SetValue(m.Value).
				SetEntityIndex(getEntityIndex(m.EntityIndex)).
				Build()
			if err != nil {
				return nil, err
			}
		}
		err := cpb.Build()
		if err != nil {
			return nil, err
		}
		if getEventTarget(cp.EventTarget) > 0 {
			err = createEventComponent(mdb, jm.DefaultContext, cp)
			if err != nil {
				return nil, err
			}
		}
	}

	for _, ei := range jm.EntityIndex {
		eib := mdb.NewEntityIndex().
			SetID(ei.ID).
			SetPrimary(ei.Primary).
			AddContext(ei.Context)

		for _, eim := range ei.Methods {
			eimb := eib.NewMethod().
				SetID(eim.ID)

			for _, m := range eim.Members {
				err := eimb.NewMember().
					SetID(m.ID).
					SetValue(m.Value).
					SetEntityIndex(getEntityIndex(m.EntityIndex)).
					Build()

				if err != nil {
					return nil, err
				}
			}
			err := eimb.Build()
			if err != nil {
				return nil, err
			}
		}
		err := eib.Build()
		if err != nil {
			return nil, err
		}
	}

	md, err := mdb.Build()
	if err != nil {
		return nil, err
	}

	return md, nil
}
