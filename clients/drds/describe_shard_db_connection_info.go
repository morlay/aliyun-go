package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeShardDbConnectionInfoRequest struct {
	requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	SubDbName      string `position:"Query" name:"SubDbName"`
}

func (req *DescribeShardDbConnectionInfoRequest) Invoke(client *sdk.Client) (resp *DescribeShardDbConnectionInfoResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeShardDbConnectionInfo", "", "")
	resp = &DescribeShardDbConnectionInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeShardDbConnectionInfoResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	ConnectionInfo DescribeShardDbConnectionInfoConnectionInfo
}

type DescribeShardDbConnectionInfoConnectionInfo struct {
	InstanceName               common.String
	InstanceUrl                common.String
	SubDbName                  common.String
	DbStatus                   common.String
	DbType                     common.String
	MinPoolSize                common.Integer
	MaxPoolSize                common.Integer
	IdleTimeOut                common.Integer
	BlockingTimeout            common.Integer
	ConnectionProperties       common.String
	PreparedStatementCacheSize common.Integer
	UserName                   common.String
}
