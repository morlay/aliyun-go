package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type MetastoreDataPreviewRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
	TableName       string `position:"Query" name:"TableName"`
}

func (req *MetastoreDataPreviewRequest) Invoke(client *sdk.Client) (resp *MetastoreDataPreviewResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "MetastoreDataPreview", "", "")
	resp = &MetastoreDataPreviewResponse{}
	err = client.DoAction(req, resp)
	return
}

type MetastoreDataPreviewResponse struct {
	responses.BaseResponse
	RequestId common.String
	Samples   common.String
}
