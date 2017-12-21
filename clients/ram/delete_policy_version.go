package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeletePolicyVersionRequest struct {
	VersionId  string `position:"Query" name:"VersionId"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (r DeletePolicyVersionRequest) Invoke(client *sdk.Client) (response *DeletePolicyVersionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeletePolicyVersionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "DeletePolicyVersion", "", "")

	resp := struct {
		*responses.BaseResponse
		DeletePolicyVersionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeletePolicyVersionResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeletePolicyVersionResponse struct {
	RequestId string
}
