package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCustinsKernelReleaseNotesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeCustinsKernelReleaseNotesRequest) Invoke(client *sdk.Client) (resp *DescribeCustinsKernelReleaseNotesResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeCustinsKernelReleaseNotes", "rds", "")
	resp = &DescribeCustinsKernelReleaseNotesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCustinsKernelReleaseNotesResponse struct {
	responses.BaseResponse
	RequestId                 common.String
	DBInstanceId              common.String
	DBInstanceDiffReleaseNote common.String
}
