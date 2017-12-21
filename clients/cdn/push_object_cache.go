package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PushObjectCacheRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ObjectPath    string `position:"Query" name:"ObjectPath"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r PushObjectCacheRequest) Invoke(client *sdk.Client) (response *PushObjectCacheResponse, err error) {
	req := struct {
		*requests.RpcRequest
		PushObjectCacheRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "PushObjectCache", "", "")

	resp := struct {
		*responses.BaseResponse
		PushObjectCacheResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PushObjectCacheResponse

	err = client.DoAction(&req, &resp)
	return
}

type PushObjectCacheResponse struct {
	RequestId  string
	PushTaskId string
}
