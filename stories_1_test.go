package main

import (
	"testing"
)

func TestPokemon(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr error
	}{
		{
			name:    "case success",
			want:    1154,
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Stories1()
			if got != test.want {
				t.Errorf("error %v %v", got, test.want)
			}
			if test.wantErr == nil {
				return
			}
			if err.Error() != test.wantErr.Error() {
				t.Errorf("error %v %v", err, test.wantErr)
			}
		})
	}

}