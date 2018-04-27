package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyFileCacheExpiredConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigID      string `position:"Query" name:"ConfigID"`
	DomainName    string `position:"Query" name:"DomainName"`
	Weight        string `position:"Query" name:"Weight"`
	CacheContent  string `position:"Query" name:"CacheContent"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	TTL           string `position:"Query" name:"TTL"`
}

func (req *ModifyFileCacheExpiredConfigRequest) Invoke(client *sdk.Client) (resp *ModifyFileCacheExpiredConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "ModifyFileCacheExpiredConfig", "", "")
	resp = &ModifyFileCacheExpiredConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyFileCacheExpiredConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
