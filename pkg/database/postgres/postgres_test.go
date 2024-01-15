package postgres

import "testing"

func TestNewConnection(t *testing.T) {

	type argument struct {
		dsn string
	}

	tests := []struct {
		name        string
		argument    argument
		expectError bool
	}{
		{
			name: "Fail: invalid dsn",
			argument: argument{
				dsn: "",
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.argument

			_, err := NewConnection(args.dsn)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}
