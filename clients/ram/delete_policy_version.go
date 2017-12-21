package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeletePolicyVersionRequest struct {
	requests.RpcRequest
	VersionId  string `position:"Query" name:"VersionId"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (req *DeletePolicyVersionRequest) Invoke(client *sdk.Client) (resp *DeletePolicyVersionResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "DeletePolicyVersion", "", "")
	resp = &DeletePolicyVersionResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeletePolicyVersionResponse struct {
	responses.BaseResponse
	RequestId string
}
