package middleware

import (
	"hifi/types"
	"log/slog"
)

func extractIDs(t *types.TidalNew, moduleIndex int) []int {
	var ids []int

	for index, row := range t.Rows {
		if index != moduleIndex {
			continue
		}

		for _, module := range row.Modules {
			for _, item := range module.PagedList.Items {
				id := item.ID
				ids = append(ids, id)

				go func(id int, title, cover string) {
					slog.Info("Tidal item",
						"title", title,
						"id", id,
						"cover", cover,
					)
				}(id, item.Title, item.Cover)
			}
		}
	}

	return ids
}
