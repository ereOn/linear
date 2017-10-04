package database

import "testing"

type testComponentProperties struct {
	Location string
}

func TestFromFile(t *testing.T) {
	database, err := FromFile("fixtures/database.yml")

	if err != nil {
		t.Errorf("no error was expected but got: %s", err)
	}

	if len(database.Components) != 2 {
		t.Errorf("expected %d components but got %d", 2, len(database.Components))
	}

	component := database.Find("service", "foo")

	if component == nil {
		t.Error("component should not be nil")
	}

	var properties testComponentProperties

	if err = component.As(&properties); err != nil {
		t.Errorf("no error was expected but got: %s", err)
	}
}

func TestFromFileNonExisting(t *testing.T) {
	_, err := FromFile("fixtures/non-existing.yml")

	if err == nil {
		t.Error("an error was expected")
	}
}
