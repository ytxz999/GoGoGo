package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 最外层响应
type CommentResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Replies []Reply `json:"replies"`
	} `json:"data"`
}

// 评论结构
type Reply struct {
	Rpid    int64 `json:"rpid"`
	Mid     int64 `json:"mid"`
	Content struct {
		Message string `json:"message"`
	} `json:"content"`
	// 子评论
	Replies []Reply `json:"replies"`
}

func main() {
	client := http.Client{}
	url := "https://api.bilibili.com/x/v2/reply/wbi/main?oid=420981979&" +
		"type=1&mode=3&pagination_str=%7B%22offset%22:%22%22%7D&plat=1&seek_rpid=&web_location=1315875&" +
		"w_rid=8c01d643aec702111f56060f1af25342&wts=1785832808"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	//浏览器身份
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/151.0.0.0 Safari/537.36 Edg/151.0.0.0")
	//来源
	req.Header.Set("referer", "https://www.bilibili.com/video/BV12341117rG/?vd_source=12746e71cfcc0fafde9f7daf4b9f1523")
	//cookie
	req.Header.Set("cookie", "enable_web_push=DISABLE; DedeUserID=480081200; DedeUserID__ckMd5=a2d951a2d80176a5; theme-tip-show=SHOWED; buvid3=DB442396-10C7-8AF9-6383-72223EDD8F7A58019infoc; b_nut=1752585158; rpdid=0zbfAHMJeu|hOG98atv|3FZ|3w1UOkiw; buvid4=E476AEFE-141E-CE42-98CF-EA731F64EB5291256-024071501-PJTLv2aQ3IMUvRtFTFKU2AbtzyIrl08OybsubTq9ELjovxpb+ZGW17RK1Zgugbwu; theme-switch-show=SHOWED; theme_style=light; PVID=1; SESSDATA=c5a58860%2C1789469111%2C7310e%2A31CjAMDZlh609q09BPnKlSvSEDw68lJ9EiXdpi-fLSzjibXyqotPyHw1JCyiyotyYzflQSVm41c3FWTDdWZFl5VEZTYmptU0U0bDhKTVZJM0NzR0J1OGtLRWlXQnZRcTFNQmd5dDl3emhnY2E0TkxWc1hMMTBtNFNLNTVjSVVFVllqRW84TEhTdU1RIIEC; bili_jct=12db2f9a5d3aa1989857c794a314e459; LIVE_BUVID=AUTO9517746715560043; buvid_fp=5441b5e436ad9231e97e7e980869a898; theme-avatar-tip-show=SHOWED; _uuid=FCA474E6-F28A-85E8-E427-87B3DAFA8D2879327infoc; CURRENT_QUALITY=80; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3ODU4NTIyMjMsImlhdCI6MTc4NTU5Mjk2MywicGx0IjotMX0.56c7y8DjYjn9VFr4c_S5u0TEyDJOa3EXdhejT_xdaKs; bili_ticket_expires=1785852163; home_feed_column=5; browser_resolution=1545-850; sid=8ecwv2n0; bp_t_offset_480081200=1232585107768868864; CURRENT_FNVAL=4048; b_lsid=60B0C8DF_19FCBEE2018")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var result CommentResponse

	err = json.Unmarshal(
		body,
		&result,
	)
	if err != nil {
		panic(err)
	}
	for _, reply := range result.Data.Replies {
		printReply(reply)
	}

}

func printReply(reply Reply) {
	fmt.Println("评论ID:", reply.Rpid)
	fmt.Println("用户:", reply.Mid)
	fmt.Println("内容:", reply.Content.Message)
	fmt.Println("----------------")
	//处理子评论
	for _, child := range reply.Replies {
		fmt.Print("子评论:")
		printReply(child)
	}
}
