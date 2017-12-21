package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileSetRequest struct {
	AutoInstall string `position:"Query" name:"AutoInstall"`
	UserId      int64  `position:"Query" name:"UserId"`
}

func (r ProfileSetRequest) Invoke(client *sdk.Client) (response *ProfileSetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ProfileSetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "ProfileSet", "cms", "")

	resp := struct {
		*responses.BaseResponse
		ProfileSetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ProfileSetResponse

	err = client.DoAction(&req, &resp)
	return
}

type ProfileSetResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
}
