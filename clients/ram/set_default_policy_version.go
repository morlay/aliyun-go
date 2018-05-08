package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetDefaultPolicyVersionRequest struct {
	requests.RpcRequest
	VersionId  string `position:"Query" name:"VersionId"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (req *SetDefaultPolicyVersionRequest) Invoke(client *sdk.Client) (resp *SetDefaultPolicyVersionResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "SetDefaultPolicyVersion", "", "")
	resp = &SetDefaultPolicyVersionResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetDefaultPolicyVersionResponse struct {
	responses.BaseResponse
	RequestId common.String
}
