package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRefreshQuotaRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeRefreshQuotaRequest) Invoke(client *sdk.Client) (resp *DescribeRefreshQuotaResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeRefreshQuota", "", "")
	resp = &DescribeRefreshQuotaResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRefreshQuotaResponse struct {
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
