package utils

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
	"time"
)

func ToTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {

		if t == reflect.TypeOf(time.Time{}) {
			switch f.Kind() {
			case reflect.String:
				return time.Parse("2006-01-02 15:04:05", data.(string))
			default:
				return data, nil
			}
		}

		return data, nil
	}
}

func Decode(input interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc()),
		Result: result,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(input); err != nil {
		return err
	}
	return err
}
