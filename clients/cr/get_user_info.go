package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetUserInfoRequest struct {
	requests.RoaRequest
}

func (req *GetUserInfoRequest) Invoke(client *sdk.Client) (resp *GetUserInfoResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetUserInfo", "/users", "", "")
	req.Method = "GET"

	resp = &GetUserInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetUserInfoResponse struct {
	responses.BaseResponse
}
