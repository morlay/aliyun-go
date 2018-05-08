package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetFileCacheExpiredConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Weight        string `position:"Query" name:"Weight"`
	CacheContent  string `position:"Query" name:"CacheContent"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	TTL           string `position:"Query" name:"TTL"`
}

func (req *SetFileCacheExpiredConfigRequest) Invoke(client *sdk.Client) (resp *SetFileCacheExpiredConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetFileCacheExpiredConfig", "", "")
	resp = &SetFileCacheExpiredConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetFileCacheExpiredConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
