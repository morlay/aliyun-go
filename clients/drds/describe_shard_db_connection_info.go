package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId      string
	Success        bool
	ConnectionInfo DescribeShardDbConnectionInfoConnectionInfo
}

type DescribeShardDbConnectionInfoConnectionInfo struct {
	InstanceName               string
	InstanceUrl                string
	SubDbName                  string
	DbStatus                   string
	DbType                     string
	MinPoolSize                int
	MaxPoolSize                int
	IdleTimeOut                int
	BlockingTimeout            int
	ConnectionProperties       string
	PreparedStatementCacheSize int
	UserName                   string
}
