package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterForInternalRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	UserId          string `position:"Query" name:"UserId"`
}

func (req *DescribeClusterForInternalRequest) Invoke(client *sdk.Client) (resp *DescribeClusterForInternalResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterForInternal", "", "")
	resp = &DescribeClusterForInternalResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterForInternalResponse struct {
	responses.BaseResponse
	RequestId   string
	Id          string
	BizId       string
	Name        string
	StartTime   int64
	StopTime    int64
	LogEnable   bool
	LogPath     string
	UserId      string
	RunningTime int
	Status      string
	ExpiredTime int64
	FailReason  DescribeClusterForInternalFailReason
}

type DescribeClusterForInternalFailReason struct {
	ErrorCode string
	ErrorMsg  string
	RequestId string
}
