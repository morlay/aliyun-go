package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetUserBlackListRequest struct {
	requests.RpcRequest
	ConfigUrl     string `position:"Query" name:"ConfigUrl"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *SetUserBlackListRequest) Invoke(client *sdk.Client) (resp *SetUserBlackListResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetUserBlackList", "", "")
	resp = &SetUserBlackListResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetUserBlackListResponse struct {
	responses.BaseResponse
	RequestId common.String
}
