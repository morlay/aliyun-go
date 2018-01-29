package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddIpControlPolicyItemRequest struct {
	requests.RpcRequest
	IpControlId string `position:"Query" name:"IpControlId"`
	AppId       string `position:"Query" name:"AppId"`
	CidrIp      string `position:"Query" name:"CidrIp"`
}

func (req *AddIpControlPolicyItemRequest) Invoke(client *sdk.Client) (resp *AddIpControlPolicyItemResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "AddIpControlPolicyItem", "apigateway", "")
	resp = &AddIpControlPolicyItemResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddIpControlPolicyItemResponse struct {
	responses.BaseResponse
	RequestId    string
	PolicyItemId string
}
