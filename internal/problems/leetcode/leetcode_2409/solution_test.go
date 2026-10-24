package leetcode2409

import "testing"

func TestCountDaysTogether(t *testing.T) {
	tests := []struct {
		name        string
		arriveAlice string
		leaveAlice  string
		arriveBob   string
		leaveBob    string
		want        int
	}{
		{
			name:        "example overlapping stays",
			arriveAlice: "08-15",
			leaveAlice:  "08-18",
			arriveBob:   "08-16",
			leaveBob:    "08-19",
			want:        3,
		},
		{
			name:        "example Bob arrives after Alice leaves",
			arriveAlice: "10-01",
			leaveAlice:  "10-31",
			arriveBob:   "11-01",
			leaveBob:    "12-31",
			want:        0,
		},
		{
			name:        "Alice arrives after Bob leaves",
			arriveAlice: "03-01",
			leaveAlice:  "03-05",
			arriveBob:   "02-27",
			leaveBob:    "02-28",
			want:        0,
		},
		{
			name:        "single day stay on shared endpoint",
			arriveAlice: "12-31",
			leaveAlice:  "12-31",
			arriveBob:   "12-01",
			leaveBob:    "12-31",
			want:        1,
		},
		{
			name:        "Bob stay contained across February boundary",
			arriveAlice: "02-01",
			leaveAlice:  "03-31",
			arriveBob:   "02-28",
			leaveBob:    "03-02",
			want:        3,
		},
		{
			name:        "identical full year stays",
			arriveAlice: "01-01",
			leaveAlice:  "12-31",
			arriveBob:   "01-01",
			leaveBob:    "12-31",
			want:        365,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDaysTogether(tt.arriveAlice, tt.leaveAlice, tt.arriveBob, tt.leaveBob); got != tt.want {
				t.Errorf("countDaysTogether() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDays(t *testing.T) {
	tests := []struct {
		name string
		d    string
		want int
	}{
		{name: "first day of year", d: "01-01", want: 1},
		{name: "last day of January", d: "01-31", want: 31},
		{name: "first day of February", d: "02-01", want: 32},
		{name: "last day of February", d: "02-28", want: 59},
		{name: "first day after February", d: "03-01", want: 60},
		{name: "last day of year", d: "12-31", want: 365},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toDays(tt.d); got != tt.want {
				t.Errorf("toDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
