package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId   common.String
	Id          common.String
	BizId       common.String
	Name        common.String
	StartTime   common.Long
	StopTime    common.Long
	LogEnable   bool
	LogPath     common.String
	UserId      common.String
	RunningTime common.Integer
	Status      common.String
	ExpiredTime common.Long
	FailReason  DescribeClusterForInternalFailReason
}

type DescribeClusterForInternalFailReason struct {
	ErrorCode common.String
	ErrorMsg  common.String
	RequestId common.String
}
