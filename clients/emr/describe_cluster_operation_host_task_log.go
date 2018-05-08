package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeClusterOperationHostTaskLogRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	OperationId     string `position:"Query" name:"OperationId"`
	HostId          string `position:"Query" name:"HostId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	TaskId          string `position:"Query" name:"TaskId"`
	Status          string `position:"Query" name:"Status"`
}

func (req *DescribeClusterOperationHostTaskLogRequest) Invoke(client *sdk.Client) (resp *DescribeClusterOperationHostTaskLogResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterOperationHostTaskLog", "", "")
	resp = &DescribeClusterOperationHostTaskLogResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterOperationHostTaskLogResponse struct {
	responses.BaseResponse
	RequestId common.String
	Stdout    common.String
	Stderr    common.String
}
