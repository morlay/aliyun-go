package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileConsumeRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r ProfileConsumeRequest) Invoke(client *sdk.Client) (response *ProfileConsumeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ProfileConsumeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileConsume", "", "")

	resp := struct {
		*responses.BaseResponse
		ProfileConsumeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ProfileConsumeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ProfileConsumeResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
