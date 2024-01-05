package db

type filterToggle struct {
	ID      string `json:"Id"`
	Enabled bool
}

type FilterToggleDB struct{ database[filterToggle] }

var filterToggleDB = FilterToggleDB{database[filterToggle]{filename: "filter-toggles.json"}}

func init() {
	filterToggleDB.load()
}

func GetFilterToggleDB() *FilterToggleDB {
	return &filterToggleDB
}

func (db *FilterToggleDB) IsFilterEnabled(id string) bool {
	toggle, found := db.first(func(item filterToggle) bool { return item.ID == id })
	if !found {
		// To be compatible with the previous version that had all filters enabled, return true if no toggle is found
		return true
	}

	return toggle.Enabled
}

func (db *FilterToggleDB) SetFilterEnabled(id string, enabled bool) {
	db.addOrReplace(
		filterToggle{ID: id, Enabled: enabled},
		func(a, b filterToggle) bool { return a.ID == b.ID },
	)
}
