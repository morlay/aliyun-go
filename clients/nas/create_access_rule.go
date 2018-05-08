package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateAccessRuleRequest struct {
	requests.RpcRequest
	RWAccessType    string `position:"Query" name:"RWAccessType"`
	SourceCidrIp    string `position:"Query" name:"SourceCidrIp"`
	UserAccessType  string `position:"Query" name:"UserAccessType"`
	Priority        int    `position:"Query" name:"Priority"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
}

func (req *CreateAccessRuleRequest) Invoke(client *sdk.Client) (resp *CreateAccessRuleResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "CreateAccessRule", "nas", "")
	resp = &CreateAccessRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateAccessRuleResponse struct {
	responses.BaseResponse
	RequestId    common.String
	AccessRuleId common.String
}
