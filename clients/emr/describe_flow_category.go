package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeFlowCategoryRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DescribeFlowCategoryRequest) Invoke(client *sdk.Client) (resp *DescribeFlowCategoryResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlowCategory", "", "")
	resp = &DescribeFlowCategoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeFlowCategoryResponse struct {
	responses.BaseResponse
	RequestId    string
	Id           string
	GmtCreate    int64
	GmtModified  int64
	Name         string
	ParentId     string
	Type         string
	CategoryType string
	ObjectType   string
	ObjectId     string
	ProjectId    string
}
