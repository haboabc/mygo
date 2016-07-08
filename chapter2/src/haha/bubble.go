package haha

/* 冒泡排序, 从大到小 */
func Bubble(values []int) {
	var max int

	for i := 0; i < len(values) -1; i++ {
		max = values[i]
		for j := i + 1; j < len(values); j++ {
			if max < values[j] {
				max = values[j]
				values[j] = values[i]
				values[i] = max
			}
		}
	}
}
