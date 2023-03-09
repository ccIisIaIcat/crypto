package main

import (
	"encoding/json"
	"fmt"
	"global"
	"io"
	"log"
	"net/http"
)

func GetInsIdInfo(InsType string, InsList []string) map[string]*global.InsidBasicInfo {
	res, err := http.Get("https://aws.okx.com/api/v5/public/instruments?instType=" + InsType)
	if err != nil {
		log.Println(err)
	}
	tool_map := make(map[string]bool)
	for i := 0; i < len(InsList); i++ {
		tool_map[InsList[i]] = true
	}
	ans_map := make(map[string]*global.InsidBasicInfo)
	temp_json, _ := io.ReadAll(res.Body)
	var temp map[string]interface{}
	json.Unmarshal(temp_json, &temp)
	info_list := temp["data"].([]interface{})
	for i := 0; i < len(info_list); i++ {
		if _, ok := tool_map[info_list[i].(map[string]interface{})["instId"].(string)]; ok {
			temp_info := info_list[i].(map[string]interface{})
			temp_basic := &global.InsidBasicInfo{}
			temp_basic.InsId = temp_info["instId"].(string)
			temp_basic.CtVal = temp_info["ctVal"].(string)
			temp_basic.CtValCcy = temp_info["ctValCcy"].(string)
			temp_basic.TickSz = temp_info["tickSz"].(string)
			temp_basic.LotSz = temp_info["lotSz"].(string)
			temp_basic.MinSz = temp_info["minSz"].(string)
			temp_basic.MaxLmtSz = temp_info["maxLmtSz"].(string)
			temp_basic.MaxMktSz = temp_info["maxMktSz"].(string)
			ans_map[temp_basic.InsId] = temp_basic
		}
	}
	return ans_map
}

func main() {
	a := GetInsIdInfo("SWAP", []string{"ETH-USDT-SWAP", "BTC-USDT-SWAP"})
	fmt.Println(a["ETH-USDT-SWAP"])
	lala, _ := json.Marshal(a)
	fmt.Println(string(lala))
}
