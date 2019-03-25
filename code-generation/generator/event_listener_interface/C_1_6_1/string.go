package generator

import "github.com/SirMetathyst/go-entitas"

// componentID ...
func componentID(c *entitas.C, cp *entitas.CP) entitas.String {
	var eventTypeSuffix = ""
	if cp.EventType() == entitas.RemovedEvent {
		eventTypeSuffix = "Removed"
	}
	var optionalContextID = ""
	if len(cp.ContextSlice()) > 1 {
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