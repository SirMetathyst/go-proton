package generator

import proton "github.com/SirMetathyst/go-proton"

// componentID ...
func componentID(c *proton.Context, cp *proton.Component) proton.String {
	var eventTypeSuffix = ""
	if cp.EventType() == proton.EventTypeRemoved {
		eventTypeSuffix = "Removed"
	}
	var optionalContextID = ""
	if len(cp.ContextSlice()) > 1 {
		optionalContextID = c.ID().String()
	}
	componentID := optionalContextID + cp.ID().WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix + "Listener"
	return proton.String(componentID)
}
