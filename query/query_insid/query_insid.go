package query_insid

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAllInfoByType(Instype string) []byte {
	if Instype != "SPOT" && Instype != "MARGIN" && Instype != "SWAP" && Instype != "FUTURES" && Instype != "OPTION" {
		panic("type doesn't existed")
	}
	response, err := http.Get("https://www.okx.com/api/v5/public/instruments?instType=" + Instype)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return body
}

// 指定查询对象类型
func GetInsByType(Instype string) []string {
	body := GetAllInfoByType(Instype)
	var basket map[string]interface{}
	err := json.Unmarshal(body, &basket)
	if err != nil {
		fmt.Println(err)
	}
	temp := basket["data"].([]interface{})
	answer := make([]string, 0)
	for i := 0; i < len(temp); i++ {
		answer = append(answer, temp[i].(map[string]interface{})["instId"].(string))
	}
	return answer
}

// 只用于查询交割合约，可以指定alias
func GetInsByTypeAndAlias(Instype string, Alias string) []string {
	body := GetAllInfoByType(Instype)
	var basket map[string]interface{}
	err := json.Unmarshal(body, &basket)
	if err != nil {
		fmt.Println(err)
	}
	temp := basket["data"].([]interface{})
	answer := make([]string, 0)
	for i := 0; i < len(temp); i++ {
		if temp[i].(map[string]interface{})["alias"].(string) == Alias {
			answer = append(answer, temp[i].(map[string]interface{})["instId"].(string))
		}
	}
	return answer
}

// func main() {
// 	a := GetInsByType("MARGIN")
// 	fmt.Println(len(a))
// }
