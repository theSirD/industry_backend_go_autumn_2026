package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	if len(nums) < 2 {
		return Stats{}
	}

	d := nums[1] - nums[0]
	s := Stats{Count: 1, Sum: d, Min: d, Max: d}

	for i := 2; i < len(nums); i++ {
		d = nums[i] - nums[i-1]
		s.Count++
		s.Sum += d
		if d < s.Min {
			s.Min = d
		}
		if d > s.Max {
			s.Max = d
		}
	}
	return s
}
