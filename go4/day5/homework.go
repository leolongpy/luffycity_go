package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//解析ini文件
type server struct {
	IP   string `ini:"ip"`
	Port int64  `ini:"port"`
}
type mysql struct {
	Host     string `ini:"host"`
	Port     int64  `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}
type config struct {
	Server server `ini:"server"`
	MySql  mysql  `ini:"mysql"`
}

// Load 加载配置文件
func Load(filename string, result interface{}) (err error) {
	//判断 result是否是结构体指针
	ptrInfo := reflect.TypeOf(result)
	structInfo := ptrInfo.Elem()
	if ptrInfo.Kind() != reflect.Ptr {
		err = errors.New("please pass into a struct ptr")
		return
	}
	if structInfo.Kind() != reflect.Struct {
		err = errors.New("please pass into a struct ptr")
		return
	}

	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		err = errors.New("文件读取失败")
	}
	//解析文件
	//lineSep := "\n"
	//if runtime.GOOS == "windows" {
	//	lineSep = "\r\n"
	//}
	lineSlice := strings.Split(string(fileData), "\n")
	fmt.Println(lineSlice)
	structFieldName := ""
	for index, line := range lineSlice {
		line = strings.TrimSpace(line)

		lineNo := index + 1
		if len(line) == 0 || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		if strings.HasPrefix(line, "[") {
			structFieldName, err = parseSection(line, lineNo, structInfo)
			if err != nil {
				return err
			}
		} else {
			err = parseItem(line, lineNo, structFieldName, result)
			if err != nil {
				return err
			}
		}
	}

	return
}

//解析section
func parseSection(line string, lineNo int, structInfo reflect.Type) (structFieldName string, err error) {
	if len(line) <= 2 || line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error,invalid section:%s line:%d", line, lineNo)
		return
	}
	sectionName := strings.TrimSpace(line[1 : len(line)-1])
	if len(sectionName) == 0 {
		err = fmt.Errorf("syntax error,invalid section:%s line:%d", line, lineNo)
		return
	}
	for i := 0; i < structInfo.NumField(); i++ {
		field := structInfo.Field(i)
		fieldTag := field.Tag.Get("ini")
		if sectionName == fieldTag {
			structFieldName = field.Name
			return
		}
	}
	err = fmt.Errorf("can't find %s form struct", sectionName)
	return
}

// 解析item
func parseItem(line string, lineNo int, structFieldName string, result interface{}) (err error) {
	index := strings.Index(line, "=")
	if index == -1 {
		err = fmt.Errorf("syntax error,line:%d", lineNo)
		return
	}
	key := strings.TrimSpace(line[:index])
	value := strings.TrimSpace(line[index+1:])
	if len(key) == 0 {
		err = fmt.Errorf("syntax error,line:%d", lineNo)
		return
	}
	//结构体
	resultValue := reflect.ValueOf(result)                          //结果值
	sectionValue := resultValue.Elem().FieldByName(structFieldName) //结果值的字段
	sectionType := sectionValue.Type()                              //结果值字段的类型

	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("%s field must be a struct", structFieldName)
		return
	}
	// 遍历内嵌结构体 根据key给result赋值
	keyName := ""
	for i := 0; i < sectionType.NumField(); i++ {
		keyFileld := sectionType.Field(i)
		keyTag := keyFileld.Tag.Get("ini")
		if keyTag == key {
			keyName = keyFileld.Name
			break
		}
	}
	if len(keyName) == 0 {
		return
	}
	fieldValue := sectionValue.FieldByName(keyName)
	switch fieldValue.Type().Kind() {
	case reflect.String:
		fieldValue.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, parseErr := strconv.ParseInt(value, 10, 64)
		if parseErr != nil {
			return parseErr
		}
		fieldValue.SetInt(intVal)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		intVal, parseErr := strconv.ParseInt(value, 10, 64)
		if parseErr != nil {
			return parseErr
		}
		fieldValue.SetInt(intVal)
	case reflect.Float32, reflect.Float64:
		floatVal, parseErr := strconv.ParseFloat(value, 64)
		if parseErr != nil {
			return parseErr
		}
		fieldValue.SetFloat(floatVal)
	case reflect.Bool:
		boolval, parseErr := strconv.ParseBool(value)
		if parseErr != nil {
			return parseErr
		}
		fieldValue.SetBool(boolval)
	default:
		err = fmt.Errorf("unsupport type:%v", fieldValue.Type().Kind())
		return
	}

	return
}
func main() {
	conf := &config{}
	err := Load("./xx.ini", conf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", conf)
}
