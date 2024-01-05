package filters

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

type Filter interface {
	IsMessageAllowed(ctx helpers.ResponseContext) bool
	ScrutinyModeOnly() bool
	ShouldSuppressMock() bool
	Description() Description
}

type Description struct {
	ID, Name, Desc string
}

// Use this notice in Filter.Description to indicate that a certain filter is unstable.
const unstableNotice = "\n      _Этот фильтр экспериментален и может работать нестабильно, " +
	"с большим количеством ложных срабатываний. Не забывайте про команду /filteroff._"

var filters []Filter

func registerFilter(filter Filter) {
	filters = append(filters, filter)
}

func List() []Filter {
	return filters
}

func ValidID(id string) bool {
	for _, f := range filters {
		if f.Description().ID == id {
			return true
		}
	}

	return false
}

func IsMessageAllowed(ctx helpers.ResponseContext) (allowed, suppressMock bool) {
	senderIsUnderScrutiny := db.GetScrutinyDB().IsUnderScrutiny(ctx.Message.From.UserName)

	for _, filter := range filters {
		isFilterEnabled := db.GetFilterToggleDB().IsFilterEnabled(filter.Description().ID)
		if !isFilterEnabled || (filter.ScrutinyModeOnly() && !senderIsUnderScrutiny) {
			continue
		}

		if !filter.IsMessageAllowed(ctx) {
			return false, filter.ShouldSuppressMock()
		}
	}

	return true, false
}
