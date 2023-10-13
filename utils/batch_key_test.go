package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadBatchKeys(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting current directory: %s", err)
	}

	keyFolder := "test_data/generated"
	expectedBatchKeys := []BatchKey{
		{
			FilePath:   fmt.Sprintf("%s/%s/%s/%s", currentDir, keyFolder, DefaultKeyFolder, "1.bls.key.json"),
			PrivateKey: "4357240271787883726453014006887765149509214591199465958688860472447752394796",
			Password:   ">a97b_u4[@h?VJ:z[^MK",
		},
		{
			FilePath:   fmt.Sprintf("%s/%s/%s/%s", currentDir, keyFolder, DefaultKeyFolder, "2.bls.key.json"),
			PrivateKey: "91585861412425212507686922436908035103777159556932541958142896602931092396",
			Password:   ">1-%2_X*z\\-_;W%=(WU8",
		},
	}
	readKeys, err := ReadBatchKeys(keyFolder, false)
	if err != nil {
		t.Errorf("Error reading batch keys: %s", err)
	}
	assert.Equal(t, expectedBatchKeys, readKeys)
}
