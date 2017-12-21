package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type WafLogRequest struct {
	JstOwnerId   int64  `position:"Query" name:"JstOwnerId"`
	PageNumber   int    `position:"Query" name:"PageNumber"`
	PageSize     int    `position:"Query" name:"PageSize"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (r WafLogRequest) Invoke(client *sdk.Client) (response *WafLogResponse, err error) {
	req := struct {
		*requests.RpcRequest
		WafLogRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "WafLog", "", "")

	resp := struct {
		*responses.BaseResponse
		WafLogResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.WafLogResponse

	err = client.DoAction(&req, &resp)
	return
}

type WafLogResponse struct {
	RequestId   string
	WebAttack   int
	NewWafUser  bool
	WafOpened   bool
	InWhiteList bool
	DomainCount int
	StartTime   string
	EndTime     string
	PageNumber  int
	PageSize    int
	TotalCount  int
	LogList     WafLogWafLogList
}

type WafLogWafLog struct {
	SourceIp string
	Time     string
	Url      string
	Type     string
	Status   int
}

type WafLogWafLogList []WafLogWafLog

func (list *WafLogWafLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]WafLogWafLog)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
