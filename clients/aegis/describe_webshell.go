package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeWebshellRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	GroupId         int    `position:"Query" name:"GroupId"`
	Remark          string `position:"Query" name:"Remark"`
}

func (req *DescribeWebshellRequest) Invoke(client *sdk.Client) (resp *DescribeWebshellResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "DescribeWebshell", "vipaegis", "")
	resp = &DescribeWebshellResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeWebshellResponse struct {
	responses.BaseResponse
	RequestId   common.String
	TotalCount  common.Integer
	PageSize    common.Integer
	CurrentPage common.Integer
}
