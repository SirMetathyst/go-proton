package cmd

import (
	"log"

	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
	language "github.com/SirMetathyst/go-proton/language"

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
		codegeneration.SetOptions(codegeneration.PO{
			ProjectPath:  ProjectPath,
			OutputFolder: OutputFolder,
		})
		if Daemonize {
			codegeneration.Daemon(language.Parse)
		} else {
			err := codegeneration.ParseGenerate(language.Parse)
			if err != nil {
				log.Printf("proton: %s\n", err)
				return
			}
			log.Println("proton: generated")
		}
	},
}

func init() {
	for _, generatorInfo := range codegeneration.GeneratorInfoSlice() {
		generateCmd.Flags().BoolVar(&generatorInfo.Enabled, generatorInfo.GeneratorVersion, generatorInfo.Enabled, "")
	}
	for _, postProcessorInfo := range codegeneration.PostProcessorInfoSlice() {
		generateCmd.Flags().BoolVar(&postProcessorInfo.Enabled, postProcessorInfo.PostProcessorVersion, postProcessorInfo.Enabled, "")
	}

	generateCmd.Flags().StringVarP(&ProjectPath, "project-path", "p", ProjectPath, "project path")
	generateCmd.Flags().StringVarP(&OutputFolder, "output-folder", "o", OutputFolder, "output folder")
	generateCmd.Flags().BoolVarP(&Daemonize, "daemonize", "d", Daemonize, "daemonize application")
}
