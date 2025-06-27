package utils

import (
	"fmt"
	"reflect"
	"strings"

	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/crypto/sha3"
)

func CapitalizeFieldName(name string) string {
	return strings.ToUpper(name[:1]) + name[1:]
}

func CopyStructAndChangeFieldName(originalStruct any, previousName string, newName string) any {
	val := reflect.ValueOf(originalStruct)
	typ := reflect.TypeOf(originalStruct)

	var newFields []reflect.StructField
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Name == previousName {
			field.Name = newName
		}
		newFields = append(newFields, field)
	}

	newStructType := reflect.StructOf(newFields)
	newStruct := reflect.New(newStructType).Elem()

	for i := 0; i < newStruct.NumField(); i++ {
		originalField2 := val.Field(i)
		newStruct.Field(i).Set(originalField2)
	}

	return newStruct.Interface()
}

func ExtractTypeFromAbi(taskManagerAbi *abi.ABI) (abi.Type, error) {
	taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "referenceTaskIndex",
			Type: "uint32",
		},
		{
			Name: "OutputValue", // Left because abi does not support purely anonymous or underscored fields
			Type: taskManagerAbi.Events["TaskResponded"].Inputs[0].Type.TupleElems[1].String(),
		},
	})
	if err != nil {
		return abi.Type{}, fmt.Errorf("error creating abi task response type: %w", err)
	}

	return taskResponseType, nil
}

func GetDefaultHashFunction(taskResponseType abi.Type) sdktypes.TaskResponseHashFunction {
	return func(taskResponse sdktypes.TaskResponseInterface) (sdktypes.TaskResponseDigest, error) {
		arguments := abi.Arguments{
			{
				Type: taskResponseType,
			},
		}

		encodeTaskResponseByte, err := arguments.Pack(taskResponse)
		if err != nil {
			return sdktypes.Bytes32{}, fmt.Errorf("error encoding task response: %w", err)
		}

		var taskResponseDigest [32]byte
		hasher := sha3.NewLegacyKeccak256()
		hasher.Write(encodeTaskResponseByte)
		copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

		return taskResponseDigest, nil
	}
}
