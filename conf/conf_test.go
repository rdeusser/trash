package conf

import (
	"testing"
)

func TestDuplicates(t *testing.T) {
	testData := []struct {
		imports    []Import
		duplicates int
	}{
		{[]Import{
			{Package: "package1", Version: "version1"},
		}, 0},
		{[]Import{
			{Package: "package1", Version: "version1"},
			{Package: "package2", Version: "version1", Repo: "repoA"},
		}, 0},
		{[]Import{
			{Package: "package1", Version: "version1"},
			{Package: "package2", Version: "version1", Repo: "repoA"},
			{Package: "package1", Version: "version1"},
		}, 1},
		{[]Import{
			{Package: "package1", Version: "version1"},
			{Package: "package2", Version: "version1", Repo: "repoA"},
			{Package: "package1", Version: "version1"},
			{Package: "package1", Version: "version1"},
		}, 2},
		{[]Import{
			{Package: "package1", Version: "version1"},
			{Package: "package2", Version: "version1", Repo: "repoA"},
			{Package: "package1", Version: "version1"},
			{Package: "package1", Version: "version1"},
			{Package: "package2", Version: "version2", Repo: "repoB"},
			{Package: "package3", Version: "version1", Repo: "repoA"},
		}, 3},
	}

	for i, d := range testData {
		trash := Conf{
			Package:   "",
			Imports:   d.imports,
			Excludes:  []string{},
			Packages:  []string{},
			ImportMap: make(map[string]Import),
			confFile:  "",
			yamlType:  false,
		}
		trash.Dedupe()

		if d.duplicates != len(d.imports)-len(trash.Imports) {
			t.Errorf("Case %d failed: expected %d duplicates but removed %d", i, d.duplicates, len(d.imports)-len(trash.Imports))
		}

	}

}
