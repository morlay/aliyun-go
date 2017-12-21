package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveVpcAccessRequest struct {
	VpcId      string `position:"Query" name:"VpcId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Port       int    `position:"Query" name:"Port"`
}

func (r RemoveVpcAccessRequest) Invoke(client *sdk.Client) (response *RemoveVpcAccessResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveVpcAccessRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveVpcAccess", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		RemoveVpcAccessResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveVpcAccessResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveVpcAccessResponse struct {
	RequestId string
}
