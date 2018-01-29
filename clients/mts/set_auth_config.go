package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetAuthConfigRequest struct {
	requests.RpcRequest
	Key1                 string `position:"Query" name:"Key.1"`
	Key2                 string `position:"Query" name:"Key.2"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *SetAuthConfigRequest) Invoke(client *sdk.Client) (resp *SetAuthConfigResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SetAuthConfig", "mts", "")
	resp = &SetAuthConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetAuthConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Key1      string
	Key2      string
}
