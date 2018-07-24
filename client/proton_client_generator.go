package client

import (
	"github.com/SirMetathyst/go-entitas"
	proton "github.com/SirMetathyst/go-proton"
	. "github.com/SirMetathyst/go-proton/generator/component/C_1_7_0"
	. "github.com/SirMetathyst/go-proton/generator/component/E_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/component_context/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/component_entity/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/component_entity_interface/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/component_lookup/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/component_matcher/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/composite_system/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/context/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/context_attribute/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/context_matcher/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/context_matcher/E_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/context_observer/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/contexts/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/entity/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/entity_index/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/event_listener_component_entity/C_1_6_1"
	. "github.com/SirMetathyst/go-proton/generator/event_listener_interface/C_1_6_1"
	. "github.com/SirMetathyst/go-proton/generator/event_system/C_1_6_1"
	. "github.com/SirMetathyst/go-proton/generator/event_systems/C_1_6_1"
	. "github.com/SirMetathyst/go-proton/generator/feature/C_1_4_2"
)

// EG ...
type EG func(md *entitas.MD) ([]entitas.FI, error)

// middleware ...
func middleware(eg EG) proton.G {
	return func(v interface{}) ([]interface{}, error) {
		slice := make([]interface{}, 0)
		md := v.(*entitas.MD)
		r, err := eg(md)
		if err != nil {
			return nil, err
		}
		for _, cr := range r {
			slice = append(slice, cr)
		}
		return slice, err
	}
}

// CSharpEntitasGenerator ...
func CSharpEntitasGenerator(p *proton.P) {

	/* Generator(s).*/
	p.AddGenerator("CSharpComponentGenerator_C_1_7_0", middleware(ComponentGenerator_C_1_7_0), false)
	p.AddGenerator("CSharpComponentGenerator_E_1_4_2", middleware(ComponentGenerator_E_1_4_2), true)
	p.AddGenerator("CSharpComponentContextGenerator_C_1_4_2", middleware(ComponentContextGenerator_C_1_4_2), true)
	p.AddGenerator("CSharpComponentEntityGenerator_C_1_4_2", middleware(ComponentEntityGenerator_C_1_4_2), true)
	p.AddGenerator("CSharpComponentEntityInterfaceGenerator_C_1_4_2", middleware(ComponentEntityInterfaceGenerator_C_1_4_2), true)
	p.AddGenerator("CSharpComponentLookupGenerator_C_1_4_2", middleware(ComponentLookupGenerator_C_1_4_2), true)
	p.AddGenerator("CSharpComponentMatcherGenerator_C_1_4_2", middleware(ComponentMatcherGenerator_C_1_4_2), true)
	p.AddGenerator("CSharpContextGenerator_C_1_4_2", middleware(ContextGenerator_C_1_4_2), true)
	p.AddGenerator("CSharpContextAttributeGenerator_C_1_4_2", middleware(ContextAttributeGenerator_C_1_4_2), true)
	p.AddGenerator("CSharpContextMatcherGenerator_C_1_4_2", middleware(ContextMatcherGenerator_C_1_4_2), true)
	p.AddGenerator("CSharpContextMatcherGenerator_E_1_4_2", middleware(ContextMatcherGenerator_E_1_4_2), false)
	p.AddGenerator("CSharpContextsGenerator_C_1_4_2", middleware(ContextsGenerator_C_1_4_2), true)
	p.AddGenerator("CSharpContextObserverGenerator_C_1_4_2", middleware(ContextObserverGenerator_C_1_4_2), true)
	p.AddGenerator("CSharpEntityGenerator_C_1_4_2", middleware(EntityGenerator_C_1_4_2), true)
	p.AddGenerator("CSharpEntityIndexGenerator_C_1_4_2", middleware(EntityIndexGenerator_C_1_4_2), true)
	p.AddGenerator("CSharpFeatureGenerator_C_1_4_2", middleware(FeatureGenerator_C_1_4_2), true)
	p.AddGenerator("CSharpEventListenerComponentEntityGenerator_C_1_6_1", middleware(EventListenerComponentEntityGenerator_C_1_6_1), true)
	p.AddGenerator("CSharpEventListenerInterfaceGenerator_C_1_6_1", middleware(EventListenerInterfaceGenerator_C_1_6_1), true)
	p.AddGenerator("CSharpEventSystemGenerator_C_1_6_1", middleware(EventSystemGenerator_C_1_6_1), true)
	p.AddGenerator("CSharpEventSystemsGenerator_C_1_6_1", middleware(EventSystemsGenerator_C_1_6_1), true)
	p.AddGenerator("CSharpCompositeSystemGenerator_C_1_6_1", middleware(CompositeSystemGenerator_C_1_4_2), true)
}
