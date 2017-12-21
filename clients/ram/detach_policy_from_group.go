package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetachPolicyFromGroupRequest struct {
	requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	GroupName  string `position:"Query" name:"GroupName"`
}

func (req *DetachPolicyFromGroupRequest) Invoke(client *sdk.Client) (resp *DetachPolicyFromGroupResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "DetachPolicyFromGroup", "", "")
	resp = &DetachPolicyFromGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetachPolicyFromGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
