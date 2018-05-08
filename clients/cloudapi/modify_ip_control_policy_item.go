package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyIpControlPolicyItemRequest struct {
	requests.RpcRequest
	IpControlId  string `position:"Query" name:"IpControlId"`
	PolicyItemId string `position:"Query" name:"PolicyItemId"`
	AppId        string `position:"Query" name:"AppId"`
	CidrIp       string `position:"Query" name:"CidrIp"`
}

func (req *ModifyIpControlPolicyItemRequest) Invoke(client *sdk.Client) (resp *ModifyIpControlPolicyItemResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyIpControlPolicyItem", "apigateway", "")
	resp = &ModifyIpControlPolicyItemResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyIpControlPolicyItemResponse struct {
	responses.BaseResponse
	RequestId common.String
}
