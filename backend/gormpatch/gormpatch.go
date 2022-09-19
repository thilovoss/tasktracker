package gormpatch

import (
	"errors"
	"reflect"
	"strings"
	"tasktracker/database"

	"gorm.io/gorm/schema"
)

type JsonPatch struct {
	Path  string `json:"path,omitempty"`
	Op    string `json:"op,omitempty"`
	Value string `json:"value,omitempty"`
}

func ApplyPatch(model interface{}, patch *JsonPatch) error {
	path := strings.TrimPrefix(patch.Path, "/")
	modelType := reflect.TypeOf(model).Elem().Elem()
	field, hasField := getFieldByJsonName(modelType, path)
	if !hasField {
		return errors.New("Json name not found: " + path)
	}

	patchable := isFieldPatchable(field)
	if !patchable {
		return errors.New("field is not patchable: " + field.Name)
	}

	columnName, err := gormNameFromJsonName(modelType, path)
	if err != nil {
		return err
	}

	switch patch.Op {
	case "replace":
		err = database.Instance.Model(model).Update(columnName, patch.Value).Error
		if err != nil {
			return err
		}

	case "remove":
		err = database.Instance.Model(model).Update(columnName, nil).Error
		if err != nil {
			return err
		}

	default:
		return errors.New("unknown patch operation")
	}
	return nil
}

func gormNameFromJsonName(modelType reflect.Type, jsonName string) (columnName string, err error) {
	field, hasField := getFieldByJsonName(modelType, jsonName)
	if !hasField {
		return "", errors.New("Json name not found: " + jsonName)
	}

	gormTag, hasGormTag := field.Tag.Lookup("gorm")

	if !hasGormTag {
		return "", errors.New("column name is missing")
	}

	settings := schema.ParseTagSetting(gormTag, ";")

	columnName, foundColumnName := settings["COLUMN"]
	if foundColumnName {
		return columnName, nil
	} else {
		return "", errors.New("column name is missing")
	}
}

func getFieldByJsonName(t reflect.Type, jsonName string) (reflect.StructField, bool) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag, hasJsonTag := field.Tag.Lookup("json")
		if hasJsonTag {
			name, hasName := getJsonNameFromTag(jsonTag)
			if hasName && name == jsonName {
				return field, true
			}
		}
	}
	return reflect.StructField{}, false
}

func getJsonNameFromTag(tag string) (name string, ok bool) {
	if strings.HasPrefix(tag, ",") {
		return name, false
	}

	parts := strings.SplitN(tag, ",", 2)

	if len(parts) >= 1 && parts[0] != "-" {
		name = parts[0]
		ok = true
	}
	return name, ok
}

func isFieldPatchable(field reflect.StructField) bool {
	tag, hasTag := field.Tag.Lookup("gorm-patch")

	if hasTag {
		parts := strings.Split(tag, ",")

		for _, part := range parts {
			if part == "patchable" {
				return true
			}
		}
	}
	return false
}
