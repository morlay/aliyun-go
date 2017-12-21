package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeOneMinuteDataRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DataTime             string `position:"Query" name:"DataTime"`
	DomainName           string `position:"Query" name:"DomainName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeOneMinuteDataRequest) Invoke(client *sdk.Client) (response *DescribeOneMinuteDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeOneMinuteDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeOneMinuteData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeOneMinuteDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeOneMinuteDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeOneMinuteDataResponse struct {
	RequestId string
	Bps       string
	Qps       string
}
