package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetDefaultPolicyVersionRequest struct {
	VersionId  string `position:"Query" name:"VersionId"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (r SetDefaultPolicyVersionRequest) Invoke(client *sdk.Client) (response *SetDefaultPolicyVersionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetDefaultPolicyVersionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "SetDefaultPolicyVersion", "", "")

	resp := struct {
		*responses.BaseResponse
		SetDefaultPolicyVersionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetDefaultPolicyVersionResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetDefaultPolicyVersionResponse struct {
	RequestId string
}
