package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeFlowCategoryTreeRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Type            string `position:"Query" name:"Type"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DescribeFlowCategoryTreeRequest) Invoke(client *sdk.Client) (resp *DescribeFlowCategoryTreeResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlowCategoryTree", "", "")
	resp = &DescribeFlowCategoryTreeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeFlowCategoryTreeResponse struct {
	responses.BaseResponse
	RequestId string
	Data      string
}
