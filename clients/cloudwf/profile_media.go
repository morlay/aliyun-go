package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileMediaRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r ProfileMediaRequest) Invoke(client *sdk.Client) (response *ProfileMediaResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ProfileMediaRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileMedia", "", "")

	resp := struct {
		*responses.BaseResponse
		ProfileMediaResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ProfileMediaResponse

	err = client.DoAction(&req, &resp)
	return
}

type ProfileMediaResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
