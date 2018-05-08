package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DescribeJobRequest) Invoke(client *sdk.Client) (resp *DescribeJobResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeJob", "", "")
	resp = &DescribeJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeJobResponse struct {
	responses.BaseResponse
	RequestId     common.String
	Id            common.String
	Name          common.String
	FailAct       common.String
	Type          common.String
	MaxRetry      common.Integer
	RetryInterval common.Integer
	RunParameter  common.String
}
