package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateAccessRuleRequest struct {
	RWAccessType    string `position:"Query" name:"RWAccessType"`
	SourceCidrIp    string `position:"Query" name:"SourceCidrIp"`
	UserAccessType  string `position:"Query" name:"UserAccessType"`
	Priority        int    `position:"Query" name:"Priority"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
}

func (r CreateAccessRuleRequest) Invoke(client *sdk.Client) (response *CreateAccessRuleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateAccessRuleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("NAS", "2017-06-26", "CreateAccessRule", "nas", "")

	resp := struct {
		*responses.BaseResponse
		CreateAccessRuleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateAccessRuleResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateAccessRuleResponse struct {
	RequestId    string
	AccessRuleId string
}
