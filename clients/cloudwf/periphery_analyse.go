package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PeripheryAnalyseRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r PeripheryAnalyseRequest) Invoke(client *sdk.Client) (response *PeripheryAnalyseResponse, err error) {
	req := struct {
		*requests.RpcRequest
		PeripheryAnalyseRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "PeripheryAnalyse", "", "")

	resp := struct {
		*responses.BaseResponse
		PeripheryAnalyseResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PeripheryAnalyseResponse

	err = client.DoAction(&req, &resp)
	return
}

type PeripheryAnalyseResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
