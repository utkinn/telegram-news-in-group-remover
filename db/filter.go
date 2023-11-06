package db

type filterToggle struct {
	Name    string
	Enabled bool
}

var filterToggleDb = database[filterToggle]{filename: "filter-toggles.json"}

func init() {
	filterToggleDb.load()
}

func IsFilterEnabled(name string) bool {
	toggle, found := filterToggleDb.first(func(item filterToggle) bool { return item.Name == name })
	if !found {
		// To be compatible with the previous version that had all filters enabled, return true if no toggle is found
		return true
	}
	return toggle.Enabled
}

func SetFilterEnabled(name string, enabled bool) {
	filterToggleDb.addOrReplace(
		filterToggle{Name: name, Enabled: enabled},
		func(a, b filterToggle) bool { return a.Name == b.Name },
	)
}
