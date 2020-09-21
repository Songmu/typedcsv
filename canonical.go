package typedcsv

func Canonical(records [][]Value) [][]Value {
	fieldTypes := make([]Type, len(records[0]))
	for _, rec := range records {
		if rec[0].Type() == TypeHeader {
			continue
		}
		for i, field := range rec {
			t1 := fieldTypes[i]
			t2 := field.Type()
			if t1 == 0 || t1 == TypeNull {
				fieldTypes[i] = t2
			}
			if t2 == TypeAny || !t1.compatible(t2) {
				// falback to TypeString
				fieldTypes[i] = TypeString
			}
		}
	}

	var ret = make([][]Value, len(records))
	for i, rec := range records {
		if rec[0].Type() == TypeHeader {
			ret[i] = rec
			continue
		}
		ret[i] = make([]Value, len(rec))
		for j, v := range rec {
			if fieldTypes[j] == TypeString && !v.Type().compatible(TypeString) {
				ret[i][j] = String(v.String())
			} else if v.Type() == TypeNull {
				ret[i][j] = Null() // remove rawValue
			} else {
				ret[i][j] = v
			}
		}
	}
	return ret
}
