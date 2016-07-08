package hehe
/* 快速排序， 从小到大 */

func _sort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p - left > 1 {
		_sort(values, left, p - 1)
	}
	if right - p > 1 {
		_sort(values, p + 1, right)
	}
}

func Asort(values []int) {
	_sort(values, 0, len(values) - 1)
}
