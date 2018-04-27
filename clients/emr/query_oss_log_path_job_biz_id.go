package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryOssLogPathJobBizIdRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	WorkNodeExecId  string `position:"Query" name:"WorkNodeExecId"`
}

func (req *QueryOssLogPathJobBizIdRequest) Invoke(client *sdk.Client) (resp *QueryOssLogPathJobBizIdResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "QueryOssLogPathJobBizId", "", "")
	resp = &QueryOssLogPathJobBizIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryOssLogPathJobBizIdResponse struct {
	responses.BaseResponse
	RequestId          string
	OssLogPathJobBizId string
}
