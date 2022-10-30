package ethereum

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"reflect"
	"time"

	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type ABI struct {
	abi.ABI
}

func (a *ABI) PackJSON(name string, args ...[]byte) (string, error) {
	if method, ok := a.Methods[name]; !ok {
		return "", fmt.Errorf("method '%s' not found", name)
	} else if len(args) != len(method.Inputs) {
		return "", fmt.Errorf("argument count mismatch: got %d for %d", len(args), len(method.Inputs))
	}

	var iArgs []interface{}
	for i, typ := range a.Methods[name].Inputs { //nolint: gocritic
		iArg, err := ParseJSONArg(args[i], &typ.Type)
		if err != nil {
			return "", fmt.Errorf("arg %q at position %v invalid: %v", typ.Name, i, err)
		}
		iArgs = append(iArgs, iArg)
	}

	payload, err := a.ABI.Pack(name, iArgs...)
	if err != nil {
		return "", err
	}

	return hexutil.Encode(payload), nil
}

func LoadABI(file string) (ABI, error) {
	f, err := os.Open(file)
	if err != nil {
		return ABI{}, err
	}

	a, err := abi.JSON(f)
	if err != nil {
		return ABI{}, err
	}

	return ABI{a}, nil
}

func ParseJSONArg(arg []byte, typ *abi.Type) (interface{}, error) {
	switch typ.T {
	case abi.IntTy:
		bii, err := unmarshalArg(arg, reflect.TypeOf(&Big{}))
		if err != nil {
			return nil, err
		}

		bi := (*big.Int)(bii.(*Big))

		switch typ.Size {
		case 8:
			if !bi.IsInt64() || bi.Int64() > math.MaxInt8 {
				return nil, fmt.Errorf("number exceed max value for int8")
			}
			return int8(bi.Int64()), nil
		case 16:
			if !bi.IsInt64() || bi.Int64() > math.MaxInt16 {
				return nil, fmt.Errorf("number exceed max value for int16")
			}
			return int16(bi.Int64()), nil
		case 32:
			if !bi.IsInt64() || bi.Int64() > math.MaxInt32 {
				return nil, fmt.Errorf("number exceed max value for int32")
			}
			return int32(bi.Int64()), nil
		case 64:
			if !bi.IsInt64() {
				return nil, fmt.Errorf("number exceed max value for int64")
			}
			return bi.Int64(), nil
		default:
			return bi, nil
		}
	case abi.UintTy:
		bii, err := unmarshalArg(arg, reflect.TypeOf(&Big{}))
		if err != nil {
			return nil, err
		}

		bi := (*big.Int)(bii.(*Big))

		switch typ.Size {
		case 8:
			if !bi.IsUint64() || bi.Uint64() > math.MaxUint8 {
				return nil, fmt.Errorf("number exceed max value for uint8")
			}
			return uint8(bi.Uint64()), nil
		case 16:
			if !bi.IsUint64() || bi.Uint64() > math.MaxUint16 {
				return nil, fmt.Errorf("number exceed max value for uint16")
			}
			return uint16(bi.Uint64()), nil
		case 32:
			if !bi.IsUint64() || bi.Uint64() > math.MaxUint32 {
				return nil, fmt.Errorf("number exceed max value for uint32")
			}
			return uint32(bi.Uint64()), nil
		case 64:
			if !bi.IsUint64() {
				return nil, fmt.Errorf("number exceed max value for uint64")
			}
			return bi.Uint64(), nil
		default:
			return bi, nil
		}
	case abi.BoolTy:
		return unmarshalArg(arg, reflect.TypeOf(false))
	case abi.AddressTy:
		return unmarshalArg(arg, reflect.TypeOf(common.Address{}))
	case abi.FixedBytesTy:
		b, err := unmarshalArg(arg, reflect.TypeOf(&hexutil.Bytes{}))
		if err != nil {
			return nil, err
		}

		v := reflect.New(reflect.ArrayOf(typ.Size, reflect.TypeOf(byte(0))))
		reflect.Copy(
			v,
			reflect.ValueOf([]byte(*(b.(*hexutil.Bytes)))).Slice(0, typ.Size),
		)
		return v.Interface(), nil
	case abi.BytesTy:
		b, err := unmarshalArg(arg, reflect.TypeOf(&hexutil.Bytes{}))
		if err != nil {
			return nil, err
		}

		return []byte(*(b.(*hexutil.Bytes))), nil
	case abi.StringTy:
		return unmarshalArg(arg, reflect.TypeOf(""))
	case abi.SliceTy:
		var raws []json.RawMessage
		err := json.Unmarshal(arg, &raws)
		if err != nil {
			return nil, err
		}

		results := reflect.New(reflect.SliceOf(typ.Elem.GetType()))
		for _, raw := range raws {
			res, err := ParseJSONArg(raw, typ.Elem)
			if err != nil {
				return nil, err
			}
			results.Elem().Set(reflect.Append(results.Elem(), reflect.ValueOf(res)))
		}
		return results.Elem().Interface(), nil
	case abi.ArrayTy:
		var raws []json.RawMessage
		err := json.Unmarshal(arg, &raws)
		if err != nil {
			return nil, err
		}

		results := reflect.New(reflect.ArrayOf(typ.Size, typ.Elem.GetType()))
		for i := range raws {
			res, err := ParseJSONArg(raws[i], typ.Elem)
			if err != nil {
				return nil, err
			}
			results.Elem().Index(i).Set(reflect.ValueOf(res))
		}
		return results.Elem().Interface(), nil
	default:
		return nil, fmt.Errorf("invalid type")
	}
}

func unmarshalArg(arg []byte, typ reflect.Type) (interface{}, error) {
	argV := reflect.New(typ)

	err := json.Unmarshal(arg, argV.Interface())
	if err != nil {
		return nil, err
	}

	return argV.Elem().Interface(), nil
}

type Big big.Int

// UnmarshalJSON implements json.Unmarshaler.
func (b *Big) UnmarshalJSON(input []byte) error {
	if isString(input) {
		input = input[1 : len(input)-1]
		if has0xPrefix(input) {
			//  attempt to parse as hex
			hex := new(hexutil.Big)
			err := hex.UnmarshalText(input)
			if err != nil {
				return err
			}
			*b = Big(big.Int(*hex))
			return nil
		}

		// attempt to parse number in base 10
		res, ok := new(big.Int).SetString(string(input), 10)
		if ok {
			*b = Big(*res)
			return nil
		}

		// attempt to parse duration
		d, err := time.ParseDuration(string(input))
		if err == nil {
			roundD := d.Round(time.Second)
			if d != roundD {
				return fmt.Errorf("invalid duration %s (second precision)", string(input))
			}
			*b = Big(*big.NewInt(int64(roundD) / 1e9))
			return nil
		}

		// attempt to parse time
		t, err := time.Parse(time.RFC3339, string(input))
		if err == nil {
			*b = Big(*big.NewInt(t.Unix()))
			return nil
		}

		return fmt.Errorf("invalid number string representation %q", string(input))
	}

	// attempt to parse number in base 10
	bi, ok := new(big.Int).SetString(string(input), 10)
	if ok {
		*b = Big(*bi)
		return nil
	}

	return fmt.Errorf("invalid number %q", string(input))
}

func isString(input []byte) bool {
	return len(input) >= 2 && input[0] == '"' && input[len(input)-1] == '"'
}

func has0xPrefix(input []byte) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}
