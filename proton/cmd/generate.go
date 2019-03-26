package cmd

import (
	. "github.com/SirMetathyst/go-proton/code-generation"
	dsl "github.com/SirMetathyst/go-proton/dsl"

	"github.com/spf13/cobra"
)

var (
	ProjectPath  = "./"
	OutputFolder = "src-gen"
	Daemonize    = false
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate entitas source code",
	Long: `you can enable/disable generators and post-processors with flags, 
change the project path (which is based on excuting directory by default), 
output folder where code is written to and keep the generator alive. 
This will find any *.proton files in the project path, track them for changes 
and re-generate the code.`,
	Run: func(cmd *cobra.Command, args []string) {
		RunApplication(ProjectPath, OutputFolder, Daemonize, dsl.Parse)
	},
}

func init() {
	for _, generatorInfo := range GeneratorInfo() {
		generateCmd.PersistentFlags().BoolVar(&generatorInfo.Enabled, generatorInfo.GeneratorVersion, generatorInfo.Enabled, "")
	}
	for _, postProcessorInfo := range PostProcessorInfo() {
		generateCmd.PersistentFlags().BoolVar(&postProcessorInfo.Enabled, postProcessorInfo.PostProcessorVersion, postProcessorInfo.Enabled, "")
	}

	generateCmd.PersistentFlags().StringVarP(&ProjectPath, "project-path", "p", ProjectPath, "project path")
	generateCmd.PersistentFlags().StringVarP(&OutputFolder, "output-folder", "o", OutputFolder, "output folder")
	generateCmd.PersistentFlags().BoolVarP(&Daemonize, "daemonize", "d", Daemonize, "daemonize application")
}
