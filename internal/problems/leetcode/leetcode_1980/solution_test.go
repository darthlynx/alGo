package leetcode1980

import "testing"

func TestFindDifferentBinaryString(t *testing.T) {
	tests := []struct {
		name  string
		nums  []string
		wants map[string]bool
	}{
		{
			name:  "test case 1",
			nums:  []string{"01", "10"},
			wants: map[string]bool{"00": true, "11": true},
		},
		{
			name:  "test case 2",
			nums:  []string{"00", "01"},
			wants: map[string]bool{"10": true, "11": true},
		},
		{
			name:  "test case 3",
			nums:  []string{"111", "011", "001"},
			wants: map[string]bool{"000": true, "010": true, "100": true, "101": true, "110": true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDifferentBinaryString(tt.nums); !tt.wants[got] {
				t.Errorf("findDifferentBinaryString() = %v, allowed = %v", got, tt.wants)
			}
		})
	}
}
