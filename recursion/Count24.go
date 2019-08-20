/**
 * @time: 2019-08-20 22:24
 * @author: louis
 */
package recursion

import "math"

const EPS = 1e-6

func IsZero(x float64) bool {
	return math.Abs(x) <= EPS
}

func Count24(a []float64, n int) bool {
	if n == 1 {
		if IsZero(a[0] - 24) {
			return true
		} else {
			return false
		}
	}
	var b []float64
	b = make([]float64, n)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ { //枚举两个数的组合
			var m int = 0
			for k := 0; k < n; k++ {
				if k != i && k != j { //将剩余m-2的数放入b
					b = append(b, a[i])
					m++
				}
			}
			if b[m] = a[i] + a[j]; Count24(b, m+1) {
				return true
			}
			if b[m] = a[i] - a[j]; Count24(b, m+1) {
				return true
			}
			if b[m] = a[j] - a[i]; Count24(b, m+1) {
				return true
			}
			if b[m] = a[i] * a[j]; Count24(b, m+1) {
				return true
			}
			if IsZero(a[j]) {
				if b[m] = a[i] / a[j]; Count24(b, m+1) {
					return true
				}
			}
			if IsZero(a[i]) {
				if b[m] = a[j] / a[i]; Count24(b, m+1) {
					return true
				}
			}
		}
	}
	return false
}
