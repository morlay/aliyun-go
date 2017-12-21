package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyAccessRuleRequest struct {
	requests.RpcRequest
	RWAccessType    string `position:"Query" name:"RWAccessType"`
	SourceCidrIp    string `position:"Query" name:"SourceCidrIp"`
	UserAccessType  string `position:"Query" name:"UserAccessType"`
	Priority        int    `position:"Query" name:"Priority"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	AccessRuleId    string `position:"Query" name:"AccessRuleId"`
}

func (req *ModifyAccessRuleRequest) Invoke(client *sdk.Client) (resp *ModifyAccessRuleResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "ModifyAccessRule", "nas", "")
	resp = &ModifyAccessRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyAccessRuleResponse struct {
	responses.BaseResponse
	RequestId string
}
