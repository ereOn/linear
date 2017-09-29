package command

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// GetCommandName gets a command name, if it has one.
func GetCommandName(prefix, binary string) string {
	r := getCommandNameRegexp(prefix)

	if matches := r.FindStringSubmatch(binary); len(matches) > 0 {
		return matches[1]
	}

	return ""
}

// A Command that can be executed.
type Command struct {
	Name        string `json:"-"`
	Path        string `json:"-"`
	Description string `json:"description"`
}

var (
	// DefaultScanCommandPrefix is the prefix for command scans.
	DefaultScanCommandPrefix = "linear"
	// DefaultScanEnvironmentVariables is the default environment variables to scan for.
	DefaultScanEnvironmentVariables = []string{"PATH"}
)

// A Scanner scans for commands.
type Scanner struct {
	CommandPrefix        string
	EnvironmentVariables []string
}

// Scan for commands.
func (s Scanner) Scan() (result []Command, errors []error) {
	environmentVariables := s.EnvironmentVariables

	if environmentVariables == nil {
		environmentVariables = DefaultScanEnvironmentVariables
	}

	for _, environmentVariable := range environmentVariables {
		paths := strings.Split(os.Getenv(environmentVariable), string(os.PathListSeparator))

		for _, path := range paths {
			commands, errs := s.ScanPath(path)
			result = append(result, commands...)
			errors = append(errors, errs...)
		}
	}

	return
}

// ScanPath scans a given path for commands.
func (s Scanner) ScanPath(path string) (result []Command, errors []error) {
	prefix := s.CommandPrefix

	if prefix == "" {
		prefix = DefaultScanCommandPrefix
	}

	if files, err := ioutil.ReadDir(path); err == nil {
		for _, file := range files {
			if !file.IsDir() {
				if name := GetCommandName(prefix, file.Name()); name != "" {
					if command, err := FromBinary(name, filepath.Join(path, file.Name())); err == nil {
						result = append(result, command)
					} else {
						errors = append(errors, err)
					}
				}
			}
		}
	}

	return
}

// FromBinary tries to build a Command from a binary at the specified path.
func FromBinary(name, path string) (Command, error) {
	cmd := exec.Command(path, "describe")
	r, err := cmd.StdoutPipe()

	if err != nil {
		return Command{}, err
	}

	defer r.Close()

	if err = cmd.Start(); err != nil {
		return Command{}, err
	}

	decoder := json.NewDecoder(r)

	result := Command{
		Name: name,
		Path: path,
	}

	if err = decoder.Decode(&result); err != nil {
		return Command{}, err
	}

	if err = cmd.Wait(); err != nil {
		return Command{}, err
	}

	return result, nil
}

// AsCobraCommand constructs a Cobra command from the Command.
func (c Command) AsCobraCommand() *cobra.Command {
	return &cobra.Command{
		Use:                c.Name,
		Short:              c.Description,
		DisableFlagParsing: true,
		SilenceUsage:       true,
		SilenceErrors:      true,
		RunE: func(cmd *cobra.Command, args []string) error {
			p := exec.Command(c.Path, args...)
			p.Stderr = cmd.OutOrStderr()
			p.Stdout = cmd.OutOrStdout()
			return p.Run()
		},
	}
}
