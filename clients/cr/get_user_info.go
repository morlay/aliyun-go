package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetUserInfoRequest struct {
}

func (r GetUserInfoRequest) Invoke(client *sdk.Client) (response *GetUserInfoResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetUserInfoRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "GetUserInfo", "/users", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetUserInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetUserInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetUserInfoResponse struct {
}
