package dcdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDcdnRefreshQuotaRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDcdnRefreshQuotaRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnRefreshQuotaResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnRefreshQuota", "dcdn", "")
	resp = &DescribeDcdnRefreshQuotaResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnRefreshQuotaResponse struct {
	responses.BaseResponse
	RequestId     string
	UrlQuota      string
	DirQuota      string
	UrlRemain     string
	DirRemain     string
	PreloadQuota  string
	BlockQuota    string
	PreloadRemain string
	BlockRemain   string
}
