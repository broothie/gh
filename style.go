package gh

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

type Style map[string]any

func (s Style) String() string {
	return fmt.Sprintf("%s;", strings.Join(lo.Map(lo.Entries(s), func(entry lo.Entry[string, any], _ int) string {
		return fmt.Sprintf("%s:%v", entry.Key, entry.Value)
	}), ";"))
}

type Class map[string]bool

func (c Class) String() string {
	return strings.Join(lo.FilterMap(lo.Entries(c), func(entry lo.Entry[string, bool], _ int) (string, bool) {
		return entry.Key, entry.Value
	}), " ")
}
