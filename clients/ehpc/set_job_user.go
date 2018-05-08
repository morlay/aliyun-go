package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetJobUserRequest struct {
	requests.RpcRequest
	RunasUserPassword string `position:"Query" name:"RunasUserPassword"`
	RunasUser         string `position:"Query" name:"RunasUser"`
	ClusterId         string `position:"Query" name:"ClusterId"`
}

func (req *SetJobUserRequest) Invoke(client *sdk.Client) (resp *SetJobUserResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "SetJobUser", "ehs", "")
	resp = &SetJobUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetJobUserResponse struct {
	responses.BaseResponse
	RequestId common.String
}
