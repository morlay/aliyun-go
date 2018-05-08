package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteCacheExpiredConfigRequest struct {
	requests.RpcRequest
	CacheType     string `position:"Query" name:"CacheType"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigID      string `position:"Query" name:"ConfigID"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteCacheExpiredConfigRequest) Invoke(client *sdk.Client) (resp *DeleteCacheExpiredConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DeleteCacheExpiredConfig", "", "")
	resp = &DeleteCacheExpiredConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCacheExpiredConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
