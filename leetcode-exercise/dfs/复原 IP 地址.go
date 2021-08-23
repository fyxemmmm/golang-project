package dfs

import "strconv"

func restoreIpAddresses(s string) []string {
	res := []string{}
	const SEG_COUNT = 4
	segments := make([]int, SEG_COUNT)

	var dfs func(s string, segID, segStart int)

	dfs = func(s string, segID, segStart int) {
		if segID == SEG_COUNT {
			if len(s) == segStart {
				ipAddr := ""
				for idx, num := range segments {
					ipAddr = ipAddr + strconv.Itoa(num)
					if idx != SEG_COUNT-1 {
						ipAddr += "."
					}
				}
				res = append(res, ipAddr)
			}
			return
		}

		if len(s) == segStart {
			return
		}

		if s[segStart] == '0' {
			segments[segID] = 0
			dfs(s, segID+1, segStart+1)
			return
		}

		num := 0
		for segEnd := segStart; segEnd < len(s); segEnd++ {
			num = num*10 + int(s[segEnd]-'0')
			if num >= 0 && num <= 0xFF {
				segments[segID] = num
				dfs(s, segID+1, segEnd+1)
			}
		}
	}

	dfs(s, 0, 0)
	return res
}
