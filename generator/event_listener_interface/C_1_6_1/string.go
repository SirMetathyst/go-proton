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
