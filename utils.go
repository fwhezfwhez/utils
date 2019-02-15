package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Execute a function n times, the function as the handler should be well format to func() (bool,error).
// Bool stands for its ok or not ,error stands for its error
// Example:
//		f := func() (bool, error) {
//		fmt.Println("connect to mysql")
//		return false, errors.New("time out")
//		}
//		if e := RetryHandler(3, f); e != nil {
//			fmt.Println(e.Error())
//		}
//
func RetryHandler(n int, f func() (bool, error)) error {
	ok, er := f()
	if ok && er == nil {
		return nil
	} else {
		if n-1 > 0 {
			return RetryHandler(n-1, f)
		} else {
			return er
		}
	}
}

// To well print a struct type value
// Example:
//		type User struct{
//			Name string
//			Age int
//			Sal float64
//			Friends []User
//		}
//		user := User{
//		Name:"ft",
//		Age: 9,
//		Sal: 1000,
//		Friends: []User{User{"f1" , 11, 0.,nil},{"f2" , 12, 0.,nil}},
//		}
//		SmartPrint(user)
func SmartPrint(i interface{}) {
	var kv = make(map[string]interface{})
	vValue := reflect.ValueOf(i)
	vType := reflect.TypeOf(i)
	for i := 0; i < vValue.NumField(); i++ {
		kv[vType.Field(i).Name] = vValue.Field(i)
	}
	fmt.Println("receive:")
	for k, v := range kv {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println()
	}
}

// generate limit offset by page size , page index , total count
func ToLimitOffset(sizeIn string, indexIn string, count int) (limit int, offset int) {
	size, _ := strconv.Atoi(sizeIn)
	index, _ := strconv.Atoi(indexIn)
	//1
	if count == 0 {
		return size, 0
	}
	var pageMax int
	//1%10
	if count%size == 0 {
		pageMax = count / size
	} else {
		//1
		pageMax = count/size + 1
	}
	//1<=1
	if pageMax <= index {
		index = pageMax
	}
	offset = size * (index - 1)

	if offset == -10 {
		offset = 0
	}
	return size, offset
}

// MD5 encrypt
func MD5(rawMsg string) string {
	data := []byte(rawMsg)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return strings.ToUpper(md5str1)
}

// Change arg to string
// if arg is a ptr kind, then change what it points to  to string
func ToString(arg interface{}) string {
	tmp := reflect.Indirect(reflect.ValueOf(arg)).Interface()
	switch v := tmp.(type) {
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case string:
		return v
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case time.Time:
		return v.Format("2006-01-02 15:04:05")
	case fmt.Stringer:
		return v.String()
	default:
		return ""
	}
}

// To judge a value whether zero or not.
// By the way, '%' '%%' is regarded as zero.
func IfZero(arg interface{}) bool {
	if arg == nil {
		return true
	}
	switch v := arg.(type) {
	case int, int32, int16, int64:
		if v == 0 {
			return true
		}
	case float32:
		r := float64(v)
		return math.Abs(r-0) < 0.0000001
	case float64:
		return math.Abs(v-0) < 0.0000001
	case string:
		if v == "" || v == "%%" || v == "%" {
			return true
		}
	case *string, *int, *int64, *int32, *int16, *int8, *float32, *float64, *time.Time:
		if v == nil {
			return true
		}
	case time.Time:
		return v.IsZero()
	case decimal.Decimal:
		tmp,_ := v.Float64()
		return math.Abs(tmp-0) < 0.0000001
	default:
		return false
	}
	return false
}

// Convert a struct type value into a format string as  a=x&&b=y&&c=z sorted by key asc like a=x&b=y&c=z not b=y,a=x,c=z
func ToParam(vx interface{}, tag string) string {
	var result string
	var tagName_FieldValue = make(map[string]string)
	var SortedArr = make([]string, 0)
	var tagValueTemp string
	var valueStrTemp interface{}

	vType := reflect.TypeOf(vx)
	vValue := reflect.ValueOf(vx)
	for i := 0; i < vType.NumField(); i++ {
		tagValueTemp = FiltTag(vType.Field(i).Tag.Get(tag))
		// set filter field
		if tagValueTemp == "" || tagValueTemp == tag || tagValueTemp == "sign" {
			continue
		}
		valueStrTemp = vValue.Field(i).Interface()
		if IfZero(valueStrTemp) {
			continue
		}
		tagName_FieldValue[tagValueTemp] = ToString(valueStrTemp)
		SortedArr = append(SortedArr, tagValueTemp)
	}
	sort.Strings(SortedArr)

	for i, v := range SortedArr {
		if i == 0 {
			result = result + v + "=" + tagName_FieldValue[v]
			continue
		}
		result = result + "&" + v + "=" + tagName_FieldValue[v]
	}
	return result
}

