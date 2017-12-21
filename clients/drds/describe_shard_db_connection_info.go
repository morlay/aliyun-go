package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeShardDbConnectionInfoRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	SubDbName      string `position:"Query" name:"SubDbName"`
}

func (r DescribeShardDbConnectionInfoRequest) Invoke(client *sdk.Client) (response *DescribeShardDbConnectionInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeShardDbConnectionInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeShardDbConnectionInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeShardDbConnectionInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeShardDbConnectionInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeShardDbConnectionInfoResponse struct {
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
