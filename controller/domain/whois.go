package domain

import (
	"fmt"
	"io/ioutil"
	"ops_tool/restapi"
	"strings"

	"github.com/domainr/whois"
	"github.com/gin-gonic/gin"
)

const (
	DomainName                 = "Domain Name"
	WhoisServer                = "Registrar WHOIS Server"
	RegisterarURL              = "Registrar URL"
	UpdatedDate                = "Updated Date"
	CreationDate               = "Creation Date"
	RegistryExpiryDate         = "Registry Expiry Date"
	Registrar                  = "Registrar:"
	RegistrarAbuseContactEmail = "Registrar Abuse Contact Email: "
	NameServer                 = "Name Server: "
)

func GetWhoisInfo(c *gin.Context) {
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
	request, _ := whois.NewRequest(domain)
	res, _ := whois.DefaultClient.Fetch(request)
	record, _ := res.Reader()
	whois, _ := ioutil.ReadAll(record)
	whois_str := string(whois)
	fmt.Println(whois_str)
	if registed(whois_str) {
		response = restapi.Response{
			Code: restapi.Failed,
			Msg:  "没有查询到",
		}
		response.Response(c, restapi.BadRequest)
		return
	}
	response = restapi.Response{
		Code: restapi.Success,
		Data: gin.H{
			"dom_upddate":   find(whois_str, UpdatedDate),
			"domain":        find(whois_str, DomainName),
			"registrar":     find(whois_str, Registrar),
			"dom_insdate":   find(whois_str, CreationDate),
			"dom_expdate":   find(whois_str, RegistryExpiryDate),
			"contact_email": find(whois_str, RegistrarAbuseContactEmail),
			"name_server":   find(whois_str, NameServer),
			"details":       whois_str,
		},
	}
	response.Response(c, restapi.OK)
}

func registed(whois string) (result bool) {
	if find := strings.Contains(whois, "No match"); find {
		return true
	}
	return false
}

func find(whois string, key string) string {
	whois_list := strings.Split(whois, "\n")
	var result []string
	for _, line := range whois_list {
		if find := strings.Contains(line, key); find {
			result = append(result, strings.Split(line, ": ")[1])
		}
	}
	fmt.Println(result)
	return strings.Join(result, "\n")
}
