package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterOperationHostTaskLogForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	OperationId     string `position:"Query" name:"OperationId"`
	HostId          string `position:"Query" name:"HostId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	UserId          string `position:"Query" name:"UserId"`
	TaskId          string `position:"Query" name:"TaskId"`
	Status          string `position:"Query" name:"Status"`
}

func (req *DescribeClusterOperationHostTaskLogForAdminRequest) Invoke(client *sdk.Client) (resp *DescribeClusterOperationHostTaskLogForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterOperationHostTaskLogForAdmin", "", "")
	resp = &DescribeClusterOperationHostTaskLogForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterOperationHostTaskLogForAdminResponse struct {
	responses.BaseResponse
	RequestId string
	Stdout    string
	Stderr    string
}
