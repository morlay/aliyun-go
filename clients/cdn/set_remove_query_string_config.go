package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetRemoveQueryStringConfigRequest struct {
	requests.RpcRequest
	KeepOssArgs   string `position:"Query" name:"KeepOssArgs"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	AliRemoveArgs string `position:"Query" name:"AliRemoveArgs"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *SetRemoveQueryStringConfigRequest) Invoke(client *sdk.Client) (resp *SetRemoveQueryStringConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetRemoveQueryStringConfig", "", "")
	resp = &SetRemoveQueryStringConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetRemoveQueryStringConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
