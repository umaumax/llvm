// Code generated by "string2enum -linecomment -type ClauseType ../../ir/enum"; DO NOT EDIT.

package enum

import (
	"fmt"

	"github.com/umaumax/llvm/ir/enum"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the string2enum command to generate them again.
	var x [1]struct{}
	_ = x[enum.ClauseTypeCatch-1]
	_ = x[enum.ClauseTypeFilter-2]
}

const _ClauseType_name = "catchfilter"

var _ClauseType_index = [...]uint8{0, 5, 11}

// ClauseTypeFromString returns the ClauseType enum corresponding to s.
func ClauseTypeFromString(s string) enum.ClauseType {
	if len(s) == 0 {
		return 0
	}
	for i := range _ClauseType_index[:len(_ClauseType_index)-1] {
		if s == _ClauseType_name[_ClauseType_index[i]:_ClauseType_index[i+1]] {
			return enum.ClauseType(i + 1)
		}
	}
	panic(fmt.Errorf("unable to locate ClauseType enum corresponding to %q", s))
}
