package types

import "testing"

func TestOperatorMetadata(t *testing.T) {
	var tests = []struct {
		name     string
		metadata OperatorMetadata
		wantErr  bool
	}{
		{
			name: "Valid metadata",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://test.com/test.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			wantErr: false,
		},
		{
			name: "Invalid metadata - no name",
			metadata: OperatorMetadata{
				Name:        "",
				Description: "test",
				Logo:        "https://test.com/test.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - no description",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "",
				Logo:        "https://test.com/test.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - wrong image format",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "My operator",
				Logo:        "https://test.com/test.svg",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - description > 500 characters",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
				Logo:        "https://test.com/test.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - no logo",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid logo extension",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://test.com/test.exe",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid logo no extension",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://test.com/test",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid website url #1",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://test.com/test.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid website url #2",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://test.com/test.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https:/test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid website url #3",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://test.com/test.png",
				Twitter:     "https://twitter.com/test",
				Website:     "ps://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid twitter url #1",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://test.com/test.png",
				Twitter:     "http",
				Website:     "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid twitter url #2",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://test.com/test.png",
				Twitter:     "ht://twitter.com/test",
				Website:     "https:/test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid twitter url #3",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://test.com/test.png",
				Twitter:     "https://twitt",
				Website:     "ps://test.com",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.metadata.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
