package db

type filterToggle struct {
	Id      string
	Enabled bool
}

type FilterToggleDB struct{ database[filterToggle] }

var filterToggleDb = FilterToggleDB{database[filterToggle]{filename: "filter-toggles.json"}}

func init() {
	filterToggleDb.load()
}

func GetFilterToggleDB() *FilterToggleDB {
	return &filterToggleDb
}

func (db *FilterToggleDB) IsFilterEnabled(id string) bool {
	toggle, found := db.first(func(item filterToggle) bool { return item.Id == id })
	if !found {
		// To be compatible with the previous version that had all filters enabled, return true if no toggle is found
		return true
	}
	return toggle.Enabled
}

func (db *FilterToggleDB) SetFilterEnabled(id string, enabled bool) {
	db.addOrReplace(
		filterToggle{Id: id, Enabled: enabled},
		func(a, b filterToggle) bool { return a.Id == b.Id },
	)
}
