// Code generated by "stringer -linecomment -type FuncAttr"; DO NOT EDIT.

package enum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FuncAttrAlwaysInline-0]
	_ = x[FuncAttrArgMemOnly-1]
	_ = x[FuncAttrBuiltin-2]
	_ = x[FuncAttrCold-3]
	_ = x[FuncAttrConvergent-4]
	_ = x[FuncAttrInaccessibleMemOrArgMemOnly-5]
	_ = x[FuncAttrInaccessibleMemOnly-6]
	_ = x[FuncAttrInlineHint-7]
	_ = x[FuncAttrJumpTable-8]
	_ = x[FuncAttrMinSize-9]
	_ = x[FuncAttrNaked-10]
	_ = x[FuncAttrNoBuiltin-11]
	_ = x[FuncAttrNoDuplicate-12]
	_ = x[FuncAttrNoFree-13]
	_ = x[FuncAttrNoImplicitFloat-14]
	_ = x[FuncAttrNoInline-15]
	_ = x[FuncAttrNonLazyBind-16]
	_ = x[FuncAttrNoRecurse-17]
	_ = x[FuncAttrNoRedZone-18]
	_ = x[FuncAttrNoReturn-19]
	_ = x[FuncAttrNoSync-20]
	_ = x[FuncAttrNoUnwind-21]
	_ = x[FuncAttrOptNone-22]
	_ = x[FuncAttrOptSize-23]
	_ = x[FuncAttrReadNone-24]
	_ = x[FuncAttrReadOnly-25]
	_ = x[FuncAttrReturnsTwice-26]
	_ = x[FuncAttrSafeStack-27]
	_ = x[FuncAttrSanitizeAddress-28]
	_ = x[FuncAttrSanitizeHWAddress-29]
	_ = x[FuncAttrSanitizeMemory-30]
	_ = x[FuncAttrSanitizeMemTag-31]
	_ = x[FuncAttrSanitizeThread-32]
	_ = x[FuncAttrSpeculatable-33]
	_ = x[FuncAttrSpeculativeLoadHardening-34]
	_ = x[FuncAttrSSP-35]
	_ = x[FuncAttrSSPReq-36]
	_ = x[FuncAttrSSPStrong-37]
	_ = x[FuncAttrStrictFP-38]
	_ = x[FuncAttrUwtable-39]
	_ = x[FuncAttrWillReturn-40]
	_ = x[FuncAttrWriteOnly-41]
}

const _FuncAttr_name = "alwaysinlineargmemonlybuiltincoldconvergentinaccessiblemem_or_argmemonlyinaccessiblememonlyinlinehintjumptableminsizenakednobuiltinnoduplicatenofreenoimplicitfloatnoinlinenonlazybindnorecursenoredzonenoreturnnosyncnounwindoptnoneoptsizereadnonereadonlyreturns_twicesafestacksanitize_addresssanitize_hwaddresssanitize_memorysanitize_memtagsanitize_threadspeculatablespeculative_load_hardeningsspsspreqsspstrongstrictfpuwtablewillreturnwriteonly"

var _FuncAttr_index = [...]uint16{0, 12, 22, 29, 33, 43, 72, 91, 101, 110, 117, 122, 131, 142, 148, 163, 171, 182, 191, 200, 208, 214, 222, 229, 236, 244, 252, 265, 274, 290, 308, 323, 338, 353, 365, 391, 394, 400, 409, 417, 424, 434, 443}

func (i FuncAttr) String() string {
	if i >= FuncAttr(len(_FuncAttr_index)-1) {
		return "FuncAttr(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FuncAttr_name[_FuncAttr_index[i]:_FuncAttr_index[i+1]]
}
