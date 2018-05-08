package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type MetastoreCreateDatabaseRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
	Description     string `position:"Query" name:"Description"`
	LocationUri     string `position:"Query" name:"LocationUri"`
}

func (req *MetastoreCreateDatabaseRequest) Invoke(client *sdk.Client) (resp *MetastoreCreateDatabaseResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "MetastoreCreateDatabase", "", "")
	resp = &MetastoreCreateDatabaseResponse{}
	err = client.DoAction(req, resp)
	return
}

type MetastoreCreateDatabaseResponse struct {
	responses.BaseResponse
	RequestId common.String
}
