package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveMixNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLiveMixNotifyConfigRequest) Invoke(client *sdk.Client) (response *DescribeLiveMixNotifyConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveMixNotifyConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveMixNotifyConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveMixNotifyConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveMixNotifyConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveMixNotifyConfigResponse struct {
	RequestId string
	NotifyUrl string
}
