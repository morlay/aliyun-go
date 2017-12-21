package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeOneMinuteDataRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DataTime             string `position:"Query" name:"DataTime"`
	DomainName           string `position:"Query" name:"DomainName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeOneMinuteDataRequest) Invoke(client *sdk.Client) (resp *DescribeOneMinuteDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeOneMinuteData", "", "")
	resp = &DescribeOneMinuteDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeOneMinuteDataResponse struct {
	responses.BaseResponse
	RequestId string
	Bps       string
	Qps       string
}
