package ethereum

import (
	"fmt"
	"math/big"
	"os"
	"reflect"

	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type ABI struct {
	abi.ABI
}

func (a *ABI) PackJSON(name string, args ...[]byte) (string, error) {
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
		switch typ.Size {
		case 8:
			return unmarshalArg(arg, reflect.TypeOf(int8(0)))
		case 16:
			return unmarshalArg(arg, reflect.TypeOf(int16(0)))
		case 32:
			return unmarshalArg(arg, reflect.TypeOf(int32(0)))
		case 64:
			return unmarshalArg(arg, reflect.TypeOf(int64(0)))
		default:
			i, err := unmarshalArg(arg, reflect.TypeOf(&hexutil.Big{}))
			if err != nil {
				return nil, err
			}
			return (*big.Int)(i.(*hexutil.Big)), nil
		}
	case abi.UintTy:
		switch typ.Size {
		case 8:
			return unmarshalArg(arg, reflect.TypeOf(uint8(0)))
		case 16:
			return unmarshalArg(arg, reflect.TypeOf(uint16(0)))
		case 32:
			return unmarshalArg(arg, reflect.TypeOf(uint32(0)))
		case 64:
			return unmarshalArg(arg, reflect.TypeOf(uint64(0)))
		default:
			i, err := unmarshalArg(arg, reflect.TypeOf(&hexutil.Big{}))
			if err != nil {
				return nil, err
			}
			return (*big.Int)(i.(*hexutil.Big)), nil
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
