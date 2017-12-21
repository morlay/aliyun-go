package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CloseCCProtectRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (r CloseCCProtectRequest) Invoke(client *sdk.Client) (response *CloseCCProtectResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CloseCCProtectRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "CloseCCProtect", "", "")

	resp := struct {
		*responses.BaseResponse
		CloseCCProtectResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CloseCCProtectResponse

	err = client.DoAction(&req, &resp)
	return
}

type CloseCCProtectResponse struct {
	RequestId string
}
