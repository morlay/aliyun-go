package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetVpcAccessRequest struct {
	requests.RpcRequest
	VpcId      string `position:"Query" name:"VpcId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Port       int    `position:"Query" name:"Port"`
	Name       string `position:"Query" name:"Name"`
}

func (req *SetVpcAccessRequest) Invoke(client *sdk.Client) (resp *SetVpcAccessResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SetVpcAccess", "apigateway", "")
	resp = &SetVpcAccessResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetVpcAccessResponse struct {
	responses.BaseResponse
	RequestId string
}
