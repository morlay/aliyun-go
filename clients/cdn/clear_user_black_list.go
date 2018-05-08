package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ClearUserBlackListRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *ClearUserBlackListRequest) Invoke(client *sdk.Client) (resp *ClearUserBlackListResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "ClearUserBlackList", "", "")
	resp = &ClearUserBlackListResponse{}
	err = client.DoAction(req, resp)
	return
}

type ClearUserBlackListResponse struct {
	responses.BaseResponse
	RequestId common.String
}
