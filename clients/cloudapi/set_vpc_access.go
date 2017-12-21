package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetVpcAccessRequest struct {
	VpcId      string `position:"Query" name:"VpcId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Port       int    `position:"Query" name:"Port"`
	Name       string `position:"Query" name:"Name"`
}

func (r SetVpcAccessRequest) Invoke(client *sdk.Client) (response *SetVpcAccessResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetVpcAccessRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SetVpcAccess", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		SetVpcAccessResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetVpcAccessResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetVpcAccessResponse struct {
	RequestId string
}
