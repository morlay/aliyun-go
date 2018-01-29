package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveIpControlPolicyItemRequest struct {
	requests.RpcRequest
	IpControlId   string `position:"Query" name:"IpControlId"`
	PolicyItemIds string `position:"Query" name:"PolicyItemIds"`
}

func (req *RemoveIpControlPolicyItemRequest) Invoke(client *sdk.Client) (resp *RemoveIpControlPolicyItemResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveIpControlPolicyItem", "apigateway", "")
	resp = &RemoveIpControlPolicyItemResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveIpControlPolicyItemResponse struct {
	responses.BaseResponse
	RequestId string
}
