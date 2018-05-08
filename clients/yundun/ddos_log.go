package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DdosLogRequest struct {
	requests.RpcRequest
	JstOwnerId   int64  `position:"Query" name:"JstOwnerId"`
	PageNumber   int    `position:"Query" name:"PageNumber"`
	PageSize     int    `position:"Query" name:"PageSize"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (req *DdosLogRequest) Invoke(client *sdk.Client) (resp *DdosLogResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "DdosLog", "", "")
	resp = &DdosLogResponse{}
	err = client.DoAction(req, resp)
	return
}

type DdosLogResponse struct {
	responses.BaseResponse
	RequestId    common.String
	AttackStatus common.Integer
	StartTime    common.String
	EndTime      common.String
	PageNumber   common.Integer
	PageSize     common.Integer
	TotalCount   common.Integer
	LogList      DdosLogDdosLogList
}

type DdosLogDdosLog struct {
	StartTime    common.String
	EndTime      common.String
	Reason       common.String
	Status       common.Integer
	Bps          common.Long
	Pps          common.Long
	Qps          common.Long
	AttackType   common.String
	AttackIpList common.String
	Type         common.Integer
}

type DdosLogDdosLogList []DdosLogDdosLog

func (list *DdosLogDdosLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DdosLogDdosLog)
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
