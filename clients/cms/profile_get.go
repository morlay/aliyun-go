package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileGetRequest struct {
	UserId int64 `position:"Query" name:"UserId"`
}

func (r ProfileGetRequest) Invoke(client *sdk.Client) (response *ProfileGetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ProfileGetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "ProfileGet", "cms", "")

	resp := struct {
		*responses.BaseResponse
		ProfileGetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ProfileGetResponse

	err = client.DoAction(&req, &resp)
	return
}

type ProfileGetResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
	UserId       int64
	AutoInstall  bool
}
