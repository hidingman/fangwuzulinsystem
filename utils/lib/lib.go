package lib

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Time struct {
	time.Time
}

// MarshalJSON 序列化为JSON
func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("\"\""), nil
	}
	stamp := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// UnmarshalJSON 反序列化为JSON
func (t *Time) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")
	if str == "" {
		// 空字符串处理
		t.Time = time.Time{}
		return nil
	}
	parsedTime, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

// String 重写String方法
func (t *Time) String() string {
	data, _ := json.Marshal(t)
	return string(data)
}

// FieldType 数据类型
func (t *Time) FieldType() int {
	return orm.TypeDateTimeField

}

// SetRaw 读取数据库值
func (t *Time) SetRaw(value interface{}) error {
	if v, ok := value.(time.Time); ok {
		t.Time = v
	}
	return nil
}

// RawValue 写入数据库
func (t *Time) RawValue() interface{} {
	str := t.Format("2006-01-02 15:04:05")
	if str == "0001-01-01 00:00:00" {
		return nil
	}
	return str
}

type Date struct {
	Time
}

// 将结构体转换为map
func StructToMap(obj interface{}) (map[string]interface{}, error) {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// 确保传入的是一个结构体
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, got %s", val.Kind())
	}

	// 创建一个map来存储结构体的字段
	data := make(map[string]interface{})

	// 遍历结构体的所有字段
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		data[typeField.Name] = valueField.Interface()
	}

	return data, nil
}

// MarshalJSON 序列化为JSON
func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte("\"\""), nil
	}
	stamp := fmt.Sprintf("\"%s\"", d.Format("2006-01-02"))
	return []byte(stamp), nil
}

// UnmarshalJSON 反序列化为JSON
func (d *Date) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")
	beego.Info("UnmarshalJSON:", str)
	if str == "" {
		// 空字符串处理
		d.Time = Time{}
		return nil
	}

	// 使用指定的日期格式进行解析
	parsedTime, err := time.Parse("2006-01-02", str)
	if err != nil {
		beego.Info("解析失败:", err)
		return err
	}
	// 保证精度为天
	d.Time = Time{parsedTime.Truncate(24 * time.Hour)}
	return nil
}

func getUserIDFromToken(tokenString string) (string, error) {
	// 解析 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法是否正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your_secret_key"), nil
	})

	if err != nil {
		return "", err
	}

	// 验证 token 是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["userID"].(string)
		return userID, nil
	}

	return "", fmt.Errorf("invalid token")
}
