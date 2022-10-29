package ethereum

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestABI(t *testing.T) {
	a, err := LoadABI("testdata/TLC.json")
	require.NoError(t, err)
	payload, err := a.PackJSON(
		"transfer",
		[]byte(`"0x2B7ff5d4C14A9Da8d5C9354c7A52aB40DdC1C01e"`),
		[]byte(`"0xffffffffffffffffffff"`),
	)
	require.NoError(t, err)
	assert.Equal(t, "0xa9059cbb0000000000000000000000002b7ff5d4c14a9da8d5c9354c7a52ab40ddc1c01e00000000000000000000000000000000000000000000ffffffffffffffffffff", payload)
}

func TestParseArg(t *testing.T) {
	testCases := []struct {
		desc           string
		solidityType   string
		arg            []byte
		expectedResult interface{}
		expectedErr    string
	}{
		{
			desc:           "uint32",
			solidityType:   "uint32",
			arg:            []byte("1234"),
			expectedResult: uint32(1234),
		},
		{
			desc:           "uint256",
			solidityType:   "uint256",
			arg:            []byte("\"0xff\""),
			expectedResult: new(big.Int).SetInt64(255),
		},
		{
			desc:           "bool",
			solidityType:   "bool",
			arg:            []byte("true"),
			expectedResult: true,
		},
		{
			desc:           "string",
			solidityType:   "string",
			arg:            []byte("\"test\""),
			expectedResult: "test",
		},
		{
			desc:           "address",
			solidityType:   "address",
			arg:            []byte("\"0x2B7ff5d4C14A9Da8d5C9354c7A52aB40DdC1C01e\""),
			expectedResult: common.HexToAddress("0x2B7ff5d4C14A9Da8d5C9354c7A52aB40DdC1C01e"),
		},
		{
			desc:           "bytes",
			solidityType:   "bytes",
			arg:            []byte("\"0x2B7ff5d4C14A9Da8d5C9354c7A52aB40DdC1C01e\""),
			expectedResult: hexutil.MustDecode("0x2B7ff5d4C14A9Da8d5C9354c7A52aB40DdC1C01e"),
		},
		{
			desc:           "int64[]",
			solidityType:   "int64[]",
			arg:            []byte("[1,2,3,4]"),
			expectedResult: []int64{1, 2, 3, 4},
		},
		{
			desc:           "int64[4]",
			solidityType:   "int64[4]",
			arg:            []byte("[1,2,3,4]"),
			expectedResult: [4]int64{1, 2, 3, 4},
		},
		{
			desc:           "int64[][]",
			solidityType:   "int64[][]",
			arg:            []byte("[[1,2],[3,4]]"),
			expectedResult: [][]int64{{1, 2}, {3, 4}},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			abiType, err := abi.NewType(tt.solidityType, "", nil)
			require.NoError(t, err)

			res, err := ParseJSONArg(tt.arg, &abiType)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedResult, res)
		})
	}
}
