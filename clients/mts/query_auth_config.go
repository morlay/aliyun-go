package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryAuthConfigRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *QueryAuthConfigRequest) Invoke(client *sdk.Client) (resp *QueryAuthConfigResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryAuthConfig", "", "")
	resp = &QueryAuthConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryAuthConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Key1      string
	Key2      string
}
