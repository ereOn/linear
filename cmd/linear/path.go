package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ereOn/linear/pkg/command"
	"github.com/spf13/cobra"
)

type commandBinary struct {
	Name string
	Path string
}

func getSubcommandBinaries(path string) (result []commandBinary) {
	components := strings.Split(path, string(os.PathListSeparator))

	set := map[string]bool{}

	for _, component := range components {
		if files, err := ioutil.ReadDir(component); err == nil {
			for _, file := range files {
				if !file.IsDir() {
					matches := subcommandBinaryPattern.FindStringSubmatch(file.Name())

					if len(matches) > 0 {
						if !set[file.Name()] {
							set[file.Name()] = true
							result = append(result, commandBinary{
								Name: matches[1],
								Path: filepath.Join(component, file.Name()),
							})
						}
					}
				}
			}
		}

	}

	return
}

func getSubcommands(binaries []commandBinary) (result []*cobra.Command) {
	for _, binary := range binaries {
		if cmd, err := describe(binary); err == nil {
			result = append(result, cmd)
		}
	}

	return
}

func describe(binary commandBinary) (*cobra.Command, error) {
	cmd := exec.Command(binary.Path, "describe")
	r, err := cmd.StdoutPipe()

	if err != nil {
		return nil, err
	}

	defer r.Close()

	if err = cmd.Start(); err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(r)

	var description command.Description

	if err = decoder.Decode(&description); err != nil {
		return nil, err
	}

	if err = cmd.Wait(); err != nil {
		return nil, err
	}

	return &cobra.Command{
		Use:                binary.Name,
		Short:              description.Short,
		DisableFlagParsing: true,
		SilenceUsage:       true,
		SilenceErrors:      true,
		RunE: func(cmd *cobra.Command, args []string) error {
			c := exec.Command(binary.Path, args...)
			c.Stderr = cmd.OutOrStderr()
			c.Stdout = cmd.OutOrStdout()
			return c.Run()
		},
	}, nil
}
