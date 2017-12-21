package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRangeDataByLocateAndIspServiceRequest struct {
	IspNames      string `position:"Query" name:"IspNames"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainNames   string `position:"Query" name:"DomainNames"`
	LocationNames string `position:"Query" name:"LocationNames"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeRangeDataByLocateAndIspServiceRequest) Invoke(client *sdk.Client) (response *DescribeRangeDataByLocateAndIspServiceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRangeDataByLocateAndIspServiceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeRangeDataByLocateAndIspService", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeRangeDataByLocateAndIspServiceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeRangeDataByLocateAndIspServiceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRangeDataByLocateAndIspServiceResponse struct {
	RequestId  string
	JsonResult string
}
