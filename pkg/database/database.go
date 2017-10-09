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

// ComponentID represents a component ID.
type ComponentID struct {
	Name string
	Type string
}

// Equals compares two component ids.
func (i ComponentID) Equals(id ComponentID) bool {
	return i.Name == id.Name && i.Type == id.Type
}

func (i ComponentID) String() string {
	return fmt.Sprintf("%s:%s", i.Type, i.Name)
}

// Component represents a project component.
type Component struct {
	ID         ComponentID
	Version    string
	Properties interface{}
}

// DecodeProperties decodes the component properties as the specified structure.
func (c Component) DecodeProperties(v interface{}) error {
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
		if component.ID.Type == t {
			components = append(components, component)
		}
	}

	return
}

// Find returns a component of the specified type and name.
func (d Database) Find(id ComponentID) *Component {
	for _, component := range d.Components {
		if component.ID.Equals(id) {
			return &component
		}
	}

	return nil
}

// Add a new component to the database.
func (d *Database) Add(component Component) error {
	if d.Find(component.ID) != nil {
		return fmt.Errorf("a component of type `%s` with the name `%s` already exists", component.ID.Type, component.ID.Name)
	}

	d.Components = append(d.Components, component)

	return nil
}

// Remove a component from the database. If found, the removed component is returned.
func (d *Database) Remove(id ComponentID) *Component {
	for i, component := range d.Components {
		if component.ID.Equals(id) {
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
