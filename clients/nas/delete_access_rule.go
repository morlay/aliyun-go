package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAccessRuleRequest struct {
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	AccessRuleId    string `position:"Query" name:"AccessRuleId"`
}

func (r DeleteAccessRuleRequest) Invoke(client *sdk.Client) (response *DeleteAccessRuleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteAccessRuleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("NAS", "2017-06-26", "DeleteAccessRule", "nas", "")

	resp := struct {
		*responses.BaseResponse
		DeleteAccessRuleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteAccessRuleResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteAccessRuleResponse struct {
	RequestId string
}
