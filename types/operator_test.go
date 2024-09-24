package types

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOperatorValidate(t *testing.T) {

	operatorMetadata := OperatorMetadata{
		Name:        "Madhur 1",
		Website:     "https://www.facebook.com",
		Description: "Madhur's first operator is best in this world",
		Logo:        "SET BELOW ONCE URL IS KNOWN",
		Twitter:     "https://twitter.com/shrimalmadhur",
	}

	operatorPngData, err := os.ReadFile("./.testdata/operator.png")
	require.NoError(t, err, "Failed to read operator.png")
	operatorPngBuffer := bytes.NewBuffer(operatorPngData)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/404metadata.json" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.URL.Path == "/operator.png" {
			w.Header().Set("Content-Type", "image/png")
			http.ServeContent(w, r, "operator.png", time.Now(), bytes.NewReader(operatorPngBuffer.Bytes()))
			return
		}
		marshalledOperatorMetadata, err := json.Marshal(operatorMetadata)
		require.NoError(t, err)
		w.Write(marshalledOperatorMetadata)
	}))
	defer ts.Close()
	fmt.Println(ts.URL)

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return nil, nil
		},
	}

	// Override the default resolver
	net.DefaultResolver = resolver

	// Create a custom dialer
	dialer := &net.Dialer{
		Resolver: resolver,
	}

	// Override the default HTTP transport
	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		host, port, err := net.SplitHostPort(addr)
		if err != nil {
			return nil, err
		}

		// Redirect specific hostnames to localhost
		switch host {
		case "mylocalserver.com":
			addr = net.JoinHostPort("127.0.0.1", port)
		}

		return dialer.DialContext(ctx, network, addr)
	}

	_, port, err := net.SplitHostPort(ts.URL[7:])
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("http://mylocalserver.com:%s", port)
	operatorMetadata.Logo = url + "/operator.png"

	defer func() {
		http.DefaultTransport = &http.Transport{}
	}()

	var tests = []struct {
		name                    string
		operator                Operator
		metadataURLresponseCode int
		operatorMetadata        OperatorMetadata
		wantErr                 bool
		expectedErr             error
	}{
		{
			name: "successful operator validation",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               url + "/metadata.json",
			},
			wantErr: false,
		},
		{
			name: "valid operator validation - ZeroAddress delegation approver address",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: ZeroAddress,
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               url + "/metadata.json",
			},
			wantErr: false,
		},
		{
			name: "failed operator validation - empty metadata url",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "",
			},
			wantErr:     true,
			expectedErr: utils.WrapError(ErrInvalidMetadataUrl, utils.ErrEmptyUrl),
		},
		{
			name: "failed operator validation - localhost metadata url disallowed",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "http://localhost:8080/metadata.json",
			},
			wantErr:     true,
			expectedErr: utils.WrapError(ErrInvalidMetadataUrl, utils.ErrUrlPointingToLocalServer),
		},
		{
			name: "failed operator validation - 127.0.0.1 metadata url disallowed",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "http://127.0.0.1:8080/metadata.json",
			},
			wantErr:     true,
			expectedErr: utils.WrapError(ErrInvalidMetadataUrl, utils.ErrUrlPointingToLocalServer),
		},
		{
			name: "failed operator validation - bad metadata",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               url + "/404metadata.json",
			},
			wantErr: true,
			expectedErr: utils.WrapError(
				ErrReadingMetadataUrlResponse,
				errors.New("error fetching url: 404 Not Found"),
			),
		},
		{
			name: "failed operator validation - wrong operator address",
			operator: Operator{
				Address:                   "0xa",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               url + "/metadata.json",
			},
			wantErr:     true,
			expectedErr: ErrInvalidOperatorAddress,
		},
		{
			name: "failed operator validation - wrong DelegationApproverAddress address",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0x12",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               url + "/metadata.json",
			},
			wantErr:     true,
			expectedErr: ErrInvalidDelegationApproverAddress,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedErr, tt.operator.Validate(), "error should be equal")
		})
	}
}
