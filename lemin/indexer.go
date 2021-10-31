package lemin

func Indexer(ret StartEnd) {
	ind := ret.EndGr.singParent
	for ind != nil {
		ind.untouch = true
		ind = ind.singParent
	}
}
