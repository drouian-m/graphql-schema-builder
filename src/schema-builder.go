package schemabuilder

import (
	"errors"
	"fmt"
	"github.com/drouian-m/array-utils"
	"github.com/graphql-go/graphql"
	"reflect"
	"strings"
)

var typeMap = map[string]*graphql.Scalar{
	"string":    graphql.String,
	"bool":      graphql.Boolean,
	"float":     graphql.Float,
	"time.Time": graphql.DateTime,
	"float32":   graphql.Float,
	"float64":   graphql.Float,
	"int":       graphql.Int,
	"int8":      graphql.Int,
	"int16":     graphql.Int,
	"int32":     graphql.Int,
	"int64":     graphql.Int,
	"uint":      graphql.Int,
	"uint8":     graphql.Int,
	"uint16":    graphql.Int,
	"uint32":    graphql.Int,
	"uint64":    graphql.Int,
}

func getTagValue(gqlTag string, key string) (string, error) {
	elems := strings.Split(gqlTag, ",")
	if len(elems) == 0 {
		return "", errors.New("getTagValue - Field doesn't have tags")
	}
	tag := array.NewArray(elems).Find(func(e string) bool {
		var vals = strings.Split(e, "=")
		typeKey := array.NewArray(vals).FindIndex(func(e string) bool {
			return e == key
		})

		return typeKey != -1
	})
	if tag == "" {
		return "", errors.New("ConversionError - Tag is missing")
	}
	return strings.Split(tag, "=")[1], nil
}

func ConvertStructToGraphqlSchema(object interface{}) (*graphql.Object, error) {
	typeReflection := reflect.TypeOf(object)
	//valReflection := reflect.ValueOf(object)

	//TODO: test that the input obj is a struct
	converted, err := structConverter(typeReflection)
	if err != nil {
		return nil, err
	}

	output := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   typeReflection.Name(),
			Fields: converted,
		},
	)

	return output, nil
}

func structConverter(tObj reflect.Type) (graphql.Fields, error) {
	fields := graphql.Fields{}

	for i := 0; i < tObj.NumField(); i++ {
		tCurr := tObj.Field(i)
		fieldGqlType, err := getGraphqlType(tCurr)
		if err != nil {
			return nil, err
		}

		var fname string

		name, _ := getTagValue(tCurr.Tag.Get("gql"), "name")
		if name != "" {
			fname = name
		} else {
			fname = tCurr.Name
		}

		fields[fname] = &graphql.Field{
			Name: fname,
			Type: fieldGqlType,
		}
	}

	return fields, nil
}

func getGraphqlType(tField reflect.StructField) (graphql.Output, error) {
	fType := tField.Type
	if fType.String() == "time.Time" {
		elemType, err := convertSimpleType(fType.String())
		if err != nil {
			return nil, fmt.Errorf("GenerationError - Cause : %s does not have a valid tag", tField.Name)
		}
		return elemType, nil
	} else if fType.String() != "time.Time" && fType.Kind() == reflect.Struct {
		return structToObject(fType)
	} else if fType.String() != "time.Time" && fType.Kind() == reflect.Slice &&
		fType.Elem().Kind() == reflect.Struct {
		elemType, err := structToObject(fType.Elem())
		if err != nil {
			return nil, err
		}
		return graphql.NewList(elemType), nil
	} else if fType.Kind() == reflect.Slice {
		elemType, err := convertSimpleType(fType.Elem().String())
		if err != nil {
			return nil, fmt.Errorf("GenerationError - Cause : %s does not have a valid tag", tField.Name)
		}
		return graphql.NewList(elemType), nil
	}

	output, err := convertSimpleType(fType.String())
	if err != nil {
		return nil, fmt.Errorf("GenerationError - Cause : %s does not have a valid tag", tField.Name)
	}
	return output, nil
}

func structToObject(objectType reflect.Type) (*graphql.Object, error) {
	fields, err := structConverter(objectType)
	if err != nil {
		return nil, err
	}

	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:   objectType.Name(),
			Fields: fields,
		},
	), nil
}

func convertSimpleType(tagType string) (*graphql.Scalar, error) {
	graphqlType, ok := typeMap[tagType]

	if !ok {
		return nil, errors.New("Invalid Type")
	}

	return graphqlType, nil
}
