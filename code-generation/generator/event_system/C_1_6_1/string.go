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

// methodID ...
func methodID(cp *proton.CP) proton.String {
	var eventTypeSuffix = ""
	if cp.EventType() == proton.RemovedEvent {
		eventTypeSuffix = "Removed"
	}
	componentID := cp.ID().WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix
	return proton.String(componentID)
}

// filter ...
func filter(c *proton.C, cp *proton.CP) string {
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
		case proton.AddedEvent:
			filter = "entity.has" + cp.ID().WithoutComponentSuffix().ToUpperFirst().String() + " && "
			break
		case proton.RemovedEvent:
			filter = "!entity.has" + cp.ID().WithoutComponentSuffix().ToUpperFirst().String() + " && "
			break
		}
	}
	if cp.EventTarget() == proton.SelfTarget {
		filter += "entity.has" + componentID(c, cp).ToUpperFirst().String()
	}
	return filter
}
