package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	UrlQuota      common.String
	DirQuota      common.String
	UrlRemain     common.String
	DirRemain     common.String
	PreloadQuota  common.String
	BlockQuota    common.String
	PreloadRemain common.String
	BlockRemain   common.String
}
