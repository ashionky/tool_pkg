/**
 * @Author pibing
 * @create 2022/1/25 11:07 AM
 */

package util

// SliceStringDifference 计算a中排除b后的差集
func SliceStringDifference(a []string, b []string) []string {
	c := make([]string, len(a))
	for _, va := range a {
		hasVa := false
		for _, vb := range b {
			if va == vb {
				hasVa = true
				break
			}
		}
		if !hasVa {
			c = append(c, va)
		}
	}
	return c
}

// SliceIntDifference 计算a中排除b后的差集
func SliceIntDifference(a []int, b []int) []int {
	c := make([]int, len(a))
	for _, va := range a {
		hasVa := false
		for _, vb := range b {
			if va == vb {
				hasVa = true
				break
			}
		}
		if !hasVa {
			c = append(c, va)
		}
	}
	return c
}

//  StringSliceEqual 比较两个slice是否完全相等
func SliceStringEqual(sliceA, sliceB []string) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}

	if (sliceA == nil) != (sliceB == nil) {
		return false
	}

	for i, v := range sliceA {
		if v != sliceB[i] {
			return false
		}
	}

	return true
}

//  SliceIntEqual 比较两个slice是否完全相等
func SliceIntEqual(sliceA, sliceB []int) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}

	if (sliceA == nil) != (sliceB == nil) {
		return false
	}

	for i, v := range sliceA {
		if v != sliceB[i] {
			return false
		}
	}

	return true
}

// SliceStringRemoveTarget 删除slice中第一个target值
func SliceStringRemoveTarget(slice *[]string, target string) {
	for index, v := range *slice {
		if v == target {
			*slice = append((*slice)[:index], (*slice)[index+1:]...)
			break
		}
	}
}

// SliceIntRemoveTarget 删除slice中第一个target值
func SliceIntRemoveTarget(slice *[]int, target int) {
	for index, v := range *slice {
		if v == target {
			*slice = append((*slice)[:index], (*slice)[index+1:]...)
			break
		}
	}
}

//  SliceStringDistinct 对string类型的slice去重
func SliceStringDistinct(slice *[]string) {
	for index, v := range *slice {
		for j := index + 1; j < len(*slice); j++ {
			if v == (*slice)[j] {
				*slice = append((*slice)[:j], (*slice)[j+1:]...)
				j--
			}
		}
	}
}

//  SliceIntDistinct 对int类型的slice去重
func SliceIntDistinct(slice *[]int) {
	for index, v := range *slice {
		for j := index + 1; j < len(*slice); j++ {
			if v == (*slice)[j] {
				*slice = append((*slice)[:j], (*slice)[j+1:]...)
				j--
			}
		}
	}
}
