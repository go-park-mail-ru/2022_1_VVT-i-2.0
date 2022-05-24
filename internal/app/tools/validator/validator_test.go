package validator

import "testing"

func TestIsSlug(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSlug(tt.args.str); got != tt.want {
				t.Errorf("IsSlug() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUserId(t *testing.T) {
	type args struct {
		num int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUserId(tt.args.num); got != tt.want {
				t.Errorf("IsUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}
