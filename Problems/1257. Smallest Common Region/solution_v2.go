package solution

// Runtime 32 ms, faster than ?%
// Memory 8.1 MB, less than ?%

func findSmallestRegionV2(regions [][]string, region1 string, region2 string) string {
	parents := make(map[string]string)
	for _, r := range regions {
		for _, rr := range r[1:] {
			parents[rr] = r[0]
		}
	}

	pathToRoot := make(map[string]struct{})
	for region1 != "" {
		pathToRoot[region1] = struct{}{}
		region1 = parents[region1]
	}

	for {
		if _, ok := pathToRoot[region2]; ok {
			break
		}
		region2 = parents[region2]
	}
	return region2
}
