package db

import (
	"github.com/dlclark/regexp2"
)

var bannedRegexesDb = database[string]{
	filename: "regexes.json",
}

func init() {
	bannedRegexesDb.load()
}

func BanRegex(regex string) error {
	_, err := regexp2.Compile(regex, regexp2.None)
	if err == nil {
		bannedRegexesDb.add(regex)
	}
	return err
}

func GetRegexes() []string {
	return bannedRegexesDb.data
}

func UnbanRegex(regex string) {
	bannedRegexesDb.removeNotMatching(func(r string) bool { return regex != r })
}
