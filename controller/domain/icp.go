package domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ops_tool/restapi"

	"github.com/gin-gonic/gin"
)

type ICPInfoResponse struct {
	Code      int    `json:"code"`
	Name      string `json:"name"`
	Nature    string `json:"nature"`
	Icp       string `json:"icp"`
	SiteName  string `json:"site_name"`
	SiteIndex string `json:"site_index"`
	SiteTime  string `json:"site_time"`
}

func ICPInfo(c *gin.Context) {
	response := restapi.Response{}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			response = restapi.Response{
				Code: restapi.Failed,
				Msg:  "查询失败",
			}
			response.Response(c, restapi.BadRequest)
		}
	}()

	domain := c.Param("domain")
	res, err := http.Get(fmt.Sprintf("https://api.oick.cn/icp/api.php?url=%s", domain))

	if err != nil {
		fmt.Println(err)
		panic("请求失败")
	}
	var icpInfo ICPInfoResponse
	result, err1 := ioutil.ReadAll(res.Body)
	err2 := json.Unmarshal(result, &icpInfo)
	if err1 != nil || err2 != nil || icpInfo.Code != 200 {
		panic("查询错误")
	}
	response = restapi.Response{
		Code: restapi.Success,
		Data: gin.H{
			"domain":     domain,
			"name":       icpInfo.Name,
			"nature":     icpInfo.Nature,
			"icp":        icpInfo.Icp,
			"site_name":  icpInfo.SiteName,
			"site_index": icpInfo.SiteIndex,
			"site_time":  icpInfo.SiteTime,
		},
	}
	response.Response(c, restapi.OK)
}
