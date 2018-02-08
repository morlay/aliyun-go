package afs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeConfigNameRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
}

func (req *DescribeConfigNameRequest) Invoke(client *sdk.Client) (resp *DescribeConfigNameResponse, err error) {
	req.InitWithApiInfo("afs", "2018-01-12", "DescribeConfigName", "", "")
	resp = &DescribeConfigNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeConfigNameResponse struct {
	responses.BaseResponse
	RequestId   string
	HasConfig   bool
	ConfigNames string
	BizCode     string
}
