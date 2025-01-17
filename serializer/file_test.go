package serializer_test

import (
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	"grpc4_test/sample"
	"grpc4_test/serializer"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobuToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := sample.NewLaptop()
	err = serializer.ReadProtobufFromFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJsonFile(laptop1, jsonFile)
	require.NoError(t, err)

}
