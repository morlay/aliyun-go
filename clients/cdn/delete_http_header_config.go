package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteHttpHeaderConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigID      string `position:"Query" name:"ConfigID"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteHttpHeaderConfigRequest) Invoke(client *sdk.Client) (resp *DeleteHttpHeaderConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DeleteHttpHeaderConfig", "", "")
	resp = &DeleteHttpHeaderConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteHttpHeaderConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
