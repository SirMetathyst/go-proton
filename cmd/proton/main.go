package main

// comment out generators and post-processors you don't want in the build

import (
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/component/E_1_4_2"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/component_context/C_1_4_2"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/component_entity/C_1_4_2"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/component_entity_interface/C_1_4_2"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/component_lookup/C_1_4_2"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/component_matcher/C_1_4_2"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/composite_system/C_1_4_2"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/context/C_1_4_2"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/context/C_1_9_0"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/context_matcher/C_1_4_2"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/context_observer/C_1_4_2"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/contexts/C_1_4_2"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/entity/C_1_4_2"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/entity_index/C_1_4_2"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/event_listener_component_entity/C_1_6_1"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/event_listener_interface/C_1_6_1"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/event_system/C_1_6_1"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/event_systems/C_1_6_1"
	//_ "github.com/SirMetathyst/go-proton/code-generation/generator/feature/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation"
	"github.com/SirMetathyst/go-proton/dsl"
)

func main() {
	RunApplication(dsl.Parse)
}
