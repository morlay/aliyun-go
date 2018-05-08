package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteAccessRuleRequest struct {
	requests.RpcRequest
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	AccessRuleId    string `position:"Query" name:"AccessRuleId"`
}

func (req *DeleteAccessRuleRequest) Invoke(client *sdk.Client) (resp *DeleteAccessRuleResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "DeleteAccessRule", "nas", "")
	resp = &DeleteAccessRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteAccessRuleResponse struct {
	responses.BaseResponse
	RequestId common.String
}
