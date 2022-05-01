package domain

import (
	"encoding/json"
	"fmt"
	"ops_tool/controller/ws"
	"ops_tool/db"
	"ops_tool/module"
	"ops_tool/restapi"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/miekg/dns"
)

type Resolve struct {
	DNSServer string      `json:"dns_server"`
	Name      string      `json:"name"`
	Type      interface{} `json:"type"`
	Data      interface{} `json:"data"`
	Status    bool        `json:"status"`
}

func DNSResolve(c *gin.Context) {
	client_tmp, _ := c.Get("ws")
	client := client_tmp.(*ws.Client)
	dnsServer := []module.DNSServer{}
	db.MySQL.Find(&dnsServer)
	domain := c.Param("domain")
	for _, server := range dnsServer {
		response := restapi.Response{}
		dnsClient := dns.Client{}
		message := dns.Msg{}
		message.SetQuestion(domain+".", dns.TypeA)
		r, _, err := dnsClient.Exchange(&message, server.ServerIP+":53")
		if err != nil || len(r.Answer) == 0 {
			response = restapi.Response{
				Code: restapi.Failed,
				Msg:  "查询失败",
				Data: []Resolve{Resolve{
					Name:      domain,
					DNSServer: server.ServerIP,
					Status:    false,
				}},
			}
		} else {
			time.Sleep(time.Microsecond * 50)
			resolves := []Resolve{}
			fmt.Println(len(r.Answer))
			fmt.Println(r.Answer)
			for _, ans := range r.Answer {
				record, isType := ans.(*dns.A)
				if isType {
					resolves = append(resolves, Resolve{
						Name:      record.Hdr.Name,
						DNSServer: server.ServerIP,
						Type:      "A",
						Data:      record.A,
						Status:    true,
					})
				}
				record1, isType := ans.(*dns.CNAME)
				if isType {
					resolves = append(resolves, Resolve{
						Name:      record1.Hdr.Name,
						DNSServer: server.ServerIP,
						Type:      "CNAME",
						Data:      record1.Target,
						Status:    true,
					})
				}

				v, _ := json.Marshal(resolves)
				fmt.Println(string(v))
			}
			response = restapi.Response{
				Code: restapi.Success,
				Data: resolves,
			}
		}
		result, _ := json.Marshal(&response)
		ws.WebsocketManager.Send(client.Id, client.Group, result)
	}
	time.Sleep(time.Second * 3)
	conn, err := c.Get("ws_conn")
	fmt.Println(err)
	conn.(*websocket.Conn).Close()
}
