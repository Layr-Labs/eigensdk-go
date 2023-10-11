package types

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOperatorValidate(t *testing.T) {
	metadata := OperatorMetadata{
		Name:        "test",
		Description: "test",
		Logo:        "https://test.com/test.png",
		Twitter:     "https://twitter.com/test",
		Website:     "https://test.com",
	}

	var tests = []struct {
		name                    string
		operator                Operator
		metadataURLresponseCode int
		operatorMetadata        OperatorMetadata
		wantErr                 bool
	}{
		{
			name: "successful operator validation",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				EarningsReceiverAddress:   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "https://localhost",
			},
			metadataURLresponseCode: 200,
			operatorMetadata:        metadata,
			wantErr:                 false,
		},
		{
			name: "valid operator validation - 0x0 delegation approver address",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0x0",
				EarningsReceiverAddress:   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "https://localhost",
			},
			metadataURLresponseCode: 200,
			operatorMetadata:        metadata,
			wantErr:                 false,
		},
		{
			name: "failed operator validation - bad metadata",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				EarningsReceiverAddress:   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "https://localhost",
			},
			metadataURLresponseCode: 200,
			operatorMetadata:        OperatorMetadata{},
			wantErr:                 true,
		},
		{
			name: "failed operator validation - wrong operator address",
			operator: Operator{
				Address:                   "0xa",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				EarningsReceiverAddress:   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "https://localhost",
			},
			metadataURLresponseCode: 200,
			operatorMetadata:        OperatorMetadata{},
			wantErr:                 true,
		},
		{
			name: "failed operator validation - wrong earning recievers address address",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				EarningsReceiverAddress:   "0xasdf",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "https://localhost",
			},
			metadataURLresponseCode: 200,
			operatorMetadata:        OperatorMetadata{},
			wantErr:                 true,
		},
		{
			name: "failed operator validation - wrong DelegationApproverAddress address",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0x12",
				EarningsReceiverAddress:   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "https://localhost",
			},
			metadataURLresponseCode: 200,
			operatorMetadata:        OperatorMetadata{},
			wantErr:                 true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metadataBytes, err := json.Marshal(tt.operatorMetadata)
			if err != nil {
				t.Errorf("Error marshalling metadata: %v", err)
			}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.metadataURLresponseCode)
				_, _ = w.Write(metadataBytes)
			}))
			// setmetadata url to the test server
			tt.operator.MetadataUrl = server.URL

			defer server.Close()
			if err := tt.operator.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Operator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
