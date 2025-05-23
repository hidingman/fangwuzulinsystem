package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-schema/models"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	KEY                    string = "test"
	DEFAULT_EXPIRE_SECONDS int    = 3600
	Date                          = "2006-01-02"
	DateTime                      = "2006-01-02 15:04:05"
	API_KEY                       = "7Otjpv2kn0ljQk45qXOXh5MO"
	SECRET_KEY                    = ""
	ApiKey                        = ""
	BaseURL                       = "https://api.coze.cn"
)

// CozeClient 表示与 Coze 智能体通信的客户端
type CozeClient struct {
	baseURL string
	apiKey  string
}

// NewCozeClient 创建一个新的 Coze 客户端实例
func NewCozeClient(baseURL, apiKey string) *CozeClient {
	return &CozeClient{
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

type MyCustomClaims struct {
	User models.Session
	jwt.StandardClaims
}

func GenerateToken(info *models.Session, expiredSeconds int) (tokenString string, err error) {
	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRE_SECONDS
	}
	// Create the Claims
	mySigningKey := []byte(KEY)
	expireAt := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	fmt.Println("token will be expired at ", time.Unix(expireAt, 0))
	// pass parameter to this func or not
	user := *info
	claims := MyCustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    user.Username,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("generate json web token failed !! error :", err)
	} else {
		tokenString = tokenStr
	}
	return
}

func ValidateToken(tokenString string) (info models.Session, err error) {
	if tokenString == "" {
		return info, errors.New("token is empty")
	}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(KEY), nil
		})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		info = claims.User
	} else {
		fmt.Println("validate tokenString failed !!!", err)
	}
	return
}
func GetDateTimeFormat(days int, pattern string) string {
	now := time.Now()
	// nowDate := now.Format("2006-01-02")
	// nowTime := now.Format("2006-01-02 15:04:05")
	days = 24 * days
	daysStr := strconv.Itoa(days) + "h"
	d, _ := time.ParseDuration(daysStr)
	new := now.Add(d)
	var newDate string
	if pattern == "yyyy-MM-dd" {
		newDate = new.Format(Date)
	}

	return newDate
}

func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", API_KEY, SECRET_KEY)
	resp, _ := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	accessTokenObj := map[string]string{}
	json.Unmarshal([]byte(body), &accessTokenObj)
	return accessTokenObj["access_token"]
}

// 将结构体转换为 map[string]interface{}，并且字段名为小写
func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	objValue := reflect.ValueOf(obj)
	objType := reflect.TypeOf(obj)

	// 确保传入的是结构体或结构体指针
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
		objType = objType.Elem()
	}

	// 遍历结构体的字段
	for i := 0; i < objValue.NumField(); i++ {
		field := objType.Field(i)
		fieldName := strings.ToLower(field.Name) // 字段名转为小写
		fieldValue := objValue.Field(i).Interface()
		result[fieldName] = fieldValue
	}

	return result
}

// 脱敏处理
func GetDesensitization(content string, typeName string) string {
	runes := []rune(content)
	switch typeName {
	case "名":
		return string(runes[:1]) + strings.Repeat("*", len(runes)-1)
	case "手":
		if len(runes) < 11 {
			return content
		} else {
			return string(runes[:3]) + strings.Repeat("*", len(runes)-7) + string(runes[len(runes)-4:])
		}
	case "身":
		if len(runes) < 18 {
			return content
		} else {
			return string(runes[:6]) + strings.Repeat("*", 10) + string(runes[16:])
		}
	case "卡":
		if len(runes) < 10 {
			return content
		} else {
			return string(runes[:6]) + strings.Repeat("*", len(runes)-10) + string(runes[len(runes)-4:])
		}
	case "邮":
		strs := strings.Split(content, "@")
		if len(strs) >= 2 {
			prefix := strs[0]
			suffix := strs[1]
			return prefix[:1] + strings.Repeat("*", len(prefix)-1) + "@" + suffix
		} else {
			return content
		}
	default:
		return content
	}
}

// 余弦相似度函数
func CosineSimilarity(a, b map[string]int) float64 {
	numerator := 0
	aDenominator := 0
	bDenominator := 0
	for key, valA := range a {
		if valB, exists := b[key]; exists {
			numerator += valA * valB
		}
		aDenominator += valA * valA
	}
	for _, valB := range b {
		bDenominator += valB * valB
	}
	return float64(numerator) / (math.Sqrt(float64(aDenominator)) * math.Sqrt(float64(bDenominator)))
}

// SendRequest 向 Coze 智能体发送请求
func (c *CozeClient) SendRequest(endpoint string, payload interface{}) ([]byte, error) {
	// 将请求数据转换为 JSON 格式
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// 构建请求 URL
	url := fmt.Sprintf("%s%s", c.baseURL, endpoint)

	// 创建 HTTP POST 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}
func JiaoYangPrint() {
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("|                      欢迎使用fangwuzulinsystem                        |")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("|                       服务已启动，端口：8080                        |")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("| 后台运行地址：http://localhost:8080/fangwuzulinsystem/admin/dist/#/login")
	fmt.Println("---------------------------------------------------------------------")
}
