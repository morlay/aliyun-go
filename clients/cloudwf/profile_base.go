package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileBaseRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r ProfileBaseRequest) Invoke(client *sdk.Client) (response *ProfileBaseResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ProfileBaseRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileBase", "", "")

	resp := struct {
		*responses.BaseResponse
		ProfileBaseResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ProfileBaseResponse

	err = client.DoAction(&req, &resp)
	return
}

type ProfileBaseResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
