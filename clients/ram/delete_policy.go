package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeletePolicyRequest struct {
	requests.RpcRequest
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (req *DeletePolicyRequest) Invoke(client *sdk.Client) (resp *DeletePolicyResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "DeletePolicy", "", "")
	resp = &DeletePolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeletePolicyResponse struct {
	responses.BaseResponse
	RequestId string
}
