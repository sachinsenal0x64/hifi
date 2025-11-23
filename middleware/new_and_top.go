package middleware

func GetNewAndTop() []int {
	ch := make(chan []int, 2)

	go func() {
		ch <- GetNew()
	}()

	go func() {
		ch <- GetTop()
	}()

	var all []int
	for range 2 {
		ids := <-ch
		all = append(all, ids...)
	}

	return all
}
