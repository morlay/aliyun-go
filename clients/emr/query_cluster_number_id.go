package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryClusterNumberIdRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *QueryClusterNumberIdRequest) Invoke(client *sdk.Client) (resp *QueryClusterNumberIdResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "QueryClusterNumberId", "", "")
	resp = &QueryClusterNumberIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryClusterNumberIdResponse struct {
	responses.BaseResponse
	RequestId       string
	ClusterNumberId string
}
