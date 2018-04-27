package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRefreshQuotaRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *DescribeRefreshQuotaRequest) Invoke(client *sdk.Client) (resp *DescribeRefreshQuotaResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "DescribeRefreshQuota", "vod", "")
	resp = &DescribeRefreshQuotaResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRefreshQuotaResponse struct {
	responses.BaseResponse
	RequestId         string
	RefreshCacheQuota DescribeRefreshQuotaRefreshCacheQuota
}

type DescribeRefreshQuotaRefreshCacheQuota struct {
	UrlQuota  string
	DirQuota  string
	UrlRemain string
	DirRemain string
}
