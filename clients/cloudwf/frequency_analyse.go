package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FrequencyAnalyseRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r FrequencyAnalyseRequest) Invoke(client *sdk.Client) (response *FrequencyAnalyseResponse, err error) {
	req := struct {
		*requests.RpcRequest
		FrequencyAnalyseRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "FrequencyAnalyse", "", "")

	resp := struct {
		*responses.BaseResponse
		FrequencyAnalyseResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.FrequencyAnalyseResponse

	err = client.DoAction(&req, &resp)
	return
}

type FrequencyAnalyseResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
