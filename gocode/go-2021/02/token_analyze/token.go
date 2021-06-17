package main

import (
	"errors"
	"fmt"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"log"
	"os"
	"strings"
)

const PrivateKeyString = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`
const PublicKeyString = `-----BEGIN CERTIFICATE-----
MIIFkTCCA3mgAwIBAgIJAM1qfyLQfpsqMA0GCSqGSIb3DQEBCwUAMF4xCzAJBgNV
BAYTAkNOMQ4wDAYDVQQIDAVzdGF0ZTELMAkGA1UEBwwCQ04xDTALBgNVBAoMBGRh
bmExDTALBgNVBAsMBGRhbmExFDASBgNVBAMMC2RhdGF0b20uY29tMCAXDTE4MDMw
NTAzMTMwM1oYDzIxMTgwMjA5MDMxMzAzWjBeMQswCQYDVQQGEwJDTjEOMAwGA1UE
CAwFc3RhdGUxCzAJBgNVBAcMAkNOMQ0wCwYDVQQKDARkYW5hMQ0wCwYDVQQLDARk
YW5hMRQwEgYDVQQDDAtkYXRhdG9tLmNvbTCCAiIwDQYJKoZIhvcNAQEBBQADggIP
ADCCAgoCggIBAOCQ4VBgOTn1btDHe3AyhT9KIDy/b6djhbJi0cE6Kkql+itG7oxx
n8MqIzgNd8j451aubcEhhUE+QTPiKbuyQHwlvigND2THG6X5tsGJqG/X4H71ziSR
YPV8VasPfFfZ6fr7kNh6jkC+TVGOpfXiQVnQsiUOOhWyC8roCh97hXxlRq1I6tH/
9V0XRQOGMPKNzbxfF6K1HoY1WuYP0X3HlO8EV+5gGuzT9H2P2Lae6IRCFzqQvMJw
NAdYkXPgvRMC9dP0i7xT+3KpNrZpH5iejwfoULkzTC3KhtIku1ApNp4kS5RJNZa0
Whr63vQcwq6DthLYVU64XIOHYTWMMcfMN+eAQD+jpFOA2GIkAM0sHJe5KNk0iuNI
OUjN6eZlvVRB8aFI0/PrFyJ7z6eGytDUec9J7n0RKLW53TVWP/F8mFF/4RjekeMb
ZnXxz/LU6WvtK+kBx/B6RdttB6oh/HIfKiuqpKBsJvtOGjuZJdRUfvtLaCWWI84t
uk3sEKSLvUb/kDzRra/Kdk88okmq+gI4Ey1zXTbL94wvB1B3XWlHVXz62MRaMTSN
0M45YZagVkhvDvRGnw/oSY2UmT1wODhCN5gyPNowz85d+i3W43S1vTNbqG5lpieC
lHJaSB06KFNTOr/YdSjFXHqf1V2SO7ZMDXgpNWoNSXXBB415gpz2gzJhAgMBAAGj
UDBOMB0GA1UdDgQWBBQlsHoRde5dvQeJl4mKZgvL1oSfODAfBgNVHSMEGDAWgBQl
sHoRde5dvQeJl4mKZgvL1oSfODAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUA
A4ICAQDRFssTgfajcS532wDCYjjqPc7Vl2nCu8lr+b4Q5e47lYQuMNMTdWLSBWPe
JqnwZK0t5m+INkyfiJC9IgeCdr5JkvNkiIsWYnCGIgnUo2G+uJYXcS1i3SUR9sb5
UDzI5rSzOsFO1lkFGXMBNPjuBiEhrCu+NNcY/oCMREHcP7JIY9MoyIz9OMa47MIq
274brXXtICp0D2I/cfiA0jcxt7gy1WuQlbp+qCJRCaYfOY7VXlhCiE8t+E7fBqGu
fJLkioLXqVnkxxEkPQq2xoBMTteAjq6TJlj5l37A0sZm5PWb2UnYgFeJiL1h7rOb
AFqNlchZ+ca6iW8oJbIURML5QCO2whGbREzqERR+Ca6jcsQE3nD5D9TX4wUkZpy7
tKSmT4uwFlr2is0aqZcwO0ql7afyDNwhOUTcpmf/htJbwUXb9uoBxFNuI3HsZxVC
b/mrRf4AzpLV1omp8JfOQxB4Y88J+G8DRJWOHcMG2eGOZtMO4gWx3YNArF/dFAbR
t+K8qI4HgZjm4PpblWMT24Jz/Bvu/MAqJPim6xR6tqsoUmKmaKwFA23AiQPBxN7T
rENr1gv1mk/QGDjJjlAu0Gw4lyiG/lt+tIGq9C/tzZd4Ppm5yCSRhVe+ILxvFQeD
tH/B9ZrQrcVbwFLF8NCQUwSHBzlS8d7wmb8tgM9jlG4nv96LkQ==
-----END CERTIFICATE-----`

var PrivateKey = []byte(PrivateKeyString)

var PublicKey = []byte(PublicKeyString)

//func GetToken() (tokenString string, err error) {
//
//	token := jwt.New(jwt.GetSigningMethod("RS256"))
//	claims, err := ggg()
//	//expire = time.Now().Add(time.Hour * 24 * 7)
//	claims["iss"] = "danastudio_auth_server"
//	claims["exp"] = 1583896190
//
//	key, err := jwt.ParseRSAPrivateKeyFromPEM(PrivateKey)
//
//	tokenString, err = token.SignedString(key)
//
//	return tokenString, err
//}

func parseToken(authHeader string) (*jwt.Token, error) {
	//var publicKeyFile = "D:\\code_space\\go\\src\\danastudio\\src\\datatom.com\\public.crt"
	//publicKey, _ := ioutil.ReadFile(publicKeyFile)

	//authHeader := "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDk3Mjk0NzUsImlhdCI6MTU0OTcyODg3NSwiaXNzIjoiZGFuYXN0dWRpb19hdXRoX3NlcnZlciIsIm1vZHVsZXMiOlsiZGV2Y2VudGVyIiwib3BlcmF0aW5nIl0sInJvbGVfaWQiOiJkZXZlbG9wZXIiLCJyb2xlX25hbWUiOiLlvIDlj5Hlt6XnqIvluIgiLCJ0a3ZlcnNpb24iOjAsInVzZXJfaWQiOiJBV2gxMGYtTjRxQWxDUVRjaVdMQSIsInVzZXJfbmFtZSI6InRlc3QwMiJ9.HAycILQpPZ0h3ptKS2x8jcPoSkfoLKuBWAVW0Q4U6mcgMhw3TGyi2zt0lK3ZAzURwx1BlSGchqbfBuN6eb3ziVqignKtc7daUOiBssI21ogd5BMMLgnaqwyRMMHfXS9Xvk787kjiMS4SfWH9TWVAo2tL8Httv8eCyP85L3rr0VHh5cNr5kwqD4auI-Xaf2m8s9fV1PDZ_vKfF0IcrKMqRuh9KRXXGVNCXbOB0EGkkwONSMjHKUTbR5KFn_-elWsUkAET-V2z6gpCMvHSs2Gwn-S2lvxnOXEyBzmD2dxu88VHx1kQQUZYUcV4fpWNydfayKIQIvEtMRHuEH6FxpFXsGHwl6Ym5de57V8N81INXI3V-yv2QJCIY9uGO19g3jD6G5_GT-1wRfZ5R-jqrvtcnqLWGtrAWl46nt8SOG9Up9Vi6Q2P-9AA_bqn6EdKtmJ4dGRZDqmVRcMZ31P_GkjYiu0rVbX9qi7DfXp2gW6DxTqWcPtcacY_-v3lBAQosQsPirTGfZ5gg97IhB4SFsX7mIQwiGLwlPrvNtKoZUOQ86KvIkFFqMaWfjwga--9o6ESBneXNFi45Kvc36Wb10W760-i382wUB7ulOxe9j4x6V0PP_D-QfyyhC6sDLXHP39-MMCgacwZ91EoNhRDt2wS5ubIepKnvzs6hXs7nLd6w64"
	parts := strings.SplitN(authHeader, " ", 2)
	return jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("RS256") != token.Method {
			return nil, errors.New("invalid signing algorithm")
		}
		return jwt.ParseRSAPublicKeyFromPEM(PublicKey)
	})
}

func getToken(token string) (*jwt.Token, error) {
	//var publicKeyFile = "D:\\code_space\\go\\src\\danastudio\\src\\datatom.com\\public.crt"
	//publicKey, _ := ioutil.ReadFile(publicKeyFile)
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("RS256") != token.Method {
			// return nil, errors.New("invalid signing algorithm")
			return nil, errors.New("无效的签名算法")
		}
		return jwt.ParseRSAPublicKeyFromPEM(PublicKey)
	})
}

func ggg(authHeader string) (jwt.MapClaims, error) {
	token, _ := parseToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)

	// 平台认证服务颁发的Token里默认有以下的几个参数
	//username := claims["user_name"].(string)
	//userID := claims["user_id"].(string)
	//roleID := claims["role_id"].(string)
	//roleName := claims["role_name"].(string)
	////modulesh := claims["modules"].([]interface{})
	//projectID := claims["project_id"].(string)
	//c:=claims["exp"].(float64)
	//d:=int64(c)
	//exp :=time.Unix(d, 0).Format("2006-01-02 15:04:05")
	//fmt.Println(exp)
	//fmt.Println(c)
	//fmt.Println(&modulesh)
	//print(username, userID, roleID, roleName+" ",projectID)
	return claims, nil
}

//token.SignedString(key)
func main() {
	//s:="Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTExMTE2NzYsImlhdCI6MTYxMDUwNjg3NiwiaXNzIjoiZGFuYXN0dWRpb19hdXRoX3NlcnZlciIsIm1vZHVsZXMiOlsiYWNjZXNzIiwibWluZGluZyIsIndvcmtmbG93IiwidmF1bHQiLCJtbGliIl0sInJvbGVfaWQiOiJkZXZlbG9wZXIiLCJyb2xlX25hbWUiOiLlvIDlj5HogIUiLCJ0a3ZlcnNpb24iOjAsInVzZXJfaWQiOiJBWE14aDZEUmNlM1RsdTAxQVljMiIsInVzZXJfbmFtZSI6ImxpdXBlbmcifQ.TGRRiMG4VZYXxIsXlRqKJGTefNnn5C6s1PaBHc4FcWY6r9o8tkMWUTLWD2JMZs9zhtnkUVKZ0Ytr8PO7_3LkQE6efrBTVNkpSFpmF6xWgbsc_RBFMyFqZhZz-NGHZGnudj3nCreDfQRtYjaqH6n3_yaWQR8cnQ0cfrCXX0Cks5xYsmmcZKDlPNXiMQI6BDAG4sjrB_mv8u7L6MHXhObwq1T4LTqzDxEu8DW2ZPKdzoLZ7kCKoXehgNbd7wzCluuRutzE7JkejkVAoEq-CtvUBajrHqpQa1KfdilpLaH-u56kHqkZIa5M9y7c7PcRTBLQNIFjfe_7TmejJhOJvXseZeUOE60w1XmfwuzVcoCISWRcXfRuvk7t5Lk0kNzK4LcHKRB9aJPkNIEd3ec1PWzvEO6bYkLKHLyhpK5DXKcAmMta4W-Tb0NEU7DSWN7GcCYrFxVDd0Pb2F38XkuzkQLHgBXSALBsjWaOWsFy38dzJOGenfnRlg5c6QIpPMoBaVyHnVydp8h0hYngnMx1tgYjVBLxakgPtMJXbBMBtLtL6FgHqT8F6PeXQPCQD8rdmx1jxkXIaFioh69uViBEzGKyspcWkhD23CnnJPqC-hN_jbEYknbX-ZtR5xhxlCUpEmFE6wSekF_PnkrRQJfBFrhGjYAYU_CcvE61l6bjDwRkJ18"
	//fmt.Println("命令行参数长度: ", len(os.Args))
	if len(os.Args) != 2 {
		log.Fatal("参数格式不正确，示例：./token \"Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTExMTE2NzYsImlhdCI6MTYxMDUwNjg3NiwiaXNzIjoiZGFuYXN0dWRpb19hdXRoX3NlcnZlciIsIm1vZHVsZXMiOlsiYWNjZXNzIiwibWluZGluZyIsIndvcmtmbG93IiwidmF1bHQiLCJtbGliIl0sInJvbGVfaWQiOiJkZXZlbG9wZXIiLCJyb2xlX25hbWUiOiLlvIDlj5HogIUiLCJ0a3ZlcnNpb24iOjAsInVzZXJfaWQiOiJBWE14aDZEUmNlM1RsdTAxQVljMiIsInVzZXJfbmFtZSI6ImxpdXBlbmcifQ.TGRRiMG4VZYXxIsXlRqKJGTefNnn5C6s1PaBHc4FcWY6r9o8tkMWUTLWD2JMZs9zhtnkUVKZ0Ytr8PO7_3LkQE6efrBTVNkpSFpmF6xWgbsc_RBFMyFqZhZz-NGHZGnudj3nCreDfQRtYjaqH6n3_yaWQR8cnQ0cfrCXX0Cks5xYsmmcZKDlPNXiMQI6BDAG4sjrB_mv8u7L6MHXhObwq1T4LTqzDxEu8DW2ZPKdzoLZ7kCKoXehgNbd7wzCluuRutzE7JkejkVAoEq-CtvUBajrHqpQa1KfdilpLaH-u56kHqkZIa5M9y7c7PcRTBLQNIFjfe_7TmejJhOJvXseZeUOE60w1XmfwuzVcoCISWRcXfRuvk7t5Lk0kNzK4LcHKRB9aJPkNIEd3ec1PWzvEO6bYkLKHLyhpK5DXKcAmMta4W-Tb0NEU7DSWN7GcCYrFxVDd0Pb2F38XkuzkQLHgBXSALBsjWaOWsFy38dzJOGenfnRlg5c6QIpPMoBaVyHnVydp8h0hYngnMx1tgYjVBLxakgPtMJXbBMBtLtL6FgHqT8F6PeXQPCQD8rdmx1jxkXIaFioh69uViBEzGKyspcWkhD23CnnJPqC-hN_jbEYknbX-ZtR5xhxlCUpEmFE6wSekF_PnkrRQJfBFrhGjYAYU_CcvE61l6bjDwRkJ18\"")
	}
	token := os.Args[1]
	claims, err := ggg(token)
	if err != nil {
		log.Fatal(err)
	}
	projectid := claims["project_id"].(string)
	fmt.Println(projectid)
}
