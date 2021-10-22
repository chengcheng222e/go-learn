package main

import (
	"encoding/json"
	"fmt"
)

// Action :
type Action struct {
	Reason string `json:"reason"`
	Data   struct {
		IpeEquipmentTypeID int    `json:"IpeEquipmentTypeId"`
		IpeStatusID        int    `json:"IpeStatusId"`
		SegmentID          int    `json:"SegmentId"`
		MultiRoom          int    `json:"MultiRoom"`
		IpeAddress         string `json:"IpeAddress"`
		IpeEquipmentID     int    `json:"IpeEquipmentId"`
		IpeIPTypeID        int    `json:"IpeIpTypeId"`
	} `json:"data"` // 如果字段数量不可收敛的话，直接用 map[string]string
	Condition map[string]string `json:"condition"`
}

type CommonReq struct {
	Params struct {
		Content struct {
			RequestInfo struct {
				Operator      string `json:"operator"`
				SceneID       string `json:"sceneId"`
				SystemID      string `json:"systemId"`
				RequestModule string `json:"requestModule"`
			} `json:"requestInfo"`
			Version  string              `json:"version"`
			Type     string              `json:"type"`
			SchemeID string              `json:"schemeId"`
			Actions  []map[string]Action `json:"actions"`
		} `json:"content"`
	} `json:"params"`
}

func main() {
	body := "{\"params\": {\"content\": {\"orderBy\": \"\", \"requestInfo\": {\"sceneId\": 3477, \"operator\": \"devcloud\", \"systemId\": \"201409050\", \"requestModule\": \"\"}, \"searchCondition\": {\"SvrAssetId\": \"TCDEVNET210827001906\"}, \"resultColumn\": {\"SvrAssetId\": \"\", \"DeptId\": \"\", \"SfwId\": \"\", \"SvrBakOperator\": \"\", \"SvrOperator\": \"\", \"SfwName\": \"\", \"DeptName\": \"\", \"SvrId\": \"\", \"ProductId\": \"\", \"BsiId\": \"\", \"serverDisabledIP\": \"\", \"BsiPath\": \"\", \"serverProductCode\": \"\", \"serverProduct\": \"\", \"serverAllIP\": \"\"}, \"version\": \"1.0\", \"schemeId\": \"Server\", \"pagingInfo\": {\"returnTotalRows\": 1, \"startIndex\": 0, \"pageSize\": 1}, \"type\": \"Json\"}}}"
	bodyBytes := []byte(body)
	payload := CommonReq{}
	if err := json.Unmarshal(bodyBytes, &payload); err != nil {
		fmt.Print(err)
	}

}
