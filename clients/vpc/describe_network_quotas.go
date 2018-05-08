package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeNetworkQuotasRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Product              string `position:"Query" name:"Product"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeNetworkQuotasRequest) Invoke(client *sdk.Client) (resp *DescribeNetworkQuotasResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeNetworkQuotas", "vpc", "")
	resp = &DescribeNetworkQuotasResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeNetworkQuotasResponse struct {
	responses.BaseResponse
	RequestId common.String
	Product   common.String
	RegionId  common.String
	Quota     common.String
}