// Get the tag's first value seperated by ','
func FiltTag(tag string) string {
	if strings.Contains(tag, ",") {
		return strings.Split(tag, ",")[0]
	} else {
		return tag
	}
}

// Transer an obj to a map by json
func Obj2MapByJson(obj interface{}) (map[string]interface{}, error) {
	var rs = make(map[string]interface{}, 0)
	buf, er := json.Marshal(obj)
	if er != nil {
		return rs, er
	}
	er = json.Unmarshal(buf, &rs)
	if er != nil {
		return rs, er
	}
	return rs, nil
}

// Transer an obj to a map by reflect
func Obj2MapByReflect(obj interface{}) (map[string]interface{}) {
	var rs = make(map[string]interface{}, 0)
	vValue := reflect.ValueOf(obj)
	vType := reflect.TypeOf(obj)
	var tag = ""
	for i := 0; i < vValue.NumField(); i++ {
		tag = vType.Field(i).Tag.Get("field")
		if tag == "-" {
			continue
		}
		rs[strings.ToLower(vType.Field(i).Name)] = vValue.Field(i).Interface()
	}
	return rs
}

// AddJSONFormTag add json and form tag for a golang struct without tag and annotation
func AddJSONFormTag(in string) string {
	var result string
	scanner := bufio.NewScanner(strings.NewReader(in))
	var oldLineTmp = ""
	var lineTmp = ""
	var propertyTmp = ""
	var seperateArr []string
	for scanner.Scan() {
		oldLineTmp = scanner.Text()
		lineTmp = strings.Trim(scanner.Text(), " ")
		if strings.Contains(lineTmp, "{") || strings.Contains(lineTmp, "}") {
			result = result + oldLineTmp + "\n"
			continue
		}
		seperateArr = Split(lineTmp, " ")
		// 接口或者父类声明不参与tag, 自带tag不参与tag
		if len(seperateArr) == 1 || len(seperateArr) == 3 {
			continue
		}
		propertyTmp = HumpToUnderLine(seperateArr[0])
		oldLineTmp = oldLineTmp + fmt.Sprintf("    `json:\"%s\" form:\"%s\"`", propertyTmp, propertyTmp)
		result = result + oldLineTmp + "\n"
	}
	return result
}

// Split 增强型Split，对  a,,,,,,,b,,c     以","进行切割成[a,b,c]
func Split(s string, sub string) []string {
	var rs = make([]string, 0, 20)
	tmp := ""
	Split2(s, sub, &tmp, &rs)
	return rs
}

// Split2 附属于Split，可独立使用
func Split2(s string, sub string, tmp *string, rs *[]string) {
	s = strings.Trim(s, sub)
	if !strings.Contains(s, sub) {
		*tmp = s
		*rs = append(*rs, *tmp)
		return
	}
	for i := range s {
		if string(s[i]) == sub {
			*tmp = s[:i]
			*rs = append(*rs, *tmp)
			s = s[i+1:]
			Split2(s, sub, tmp, rs)
			return
		}
	}
}

// HumpToUnderLine 驼峰转下划线
func HumpToUnderLine(s string) string {
	if s == "ID" {
		return "id"
	}
	var rs string
	elements := FindUpperElement(s)
	for _, e := range elements {
		s = strings.Replace(s, e, "_"+strings.ToLower(e), -1)
	}
	rs = strings.Trim(s, " ")
	rs = strings.Trim(rs, "\t")
	return strings.Trim(rs, "_")
}

// FindUpperElement 找到字符串中大写字母的列表,附属于HumpToUnderLine
func FindUpperElement(s string) []string {
	var rs = make([]string, 0, 10)
	for i := range s {
		if s[i] >= 65 && s[i] <= 90 {
			rs = append(rs, string(s[i]))
		}
	}
	return rs
}

