package middleware

import "hifi/types"

func extractItems(t *types.TidalNew, moduleIndex int) []types.ExploreItem {
	var items []types.ExploreItem

	for index, row := range t.Rows {
		if index != moduleIndex {
			continue
		}

		for _, module := range row.Modules {
			for _, item := range module.PagedList.Items {
				items = append(items, types.ExploreItem{
					ID:    item.ID,
					Title: item.Title,
					Cover: item.Cover,
				})
			}
		}
	}

	return items
}
