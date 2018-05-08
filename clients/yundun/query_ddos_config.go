package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryDdosConfigRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (req *QueryDdosConfigRequest) Invoke(client *sdk.Client) (resp *QueryDdosConfigResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "QueryDdosConfig", "", "")
	resp = &QueryDdosConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDdosConfigResponse struct {
	responses.BaseResponse
	RequestId        common.String
	Bps              common.Long
	Pps              common.Long
	Qps              common.Long
	Sipconn          common.Long
	Sipnew           common.Long
	Layer7Config     bool
	FlowPosition     common.Integer
	QpsPosition      common.Integer
	StrategyPosition common.Integer
	Level            common.Integer
	HoleBps          common.String
	ConfigType       common.String
}
