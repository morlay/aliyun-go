package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	Id           common.String
	GmtCreate    common.Long
	GmtModified  common.Long
	Name         common.String
	ParentId     common.String
	Type         common.String
	CategoryType common.String
	ObjectType   common.String
	ObjectId     common.String
	ProjectId    common.String
}
