package slice

func contains(s []interface{}, e interface{}) bool {
	for _, i := range s {
		if i == e {
			return true
		}
	}

	return false
}
