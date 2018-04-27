package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type MetastoreDropTableRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
	TableName       string `position:"Query" name:"TableName"`
}

func (req *MetastoreDropTableRequest) Invoke(client *sdk.Client) (resp *MetastoreDropTableResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "MetastoreDropTable", "", "")
	resp = &MetastoreDropTableResponse{}
	err = client.DoAction(req, resp)
	return
}

type MetastoreDropTableResponse struct {
	responses.BaseResponse
	RequestId string
}
