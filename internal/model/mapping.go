package model

func MapRegionsNames(regions []*Region) map[string]string {
	m := make(map[string]string, len(regions))
	for _, region := range regions {
		m[region.DisplayName] = region.ID
	}
	return m
}

func MapSummitNames(summits []*Summit) map[string]string {
	m := make(map[string]string, len(summits))
	for _, summit := range summits {
		m[summit.DisplayName] = summit.ID
	}
	return m
}