// postgres数据库表转go model
// 使用时注入驱动 	_ "github.com/jinzhu/gorm/dialects/postgres"
func TableToStruct(dataSource string, tableName string) string {

	columnString := ""
	tmp := ""
	columns := FindColumns(dataSource, tableName)
	for _, column := range columns {

		tmp = fmt.Sprintf("    %s  %s\n", UnderLineToHump(column.ColumnName), typeConvert(column.ColumnType))
		columnString = columnString + tmp
	}

	rs := fmt.Sprintf("type %s struct{\n%s}", UnderLineToHump(HumpToUnderLine(tableName)), columnString)
	return rs
}

// 类型转换pg->go
func typeConvert(s string) string {
	if strings.Contains(s, "char") || In(s, []string{
		"text",
	}) {
		return "string"
	}
	if In(s, []string{"bigint", "bigserial", "integer", "smallint", "serial", "big serial"}) {
		return "int"
	}
	if In(s, []string{"numeric", "decimal", "real"}) {
		return "decimal.Decimal"
	}
	if In(s, []string{"bytea"}) {
		return "[]byte"
	}
	if strings.Contains(s, "time") || In(s, []string{"date"}) {
		return "time.Time"
	}
	if In(s, []string{"bigint", "bigserial", ""}) {
		return "json.RawMessage"
	}
	return "interface{}"
}

// s 是否in arr
func In(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

// UnderLineToHump 下划线转驼峰
func UnderLineToHump(s string) string {
	arr := strings.Split(s, "_")
	for i, v := range arr {
		arr[i] = strings.ToUpper(string(v[0])) + string(v[1:])
	}
	return strings.Join(arr, "")
}

// 数据库列属性
type Column struct {
	ColumnNumber int    `gorm:"column_number"` // column index
	ColumnName   string `gorm:"column_name"`   // column_name
	ColumnType   string `gorm:"column_type"`   // column_type
}

// 根据数据源，表明获取列属性
// 使用时注入驱动 	_ "github.com/jinzhu/gorm/dialects/postgres"
func FindColumns(dataSource string, tableName string) []Column {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(fmt.Sprintf("recover from a fatal error : %v", e))
		}
	}()
	var FindColumnsSql = `
        SELECT
            a.attnum AS column_number,
            a.attname AS column_name,
            --format_type(a.atttypid, a.atttypmod) AS column_type,
            a.attnotnull AS not_null,
			COALESCE(pg_get_expr(ad.adbin, ad.adrelid), '') AS default_value,
    		COALESCE(ct.contype = 'p', false) AS  is_primary_key,
    		CASE
        	WHEN a.atttypid = ANY ('{int,int8,int2}'::regtype[])
          		AND EXISTS (
				SELECT 1 FROM pg_attrdef ad
             	WHERE  ad.adrelid = a.attrelid
             	AND    ad.adnum   = a.attnum
             	AND    ad.adsrc = 'nextval('''
                	|| (pg_get_serial_sequence (a.attrelid::regclass::text
                	                          , a.attname))::regclass
                	|| '''::regclass)'
             	)
            THEN CASE a.atttypid
                    WHEN 'int'::regtype  THEN 'serial'
                    WHEN 'int8'::regtype THEN 'bigserial'
                    WHEN 'int2'::regtype THEN 'smallserial'
                 END
			WHEN a.atttypid = ANY ('{uuid}'::regtype[]) AND COALESCE(pg_get_expr(ad.adbin, ad.adrelid), '') != ''
            THEN 'autogenuuid'
        	ELSE format_type(a.atttypid, a.atttypmod)
    		END AS column_type
		FROM pg_attribute a
		JOIN ONLY pg_class c ON c.oid = a.attrelid
		JOIN ONLY pg_namespace n ON n.oid = c.relnamespace
		LEFT JOIN pg_constraint ct ON ct.conrelid = c.oid
		AND a.attnum = ANY(ct.conkey) AND ct.contype = 'p'
		LEFT JOIN pg_attrdef ad ON ad.adrelid = c.oid AND ad.adnum = a.attnum
		WHERE a.attisdropped = false
		AND n.nspname = 'public'
		AND c.relname = ?
		AND a.attnum > 0
		ORDER BY a.attnum
	`
	db, err := gorm.Open("postgres", dataSource)
	db.SingularTable(true)
	//db.LogMode(true)
	if err != nil {
		panic(err)
	}
	var columns = make([]Column, 0, 10)
	db.Raw(FindColumnsSql, tableName).Find(&columns)
	return columns
}
