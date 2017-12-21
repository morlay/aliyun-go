package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDdosConfigRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (r QueryDdosConfigRequest) Invoke(client *sdk.Client) (response *QueryDdosConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryDdosConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "QueryDdosConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryDdosConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryDdosConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryDdosConfigResponse struct {
	RequestId        string
	Bps              int64
	Pps              int64
	Qps              int64
	Sipconn          int64
	Sipnew           int64
	Layer7Config     bool
	FlowPosition     int
	QpsPosition      int
	StrategyPosition int
	Level            int
	HoleBps          string
	ConfigType       string
}
