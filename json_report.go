package main

import (
	"encoding/json"
	"maps"
	"os"
	"slices"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	keys := slices.Collect(maps.Keys(pages))
	slices.Sort(keys)

	data, err := json.MarshalIndent(pages, "", "  ")

	if err != nil {
		return err
	}

	os.WriteFile(filename, data, 0644)
	return nil
}