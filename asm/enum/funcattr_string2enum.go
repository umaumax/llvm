// Code generated by "string2enum -linecomment -type FuncAttr ../../ir/enum"; DO NOT EDIT.

package enum

import (
	"fmt"

	"github.com/umaumax/llvm/ir/enum"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the string2enum command to generate them again.
	var x [1]struct{}
	_ = x[enum.FuncAttrAlwaysInline-0]
	_ = x[enum.FuncAttrArgMemOnly-1]
	_ = x[enum.FuncAttrBuiltin-2]
	_ = x[enum.FuncAttrCold-3]
	_ = x[enum.FuncAttrConvergent-4]
	_ = x[enum.FuncAttrInaccessibleMemOrArgMemOnly-5]
	_ = x[enum.FuncAttrInaccessibleMemOnly-6]
	_ = x[enum.FuncAttrInlineHint-7]
	_ = x[enum.FuncAttrJumpTable-8]
	_ = x[enum.FuncAttrMinSize-9]
	_ = x[enum.FuncAttrNaked-10]
	_ = x[enum.FuncAttrNoBuiltin-11]
	_ = x[enum.FuncAttrNoDuplicate-12]
	_ = x[enum.FuncAttrNoFree-13]
	_ = x[enum.FuncAttrNoImplicitFloat-14]
	_ = x[enum.FuncAttrNoInline-15]
	_ = x[enum.FuncAttrNonLazyBind-16]
	_ = x[enum.FuncAttrNoRecurse-17]
	_ = x[enum.FuncAttrNoRedZone-18]
	_ = x[enum.FuncAttrNoReturn-19]
	_ = x[enum.FuncAttrNoSync-20]
	_ = x[enum.FuncAttrNoUnwind-21]
	_ = x[enum.FuncAttrOptNone-22]
	_ = x[enum.FuncAttrOptSize-23]
	_ = x[enum.FuncAttrReadNone-24]
	_ = x[enum.FuncAttrReadOnly-25]
	_ = x[enum.FuncAttrReturnsTwice-26]
	_ = x[enum.FuncAttrSafeStack-27]
	_ = x[enum.FuncAttrSanitizeAddress-28]
	_ = x[enum.FuncAttrSanitizeHWAddress-29]
	_ = x[enum.FuncAttrSanitizeMemory-30]
	_ = x[enum.FuncAttrSanitizeMemTag-31]
	_ = x[enum.FuncAttrSanitizeThread-32]
	_ = x[enum.FuncAttrSpeculatable-33]
	_ = x[enum.FuncAttrSpeculativeLoadHardening-34]
	_ = x[enum.FuncAttrSSP-35]
	_ = x[enum.FuncAttrSSPReq-36]
	_ = x[enum.FuncAttrSSPStrong-37]
	_ = x[enum.FuncAttrStrictFP-38]
	_ = x[enum.FuncAttrUwtable-39]
	_ = x[enum.FuncAttrWillReturn-40]
	_ = x[enum.FuncAttrWriteOnly-41]
}

const _FuncAttr_name = "alwaysinlineargmemonlybuiltincoldconvergentinaccessiblemem_or_argmemonlyinaccessiblememonlyinlinehintjumptableminsizenakednobuiltinnoduplicatenofreenoimplicitfloatnoinlinenonlazybindnorecursenoredzonenoreturnnosyncnounwindoptnoneoptsizereadnonereadonlyreturns_twicesafestacksanitize_addresssanitize_hwaddresssanitize_memorysanitize_memtagsanitize_threadspeculatablespeculative_load_hardeningsspsspreqsspstrongstrictfpuwtablewillreturnwriteonly"

var _FuncAttr_index = [...]uint16{0, 12, 22, 29, 33, 43, 72, 91, 101, 110, 117, 122, 131, 142, 148, 163, 171, 182, 191, 200, 208, 214, 222, 229, 236, 244, 252, 265, 274, 290, 308, 323, 338, 353, 365, 391, 394, 400, 409, 417, 424, 434, 443}

// FuncAttrFromString returns the FuncAttr enum corresponding to s.
func FuncAttrFromString(s string) enum.FuncAttr {
	if len(s) == 0 {
		return 0
	}
	for i := range _FuncAttr_index[:len(_FuncAttr_index)-1] {
		if s == _FuncAttr_name[_FuncAttr_index[i]:_FuncAttr_index[i+1]] {
			return enum.FuncAttr(i)
		}
	}
	panic(fmt.Errorf("unable to locate FuncAttr enum corresponding to %q", s))
}
