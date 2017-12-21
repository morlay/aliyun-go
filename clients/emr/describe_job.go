package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeJobRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (r DescribeJobRequest) Invoke(client *sdk.Client) (response *DescribeJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeJob", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeJobResponse struct {
	RequestId    string
	Id           string
	Name         string
	FailAct      string
	Type         string
	RunParameter string
}
