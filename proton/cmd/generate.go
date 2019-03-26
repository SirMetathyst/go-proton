package cmd

import (
	. "github.com/SirMetathyst/go-proton/code-generation"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/clean_target_directory/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/file_header/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/merge_content/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/print_file/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/print_file_content/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/write_to_disk/C_1_4_2"

	dsl "github.com/SirMetathyst/go-proton/dsl"

	"github.com/spf13/cobra"
)

var (
	Directory = "./"
	Daemonize = false
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate entitas source code",
	Run: func(cmd *cobra.Command, args []string) {
		RunApplication(Directory, Daemonize, dsl.Parse)
	},
}

func init() {
	for _, generatorInfo := range GeneratorInfo() {
		generateCmd.PersistentFlags().BoolVar(&generatorInfo.Enabled, generatorInfo.GeneratorVersion, generatorInfo.Enabled, "")
	}
	for _, postProcessorInfo := range PostProcessorInfo() {
		generateCmd.PersistentFlags().BoolVar(&postProcessorInfo.Enabled, postProcessorInfo.PostProcessorVersion, postProcessorInfo.Enabled, "")
	}

	generateCmd.PersistentFlags().StringVarP(&Directory, "project", "p", Directory, "")
	generateCmd.PersistentFlags().BoolVarP(&Daemonize, "daemonize", "d", Daemonize, "")

	AddPostProcessor("MergeContentPostProcessor_C_1_4_2", MergeContentPostProcessor_C_1_4_2, true)
	AddPostProcessor("FileHeaderPostProcessor_C_1_4_2", FileHeaderPostProcessor_C_1_4_2, true)
	AddPostProcessor("PrintFilePostProcessor_C_1_4_2", PrintFilePostProcessor_C_1_4_2, false)
	AddPostProcessor("PrintFileContentPostProcessor_C_1_4_2", PrintFileContentPostProcessor_C_1_4_2, false)
	AddPostProcessor("CleanTargetDirectoryPostProcessor_C_1_4_2", CleanTargetDirectoryPostProcessor_C_1_4_2, true)
	AddPostProcessor("WriteToDiskPostProcessor_C_1_4_2", WriteToDiskPostProcessor_C_1_4_2, true)

}
