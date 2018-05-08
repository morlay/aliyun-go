package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetCrackStatisticsRequest struct {
	requests.RpcRequest
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
}

func (req *GetCrackStatisticsRequest) Invoke(client *sdk.Client) (resp *GetCrackStatisticsResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "GetCrackStatistics", "vipaegis", "")
	resp = &GetCrackStatisticsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetCrackStatisticsResponse struct {
	responses.BaseResponse
	RequestId common.String
	Code      common.String
	Success   bool
	Message   common.String
	Data      GetCrackStatisticsData
}

type GetCrackStatisticsData struct {
	Intercepted common.Integer
}
