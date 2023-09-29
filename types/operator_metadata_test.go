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
				Name:              "test",
				Description:       "test",
				Logo:              "https://test.com/test.png",
				TwitterProfileUrl: "https://twitter.com/test",
				Website:           "https://test.com",
			},
			wantErr: false,
		},
		{
			name: "Invalid metadata - no name",
			metadata: OperatorMetadata{
				Name:              "",
				Description:       "test",
				Logo:              "https://test.com/test.png",
				TwitterProfileUrl: "https://twitter.com/test",
				Website:           "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - no description",
			metadata: OperatorMetadata{
				Name:              "test",
				Description:       "",
				Logo:              "https://test.com/test.png",
				TwitterProfileUrl: "https://twitter.com/test",
				Website:           "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - no logo",
			metadata: OperatorMetadata{
				Name:              "test",
				Description:       "test",
				Logo:              "",
				TwitterProfileUrl: "https://twitter.com/test",
				Website:           "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid logo extension",
			metadata: OperatorMetadata{
				Name:              "test",
				Description:       "test",
				Logo:              "https://test.com/test.exe",
				TwitterProfileUrl: "https://twitter.com/test",
				Website:           "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid logo no extension",
			metadata: OperatorMetadata{
				Name:              "test",
				Description:       "test",
				Logo:              "https://test.com/test",
				TwitterProfileUrl: "https://twitter.com/test",
				Website:           "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid website url #1",
			metadata: OperatorMetadata{
				Name:              "test",
				Description:       "test",
				Logo:              "https://test.com/test.png",
				TwitterProfileUrl: "https://twitter.com/test",
				Website:           "https",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid website url #2",
			metadata: OperatorMetadata{
				Name:              "test",
				Description:       "test",
				Logo:              "https://test.com/test.png",
				TwitterProfileUrl: "https://twitter.com/test",
				Website:           "https:/test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid website url #3",
			metadata: OperatorMetadata{
				Name:              "test",
				Description:       "test",
				Logo:              "https://test.com/test.png",
				TwitterProfileUrl: "https://twitter.com/test",
				Website:           "ps://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid twitter url #1",
			metadata: OperatorMetadata{
				Name:              "test",
				Description:       "test",
				Logo:              "https://test.com/test.png",
				TwitterProfileUrl: "http",
				Website:           "https://test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid twitter url #2",
			metadata: OperatorMetadata{
				Name:              "test",
				Description:       "test",
				Logo:              "https://test.com/test.png",
				TwitterProfileUrl: "ht://twitter.com/test",
				Website:           "https:/test.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid metadata - invalid twitter url #3",
			metadata: OperatorMetadata{
				Name:              "test",
				Description:       "test",
				Logo:              "https://test.com/test.png",
				TwitterProfileUrl: "https://twitt",
				Website:           "ps://test.com",
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
