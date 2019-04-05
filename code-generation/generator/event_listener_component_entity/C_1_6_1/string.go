package generator

import proton "github.com/SirMetathyst/go-proton"

// componentID ...
func componentID(c *proton.C, cp *proton.CP) proton.String {
	var eventTypeSuffix = ""
	if cp.EventType() == proton.RemovedEvent {
		eventTypeSuffix = "Removed"
	}
	var optionalContextID = ""
	if len(cp.ContextSlice()) > 1 {
		optionalContextID = c.ID().String()
	}
	componentID := optionalContextID + cp.ID().WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix + "Listener"
	return proton.String(componentID)
}
