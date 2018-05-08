package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DdosFlowGraphRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (req *DdosFlowGraphRequest) Invoke(client *sdk.Client) (resp *DdosFlowGraphResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "DdosFlowGraph", "", "")
	resp = &DdosFlowGraphResponse{}
	err = client.DoAction(req, resp)
	return
}

type DdosFlowGraphResponse struct {
	responses.BaseResponse
	RequestId   common.String
	NormalFlows DdosFlowGraphNormalFlowList
	TotalFlows  DdosFlowGraphTotalFlowList
}

type DdosFlowGraphNormalFlow struct {
	Time    common.Long
	BitRecv common.Long
	BitSend common.Long
	PktRecv common.Long
	PktSend common.Long
}

type DdosFlowGraphTotalFlow struct {
	Time    common.Long
	BitRecv common.Long
	PktRecv common.Long
}

type DdosFlowGraphNormalFlowList []DdosFlowGraphNormalFlow

func (list *DdosFlowGraphNormalFlowList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DdosFlowGraphNormalFlow)
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

type DdosFlowGraphTotalFlowList []DdosFlowGraphTotalFlow

func (list *DdosFlowGraphTotalFlowList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DdosFlowGraphTotalFlow)
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
