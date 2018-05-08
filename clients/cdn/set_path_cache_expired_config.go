package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetPathCacheExpiredConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Weight        string `position:"Query" name:"Weight"`
	CacheContent  string `position:"Query" name:"CacheContent"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	TTL           string `position:"Query" name:"TTL"`
}

func (req *SetPathCacheExpiredConfigRequest) Invoke(client *sdk.Client) (resp *SetPathCacheExpiredConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetPathCacheExpiredConfig", "", "")
	resp = &SetPathCacheExpiredConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetPathCacheExpiredConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
