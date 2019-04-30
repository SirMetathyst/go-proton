package main

// comment out generators and post-processors you don't want in the build
// the import order dictates the order it will generate/post-process in

import (
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/component"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/component_context"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/component_entity"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/component_entity_interface"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/component_lookup"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/component_matcher"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/composite_system"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/context"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/context_matcher"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/context_observer"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/contexts"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/entity"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/entity_index"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/event_listener_component_entity"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/event_listener_interface"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/event_system"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/event_systems"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/feature"
	_ "github.com/SirMetathyst/go-proton/code-generation/postprocessor/clean_target_directory/C_1_4_2"
	_ "github.com/SirMetathyst/go-proton/code-generation/postprocessor/file_header/C_1_4_2"
	_ "github.com/SirMetathyst/go-proton/code-generation/postprocessor/merge_content/C_1_4_2"
	_ "github.com/SirMetathyst/go-proton/code-generation/postprocessor/print_file/C_1_4_2"
	_ "github.com/SirMetathyst/go-proton/code-generation/postprocessor/print_file_content/C_1_4_2"
	_ "github.com/SirMetathyst/go-proton/code-generation/postprocessor/write_to_disk/C_1_4_2"

	"github.com/SirMetathyst/go-proton/proton/cmd"
)

func main() {
	cmd.Execute()
}
