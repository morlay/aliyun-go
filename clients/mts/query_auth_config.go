package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryAuthConfigRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *QueryAuthConfigRequest) Invoke(client *sdk.Client) (resp *QueryAuthConfigResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryAuthConfig", "mts", "")
	resp = &QueryAuthConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryAuthConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
	Key1      common.String
	Key2      common.String
}
