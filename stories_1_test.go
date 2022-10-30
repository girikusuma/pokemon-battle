package main

import (
	"errors"
	"testing"
)

// Pokemon pokemon by given poke.
//
// It returns int, and nil error when successful.
// Otherwise, empty int, and error will be returned.
func PokemonTDD(poke int) (int, error) {
	if poke == 1 {
		return 0, errors.New("error")
	}

	return poke, nil
}

// TestPokemon test pokemon by given t pointer of testing.T.
func TestPokemon(t *testing.T) {
	tests := []struct {
		name    string
		param   int
		want    int
		wantErr error
	}{
		{
			name:    "case success",
			param:   2,
			want:    2,
			wantErr: nil,
		},
		{
			name:    "case fail",
			param:   1,
			want:    0,
			wantErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := PokemonTDD(test.param)
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