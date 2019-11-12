package solution

// Runtime 0 ms, faster than 100%
// Memory 2.6 MB, less than 60%

func addStrings(num1 string, num2 string) string {
	b1 := []byte(num1)
	b2 := []byte(num2)
	tmp := make([]byte, 0, len(b1)+len(b2))

	i, j := len(b1)-1, len(b2)-1
	carry := 0
	for i >= 0 || j >= 0 || carry == 1 {
		var d1, d2 int
		if i >= 0 {
			d1 = int(b1[i] - '0')
		}
		if j >= 0 {
			d2 = int(b2[j] - '0')
		}

		tmp = append(tmp, byte((d1+d2+carry)%10)+'0')
		carry = (d1 + d2 + carry) / 10
		i--
		j--
	}

	result := make([]byte, len(tmp))
	for i := range tmp {
		result[len(tmp)-i-1] = tmp[i]
	}

	return string(result)
}
