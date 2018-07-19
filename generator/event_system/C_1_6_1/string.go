package generator

import "github.com/SirMetathyst/proton/model"

// ComponentID ...
func ComponentID(c *model.C, cp *model.CP) model.String {
	var eventTypeSuffix = ""
	if cp.EventType() == model.RemovedEvent {
		eventTypeSuffix = "Removed"
	}
	var optionalContextID = ""
	if len(cp.ContextList()) > 1 {
		optionalContextID = c.ID().String()
	}
	componentID := optionalContextID + cp.ID().WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix + "Listener"
	return model.String(componentID)
}

// MethodID ...
func MethodID(cp *model.CP) model.String {
	var eventTypeSuffix = ""
	if cp.EventType() == model.RemovedEvent {
		eventTypeSuffix = "Removed"
	}
	componentID := cp.ID().WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix
	return model.String(componentID)
}

// Filter ...
func Filter(c *model.C, cp *model.CP) string {
	filter := ""
	if len(cp.MemberList()) == 0 {
		switch cp.EventType() {
		case model.AddedEvent:
			filter = "entity." + cp.FlagPrefixOrDefault().String() + cp.ID().WithoutComponentSuffix().ToUpperFirst().String()
			break
		case model.RemovedEvent:
			filter = "!entity." + cp.FlagPrefixOrDefault().String() + cp.ID().WithoutComponentSuffix().ToUpperFirst().String()
			break
		}
	} else {
		switch cp.EventType() {
		case model.AddedEvent:
			filter = "entity.has" + cp.ID().WithoutComponentSuffix().ToUpperFirst().String()
			break
		case model.RemovedEvent:
			filter = "!entity.has" + cp.ID().WithoutComponentSuffix().ToUpperFirst().String()
			break
		}
	}
	if cp.EventTarget() == model.SelfTarget {
		filter += " && entity.has" + ComponentID(c, cp).ToUpperFirst().String()
	}
	return filter
}
