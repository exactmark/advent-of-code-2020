package Day13

import "testing"

func Test_findPt2TimeStamp(t *testing.T) {
	type args struct {
		idList string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"case1",args{"7,13,x,x,59,x,31,19"},1068781},
		{"case2",args{"17,x,13,19"},3417},
		{"case3",args{"67,7,59,61"}, 754018},
		{"case4",args{"67,x,7,59,61"}, 779210},
		{"case5",args{"67,7,x,59,61"}, 1261476},
		{"case6",args{"1789,37,47,1889"}, 1202161486},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPt2TimeStamp(tt.args.idList); got != tt.want {
				t.Errorf("findPt2TimeStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}