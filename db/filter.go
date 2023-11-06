package db

type filterToggle struct {
	Id      string
	Enabled bool
}

var filterToggleDb = database[filterToggle]{filename: "filter-toggles.json"}

func init() {
	filterToggleDb.load()
}

func IsFilterEnabled(id string) bool {
	toggle, found := filterToggleDb.first(func(item filterToggle) bool { return item.Id == id })
	if !found {
		// To be compatible with the previous version that had all filters enabled, return true if no toggle is found
		return true
	}
	return toggle.Enabled
}

func SetFilterEnabled(id string, enabled bool) {
	filterToggleDb.addOrReplace(
		filterToggle{Id: id, Enabled: enabled},
		func(a, b filterToggle) bool { return a.Id == b.Id },
	)
}
