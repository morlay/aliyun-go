package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeEmrStackVersionRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	MainVersion     string `position:"Query" name:"MainVersion"`
}

func (req *DescribeEmrStackVersionRequest) Invoke(client *sdk.Client) (resp *DescribeEmrStackVersionResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeEmrStackVersion", "", "")
	resp = &DescribeEmrStackVersionResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeEmrStackVersionResponse struct {
	responses.BaseResponse
	RequestId    string
	StackVersion string
}
