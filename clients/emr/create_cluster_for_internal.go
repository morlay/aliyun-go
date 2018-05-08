package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateClusterForInternalRequest struct {
	requests.RpcRequest
	ResourceOwnerId   int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterTemplateId int64  `position:"Query" name:"ClusterTemplateId"`
	ClusterCondition  string `position:"Query" name:"ClusterCondition"`
	UserId            string `position:"Query" name:"UserId"`
}

func (req *CreateClusterForInternalRequest) Invoke(client *sdk.Client) (resp *CreateClusterForInternalResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateClusterForInternal", "", "")
	resp = &CreateClusterForInternalResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateClusterForInternalResponse struct {
	responses.BaseResponse
	RequestId common.String
	ClusterId common.String
}
