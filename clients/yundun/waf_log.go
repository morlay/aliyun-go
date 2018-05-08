package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type WafLogRequest struct {
	requests.RpcRequest
	JstOwnerId   int64  `position:"Query" name:"JstOwnerId"`
	PageNumber   int    `position:"Query" name:"PageNumber"`
	PageSize     int    `position:"Query" name:"PageSize"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (req *WafLogRequest) Invoke(client *sdk.Client) (resp *WafLogResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "WafLog", "", "")
	resp = &WafLogResponse{}
	err = client.DoAction(req, resp)
	return
}

type WafLogResponse struct {
	responses.BaseResponse
	RequestId   common.String
	WebAttack   common.Integer
	NewWafUser  bool
	WafOpened   bool
	InWhiteList bool
	DomainCount common.Integer
	StartTime   common.String
	EndTime     common.String
	PageNumber  common.Integer
	PageSize    common.Integer
	TotalCount  common.Integer
	LogList     WafLogWafLogList
}

type WafLogWafLog struct {
	SourceIp common.String
	Time     common.String
	Url      common.String
	Type     common.String
	Status   common.Integer
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
