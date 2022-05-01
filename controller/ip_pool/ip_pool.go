package ippool

import (
	"ops_tool/restapi"

	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

func GetIPInfoByIP(c *gin.Context) {
	response := restapi.Response{}
	ip := c.Param("ip")
	if ip == "" {
		ip = c.Request.Header.Get("X-Forward-For")
	}
	record, err := GetIPInfo(ip)
	if err != nil {
		response = restapi.Response{
			Code: restapi.Failed,
			Msg:  err.Error(),
		}
		response.Response(c, restapi.BadRequest)
		return
	}
	response = restapi.Response{
		Code: restapi.Success,
		Data: gin.H{
			"ip":       ip,
			"isp":      record.ISP,
			"city":     record.City,
			"country":  record.Country,
			"region":   record.Region,
			"province": record.Province,
		},
	}
	response.Response(c, restapi.OK)
}

func GetIPInfo(ip string) (ipinfo ip2region.IpInfo, err error) {
	db, err := ip2region.New("ip2region.db")
	if err != nil {
		return
	}
	ipinfo, err = db.BinarySearch(ip)
	if err != nil {
		return
	}
	return
}
