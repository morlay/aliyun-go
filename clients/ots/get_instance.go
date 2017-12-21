package ots

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetInstanceRequest struct {
	Accept      string `position:"Header" name:"Accept"`
	VersionName string `position:"Header" name:"VersionName"`
}

func (r GetInstanceRequest) Invoke(client *sdk.Client) (response *GetInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ots", "2013-09-12", "GetInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		GetInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetInstanceResponse struct {
	RequestId string
}
