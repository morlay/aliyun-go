package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetStatisticsRequest struct {
	requests.RpcRequest
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
}

func (req *GetStatisticsRequest) Invoke(client *sdk.Client) (resp *GetStatisticsResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "GetStatistics", "", "")
	resp = &GetStatisticsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetStatisticsResponse struct {
	responses.BaseResponse
	RequestId string
	Code      string
	Success   bool
	Message   string
	Data      GetStatisticsData
}

type GetStatisticsData struct {
	Account int
	Health  int
	Patch   int
	Trojan  int
}
