package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveMixNotifyConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveMixNotifyConfigRequest) Invoke(client *sdk.Client) (resp *DescribeLiveMixNotifyConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveMixNotifyConfig", "live", "")
	resp = &DescribeLiveMixNotifyConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveMixNotifyConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
	NotifyUrl common.String
}
