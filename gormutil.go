package util

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"strings"
	"time"
)

// 批量插入数据  values 参数必须为 数组， validColList 为想插入的字段
func GormBatchInsert(db *gorm.DB, values interface{}, validColList []string) error {

	t1 := time.Now()
	defer func() {
		elapsed := time.Since(t1)
		fmt.Println("App elapsed: ", elapsed)
	}()

	dataType := reflect.TypeOf(values)
	if dataType.Kind() != reflect.Slice {
		return errors.New("values must be a slice!")
	}

	val := reflect.ValueOf(values)
	if val.Len() <= 0 {
		return nil
	}

	scope := db.NewScope(val.Index(0).Interface())
	var realColList []string
	if len(validColList) == 0 {
		for _, field := range scope.Fields() {
			realColList = append(realColList, field.DBName)
		}
	} else {
		for _, colName := range validColList {
			realColList = append(realColList, colName)
		}
	}

	var args []string
	for i := 0; i < len(realColList); i++ {
		args = append(args, "?")
	}

	rowSQL := "(" + strings.Join(args, ", ") + ")"

	sqlStr := "INSERT INTO " + scope.TableName() + "(" + strings.Join(realColList, ",") + ") VALUES "

	var vals []interface{}

	var inserts []string

	for sliceIndex := 0; sliceIndex < val.Len(); sliceIndex++ {
		data := val.Index(sliceIndex).Interface()

		inserts = append(inserts, rowSQL)
		//vals = append(vals, elem.Prop1, elem.Prop2, elem.Prop3)
		elemScope := db.NewScope(data)
		for _, validCol := range realColList {
			field, ok := elemScope.FieldByName(validCol)
			if !ok {
				return errors.New("can not find col(" + validCol + ")")
			}

			vals = append(vals, field.Field.Interface())
		}
	}

	sqlStr = sqlStr + strings.Join(inserts, ",")

	err := db.Exec(sqlStr, vals...).Error
	if err != nil {

	}

	return err
}

// 批量插入数据  values 参数必须为 数组， validColList 为想插入的字段
func GormBatchInsertOnDuplicate(db *gorm.DB, values interface{}, validColList []string, duplicateUpdateColList []string) error {

	t1 := time.Now()
	defer func() {
		elapsed := time.Since(t1)
		fmt.Println("App elapsed: ", elapsed)
	}()

	dataType := reflect.TypeOf(values)
	if dataType.Kind() != reflect.Slice {
		return errors.New("values muset be a slice!")
	}

	val := reflect.ValueOf(values)
	if val.Len() <= 0 {
		return nil
	}

	scope := db.NewScope(val.Index(0).Interface())

	var realColList []string
	if len(validColList) == 0 {
		for _, field := range scope.Fields() {
			realColList = append(realColList, field.DBName)
		}
	} else {
		for _, colName := range validColList {
			realColList = append(realColList, colName)
		}
	}

	var args []string
	for i := 0; i < len(realColList); i++ {
		args = append(args, "?")
	}

	rowSQL := "(" + strings.Join(args, ", ") + ")"

	sqlStr := "INSERT INTO " + scope.TableName() + "(" + strings.Join(realColList, ",") + ") VALUES "

	var vals []interface{}

	var inserts []string

	for sliceIndex := 0; sliceIndex < val.Len(); sliceIndex++ {
		data := val.Index(sliceIndex).Interface()

		inserts = append(inserts, rowSQL)
		//vals = append(vals, elem.Prop1, elem.Prop2, elem.Prop3)
		elemScope := db.NewScope(data)
		for _, validCol := range realColList {
			field, ok := elemScope.FieldByName(validCol)
			if !ok {
				return errors.New("can not find col(" + validCol + ")")
			}

			vals = append(vals, field.Field.Interface())
		}
	}

	sqlStr = sqlStr + strings.Join(inserts, ", ")

	if len(duplicateUpdateColList) > 0 {
		dulicateStr := " ON DUPLICATE KEY UPDATE "
		var dulicateList []string
		for _, duplicateUpdateCol := range duplicateUpdateColList {
			dulicateList = append(dulicateList, fmt.Sprintf("%s = VALUES(%s)", duplicateUpdateCol, duplicateUpdateCol))
		}

		sqlStr += dulicateStr + strings.Join(dulicateList, ", ")
	}

	err := db.Exec(sqlStr, vals...).Error
	if err != nil {

	}

	return err
}

// 批量插入数据  values 参数必须为 数组， validColList 为想插入的字段
func GormBatchInsertOnDuplicate2(db *gorm.DB, values interface{}, validColList []string, duplicateUpdateColList []string, duplicateUpdateColMap map[string]string) error {

	t1 := time.Now()
	defer func() {
		elapsed := time.Since(t1)
		fmt.Println("App elapsed: ", elapsed)
	}()

	dataType := reflect.TypeOf(values)
	if dataType.Kind() != reflect.Slice {
		return errors.New("values muset be a slice!")
	}

	val := reflect.ValueOf(values)
	if val.Len() <= 0 {
		return nil
	}

	scope := db.NewScope(val.Index(0).Interface())

	var realColList []string
	if len(validColList) == 0 {
		for _, field := range scope.Fields() {
			realColList = append(realColList, field.DBName)
		}
	} else {
		for _, colName := range validColList {
			realColList = append(realColList, colName)
		}
	}

	var args []string
	for i := 0; i < len(realColList); i++ {
		args = append(args, "?")
	}

	rowSQL := "(" + strings.Join(args, ", ") + ")"

	sqlStr := "INSERT INTO " + scope.TableName() + "(" + strings.Join(realColList, ",") + ") VALUES "

	var vals []interface{}

	var inserts []string

	for sliceIndex := 0; sliceIndex < val.Len(); sliceIndex++ {
		data := val.Index(sliceIndex).Interface()

		inserts = append(inserts, rowSQL)
		//vals = append(vals, elem.Prop1, elem.Prop2, elem.Prop3)
		elemScope := db.NewScope(data)
		for _, validCol := range realColList {
			field, ok := elemScope.FieldByName(validCol)
			if !ok {
				return errors.New("can not find col(" + validCol + ")")
			}

			vals = append(vals, field.Field.Interface())
		}
	}

	sqlStr = sqlStr + strings.Join(inserts, ", ")

	var dulicateList []string
	dulicateStr := " ON DUPLICATE KEY UPDATE "
	if len(duplicateUpdateColList) > 0 {
		for _, duplicateUpdateCol := range duplicateUpdateColList {
			dulicateList = append(dulicateList, fmt.Sprintf("%s = VALUES(%s)", duplicateUpdateCol, duplicateUpdateCol))
		}
	}

	if len(duplicateUpdateColMap) > 0 {
		for duplicateUpdateColName, duplicateUpdateColValue := range duplicateUpdateColMap {
			dulicateList = append(dulicateList, fmt.Sprintf("%s = %s ", duplicateUpdateColName, duplicateUpdateColValue))
		}
	}

	if len(dulicateList) > 0 {
		sqlStr += dulicateStr + strings.Join(dulicateList, ", ")
	}

	err := db.Exec(sqlStr, vals...).Error
	if err != nil {

	}

	return err
}
