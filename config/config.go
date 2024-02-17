package config

import (
	"os"
	"reflect"
	"strconv"
)

type AppParams struct {
	S3Endpoint string `param:"S3_ENDPOINT"`
}

var Params = InitConfig()

func InitConfig() (params AppParams) {
	st := reflect.TypeOf(params)
	sv := reflect.ValueOf(&params).Elem()

	for i := 0; i < st.NumField(); i++ {
		envName, ok := st.Field(i).Tag.Lookup("param")
		if !ok {
			continue
		}

		envVal, ok := os.LookupEnv(envName)
		if !ok {
			continue
		}

		if st.Field(i).Type.Name() == "uint" {
			intVal, err := strconv.Atoi(envVal)
			if err != nil {
				continue
			}
			sv.Field(i).SetUint(uint64(intVal))
		} else {
			sv.Field(i).SetString(envVal)
		}
	}

	return params
}

func GetValue(configValue string, defaultValue int) (int, error) {
	if configValue == "" {
		return defaultValue, nil
	}

	return strconv.Atoi(configValue)
}
