package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DetachPolicyFromUserRequest struct {
	requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	UserName   string `position:"Query" name:"UserName"`
}

func (req *DetachPolicyFromUserRequest) Invoke(client *sdk.Client) (resp *DetachPolicyFromUserResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "DetachPolicyFromUser", "", "")
	resp = &DetachPolicyFromUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetachPolicyFromUserResponse struct {
	responses.BaseResponse
	RequestId common.String
}
