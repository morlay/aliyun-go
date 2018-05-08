package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PushObjectCacheRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ObjectPath    string `position:"Query" name:"ObjectPath"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *PushObjectCacheRequest) Invoke(client *sdk.Client) (resp *PushObjectCacheResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "PushObjectCache", "", "")
	resp = &PushObjectCacheResponse{}
	err = client.DoAction(req, resp)
	return
}

type PushObjectCacheResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PushTaskId common.String
}
