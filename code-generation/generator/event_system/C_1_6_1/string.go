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

// methodID ...
func methodID(cp *proton.Component) proton.String {
	var eventTypeSuffix = ""
	if cp.EventType() == proton.EventTypeRemoved {
		eventTypeSuffix = "Removed"
	}
	componentID := cp.ID().WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix
	return proton.String(componentID)
}

// filter ...
func filter(c *proton.Context, cp *proton.Component) string {
	filter := ""
	if len(cp.MemberSlice()) == 0 {
		//switch cp.EventType() {
		//case proton.AddedEvent:
		//	filter = "entity." + cp.FlagPrefixOrDefault().String() + cp.ID().WithoutComponentSuffix().ToUpperFirst().String()
		//	break
		//case proton.RemovedEvent:
		//	filter = "!entity." + cp.FlagPrefixOrDefault().String() + cp.ID().WithoutComponentSuffix().ToUpperFirst().String()
		//	break
		//}
	} else {
		switch cp.EventType() {
		case proton.EventTypeAdded:
			filter = "entity.has" + cp.ID().WithoutComponentSuffix().ToUpperFirst().String() + " && "
			break
		case proton.EventTypeRemoved:
			filter = "!entity.has" + cp.ID().WithoutComponentSuffix().ToUpperFirst().String() + " && "
			break
		}
	}
	if cp.EventTarget() == proton.EventTargetSelf {
		filter += "entity.has" + componentID(c, cp).ToUpperFirst().String()
	}
	return filter
}
