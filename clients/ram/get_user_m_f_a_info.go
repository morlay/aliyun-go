package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetUserMFAInfoRequest struct {
	UserName string `position:"Query" name:"UserName"`
}

func (r GetUserMFAInfoRequest) Invoke(client *sdk.Client) (response *GetUserMFAInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetUserMFAInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "GetUserMFAInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		GetUserMFAInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetUserMFAInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetUserMFAInfoResponse struct {
	RequestId string
	MFADevice GetUserMFAInfoMFADevice
}

type GetUserMFAInfoMFADevice struct {
	SerialNumber string
}
