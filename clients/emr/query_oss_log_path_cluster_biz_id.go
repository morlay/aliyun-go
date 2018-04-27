package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryOssLogPathClusterBizIdRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *QueryOssLogPathClusterBizIdRequest) Invoke(client *sdk.Client) (resp *QueryOssLogPathClusterBizIdResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "QueryOssLogPathClusterBizId", "", "")
	resp = &QueryOssLogPathClusterBizIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryOssLogPathClusterBizIdResponse struct {
	responses.BaseResponse
	RequestId              string
	OssLogPathClusterBizId string
}
