package util

import (
	"errors"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

const tagName = "env"

func LoadEnv(obj interface{}, path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return errors.New("failed to load env vars: " + err.Error())
	}

	p := reflect.TypeOf(obj)

	if p.Kind() != reflect.Ptr {
		return errors.New("object must be a pointer")
	}
	s := p.Elem()

	switch s.Kind() {
	case reflect.Struct:
		return loadVars(obj, s)
	default:
		return errors.New("object type is not struct")
	}
}

func loadVars(obj interface{}, s reflect.Type) error {
	for i := 0; i < s.NumField(); i++ {
		var (
			field reflect.Value
			f     = s.Field(i)
		)

		if f.Type.Kind() == reflect.Struct {
			err := loadVars(reflect.ValueOf(obj).Elem().Field(i), f.Type)
			if err != nil {
				return err
			}
			continue
		}

		switch reflect.TypeOf(obj).Kind() {
		case reflect.Ptr:
			field = reflect.ValueOf(obj).Elem().Field(i)
		case reflect.Struct:
			field = obj.(reflect.Value).Field(i)
		}

		tag := f.Tag.Get(tagName)
		if tag == "" {
			return errors.New("field `" + f.Name + "` does't have tag `" + tagName + "`")
		}

		param := os.Getenv(tag)
		if param == "" {
			return errors.New("missing `" + tag + "` env var")
		}

		switch f.Type.Kind() {
		case reflect.String:
			field.SetString(param)
		case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int16, reflect.Int8:
			v, err := strconv.ParseInt(param, 10, 64)
			if err != nil {
				return err
			}
			if field.OverflowInt(v) {
				return errors.New("value of `" + tag + "` is overflowing type `" + f.Type.String() + "`")
			}
			field.SetInt(v)
		case reflect.Uint, reflect.Uint32, reflect.Uint64, reflect.Uint16, reflect.Uint8:
			v, err := strconv.ParseUint(param, 10, 64)
			if err != nil {
				return err
			}
			if field.OverflowUint(v) {
				return errors.New("value of `" + tag + "` is overflowing type `" + f.Type.String() + "`")
			}
			field.SetUint(v)
		case reflect.Bool:
			v, err := strconv.ParseBool(param)
			if err != nil {
				return err
			}
			field.SetBool(v)
		case reflect.Float64, reflect.Float32:
			v, err := strconv.ParseFloat(param, 64)
			if err != nil {
				return err
			}

			if field.OverflowFloat(v) {
				return errors.New("value of `" + tag + "` is overflowing type `" + f.Type.String() + "`")
			}
			field.SetFloat(v)
		default:
			return errors.New("`" + f.Type.String() + "` is unsupported")
		}
	}

	return nil
}
