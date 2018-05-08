package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AuthRealNameRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
}

func (req *AuthRealNameRequest) Invoke(client *sdk.Client) (resp *AuthRealNameResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "AuthRealName", "", "")
	resp = &AuthRealNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type AuthRealNameResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
}
