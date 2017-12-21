package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRefreshQuotaRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeRefreshQuotaRequest) Invoke(client *sdk.Client) (response *DescribeRefreshQuotaResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRefreshQuotaRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeRefreshQuota", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeRefreshQuotaResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeRefreshQuotaResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRefreshQuotaResponse struct {
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
