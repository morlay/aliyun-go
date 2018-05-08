package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type MetastoreDropDatabaseRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
}

func (req *MetastoreDropDatabaseRequest) Invoke(client *sdk.Client) (resp *MetastoreDropDatabaseResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "MetastoreDropDatabase", "", "")
	resp = &MetastoreDropDatabaseResponse{}
	err = client.DoAction(req, resp)
	return
}

type MetastoreDropDatabaseResponse struct {
	responses.BaseResponse
	RequestId common.String
}
