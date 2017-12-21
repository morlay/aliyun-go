package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeletePolicyRequest struct {
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (r DeletePolicyRequest) Invoke(client *sdk.Client) (response *DeletePolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeletePolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "DeletePolicy", "", "")

	resp := struct {
		*responses.BaseResponse
		DeletePolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeletePolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeletePolicyResponse struct {
	RequestId string
}
