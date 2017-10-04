package database

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// Database represents a project datase.
type Database struct {
	Components []Component
}

// Component represents a project component.
type Component struct {
	Name       string
	Type       string
	Properties map[string]interface{}
}

// As decodes the component as the specified structure.
func (c Component) As(v interface{}) error {
	return mapstructure.Decode(c.Properties, v)
}

// FromFile loads a Database from a file on disk.
func FromFile(path string) (*Database, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	return FromReader(f)
}

// FromReader loads a Database from a reader.
func FromReader(r io.Reader) (*Database, error) {
	data, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	var database Database

	if err := yaml.Unmarshal(data, &database); err != nil {
		return nil, err
	}

	return &database, nil
}

// ByType returns all the components of the specified type.
func (d Database) ByType(t string) (components []Component) {
	for _, component := range d.Components {
		if component.Type == t {
			components = append(components, component)
		}
	}

	return
}

// Find returns a component of the specified type and name.
func (d Database) Find(t, name string) *Component {
	for _, component := range d.Components {
		if component.Type == t && component.Name == name {
			return &component
		}
	}

	return nil
}

// Add a new component to the database.
func (d *Database) Add(component Component) error {
	if d.Find(component.Type, component.Name) != nil {
		return fmt.Errorf("a component of type `%s` with the name `%s` already exists", component.Type, component.Name)
	}

	d.Components = append(d.Components, component)

	return nil
}

// Remove a component from the database. If found, the removed component is returned.
func (d *Database) Remove(t, name string) *Component {
	for i, component := range d.Components {
		if component.Type == t && component.Name == name {
			d.Components = append(d.Components[:i], d.Components[i+1:]...)

			return &component
		}
	}

	return nil
}

// ToWriter writes the database to the specified writer.
func (d Database) ToWriter(w io.Writer) error {
	data, err := yaml.Marshal(d)

	if err != nil {
		return err
	}

	_, err = w.Write(data)

	return err
}

// ToFile writes a Database to file.
func (d Database) ToFile(path string) error {
	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	return d.ToWriter(f)
}
