// Code generated by "stringer -type Section"; DO NOT EDIT.

package influxdata

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MeasurementSection-0]
	_ = x[TagSection-1]
	_ = x[FieldSection-2]
	_ = x[TimeSection-3]
	_ = x[newlineSection-4]
	_ = x[endSection-5]
}

const _Section_name = "MeasurementSectionTagSectionFieldSectionTimeSectionnewlineSectionendSection"

var _Section_index = [...]uint8{0, 18, 28, 40, 51, 65, 75}

func (i Section) String() string {
	if i >= Section(len(_Section_index)-1) {
		return "Section(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Section_name[_Section_index[i]:_Section_index[i+1]]
}
