package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetStatisticsRequest struct {
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
}

func (r GetStatisticsRequest) Invoke(client *sdk.Client) (response *GetStatisticsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetStatisticsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("aegis", "2016-11-11", "GetStatistics", "", "")

	resp := struct {
		*responses.BaseResponse
		GetStatisticsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetStatisticsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetStatisticsResponse struct {
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
