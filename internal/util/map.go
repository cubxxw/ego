package util

import (
	"fmt"
	"reflect"

	"github.com/spf13/cast"
)

// MergeStringMap merge two map
// Deprecated: this function will be moved to internal package, user should not use it anymore.
func MergeStringMap(dest, src map[string]any) {
	for sk, sv := range src {
		tv, ok := dest[sk]
		if !ok {
			// val不存在时，直接赋值
			dest[sk] = sv
			continue
		}

		svType := reflect.TypeOf(sv)
		tvType := reflect.TypeOf(tv)
		if svType != tvType {
			fmt.Println("continue, type is different")
			continue
		}

		switch ttv := tv.(type) {
		case map[any]any:
			tsv := sv.(map[any]any)
			ssv := ToMapStringInterface(tsv)
			stv := ToMapStringInterface(ttv)
			MergeStringMap(stv, ssv)
			dest[sk] = stv
		case map[string]any:
			MergeStringMap(ttv, sv.(map[string]any))
			dest[sk] = ttv
		default:
			dest[sk] = sv
		}
	}
}

// ToMapStringInterface cast map[any]any to map[string]any
// Deprecated: this function will be moved to internal package, user should not use it anymore.
func ToMapStringInterface(src map[any]any) map[string]any {
	tgt := map[string]any{}
	for k, v := range src {
		tgt[fmt.Sprintf("%v", k)] = v
	}
	return tgt
}

// DeepSearchInMap deep search in map
// Deprecated: this function will be moved to internal package, user should not use it anymore.
func DeepSearchInMap(m map[string]any, paths ...string) map[string]any {
	mtmp := make(map[string]any)
	for k, v := range m {
		mtmp[k] = v
	}
	for _, k := range paths {
		m2, ok := mtmp[k]
		if !ok {
			m3 := make(map[string]any)
			mtmp[k] = m3
			mtmp = m3
			continue
		}

		m3, err := cast.ToStringMapE(m2)
		if err != nil {
			m3 = make(map[string]any)
			mtmp[k] = m3
		}
		// continue search
		mtmp = m3
	}
	return mtmp
}
