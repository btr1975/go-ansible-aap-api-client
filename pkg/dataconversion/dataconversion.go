/*
Package dataconversion provides some helpful functions for converting data
*/
package dataconversion

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"net/http"
	"reflect"
	"strings"
)

// DataConverterInterface is the data conversion interface
type DataConverterInterface interface {
	StructToJSONString(structData any) (jsonString string, err error)
	JSONStringToStruct(structData any, jsonString string) (err error)
	StructToYAMLString(structData any) (yamlString string, err error)
	YAMLStringToStruct(structData any, yamlString string) (err error)
	ResponseBodyToStruct(structData any, response http.Response) (err error)
}

type DataConverter struct{}

func NewDataConverter() DataConverterInterface {
	return &DataConverter{}
}

// StructToJSONString converts a struct to a JSON string
//
//	:param structData: The struct to use to convert
func (dc *DataConverter) StructToJSONString(structData any) (jsonString string, err error) {
	var destinationString strings.Builder
	encoder := json.NewEncoder(&destinationString)
	err = encoder.Encode(&structData)

	if err != nil {
		return "", fmt.Errorf("error StructToJsonString encoder.Encode %v", err)
	}

	return destinationString.String(), nil
}

// JSONStringToStruct converts a JSON string to a struct
//
//	:param structData: The struct to use to convert
//	:param jsonString: The JSON string to convert
func (dc *DataConverter) JSONStringToStruct(structData any, jsonString string) (err error) {
	if reflect.ValueOf(structData).Kind() != reflect.Ptr {
		return errors.New("error JSONStringToStruct structData must be a pointer")
	}

	stringReader := strings.NewReader(jsonString)
	decoder := json.NewDecoder(stringReader)
	err = decoder.Decode(structData)

	if err != nil {
		return errors.New(fmt.Sprintf("error JSONStringToStruct decoder.Decode %v", err))
	}

	return nil
}

// StructToYAMLString converts a struct to a YAML string
//
//	:param structData: The struct to use to convert
func (dc *DataConverter) StructToYAMLString(structData any) (yamlString string, err error) {
	var destinationString strings.Builder
	encoder := yaml.NewEncoder(&destinationString)
	err = encoder.Encode(&structData)

	if err != nil {
		return "", errors.New(fmt.Sprintf("error StructToYAMLString encoder.Encode %v", err))
	}

	return destinationString.String(), nil

}

// YAMLStringToStruct converts a YAML string to a struct
//
//	:param structData: The struct to use to convert
//	:param yamlString: The YAML string to convert
func (dc *DataConverter) YAMLStringToStruct(structData any, yamlString string) (err error) {
	if reflect.ValueOf(structData).Kind() != reflect.Ptr {
		return errors.New("error YAMLStringToStruct structData must be a pointer")
	}

	stringReader := strings.NewReader(yamlString)
	decoder := yaml.NewDecoder(stringReader)
	err = decoder.Decode(structData)
	if err != nil {
		return errors.New(fmt.Sprintf("error YAMLStringToStruct decoder.Decode %v", err))
	}

	return nil

}

// ResponseBodyToStruct converts a http.Response to a struct
//
//	:param structData: The struct to use to convert
//	:param response: The http.Response to convert
func (dc *DataConverter) ResponseBodyToStruct(structData any, response http.Response) (err error) {
	if reflect.ValueOf(structData).Kind() != reflect.Ptr {
		return errors.New("error ResponseBodyToStruct structData must be a pointer")
	}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(structData)

	if err != nil {
		return errors.New(fmt.Sprintf("error ResponseBodyToStruct decoder.Decode %v", err))
	}

	return nil
}
