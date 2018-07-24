package generator

import "github.com/SirMetathyst/go-entitas"

// componentID ...
func componentID(c *entitas.C, cp *entitas.CP) entitas.String {
	var eventTypeSuffix = ""
	if cp.EventType() == entitas.RemovedEvent {
		eventTypeSuffix = "Removed"
	}
	var optionalContextID = ""
	if len(cp.ContextList()) > 1 {
		optionalContextID = c.ID().String()
	}
	componentID := optionalContextID + cp.ID().WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix + "Listener"
	return entitas.String(componentID)
}

// methodID ...
func methodID(cp *entitas.CP) entitas.String {
	var eventTypeSuffix = ""
	if cp.EventType() == entitas.RemovedEvent {
		eventTypeSuffix = "Removed"
	}
	componentID := cp.ID().WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix
	return entitas.String(componentID)
}

// filter ...
func filter(c *entitas.C, cp *entitas.CP) string {
	filter := ""
	if len(cp.MemberList()) == 0 {
		switch cp.EventType() {
		case entitas.AddedEvent:
			filter = "entity." + cp.FlagPrefixOrDefault().String() + cp.ID().WithoutComponentSuffix().ToUpperFirst().String()
			break
		case entitas.RemovedEvent:
			filter = "!entity." + cp.FlagPrefixOrDefault().String() + cp.ID().WithoutComponentSuffix().ToUpperFirst().String()
			break
		}
	} else {
		switch cp.EventType() {
		case entitas.AddedEvent:
			filter = "entity.has" + cp.ID().WithoutComponentSuffix().ToUpperFirst().String()
			break
		case entitas.RemovedEvent:
			filter = "!entity.has" + cp.ID().WithoutComponentSuffix().ToUpperFirst().String()
			break
		}
	}
	if cp.EventTarget() == entitas.SelfTarget {
		filter += " && entity.has" + componentID(c, cp).ToUpperFirst().String()
	}
	return filter
}
