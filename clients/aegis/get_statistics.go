package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetStatisticsRequest struct {
	requests.RpcRequest
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
}

func (req *GetStatisticsRequest) Invoke(client *sdk.Client) (resp *GetStatisticsResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "GetStatistics", "vipaegis", "")
	resp = &GetStatisticsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetStatisticsResponse struct {
	responses.BaseResponse
	RequestId common.String
	Code      common.String
	Success   bool
	Message   common.String
	Data      GetStatisticsData
}

type GetStatisticsData struct {
	Account common.Integer
	Health  common.Integer
	Patch   common.Integer
	Trojan  common.Integer
}
