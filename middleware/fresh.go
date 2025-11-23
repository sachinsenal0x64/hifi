package middleware

import "hifi/types"

func GetNewAndTopItems() []types.ExploreItem {

	ch := make(chan []types.ExploreItem, 2)

	go func() {
		ch <- GetNewItems()
	}()

	go func() {
		ch <- GetTopItems()
	}()

	var all []types.ExploreItem
	for range 2 {
		items := <-ch
		all = append(all, items...)
	}

	return all
}
