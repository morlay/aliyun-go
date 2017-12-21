package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseClusterRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ForceRelease    string `position:"Query" name:"ForceRelease"`
}

func (req *ReleaseClusterRequest) Invoke(client *sdk.Client) (resp *ReleaseClusterResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ReleaseCluster", "", "")
	resp = &ReleaseClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReleaseClusterResponse struct {
	responses.BaseResponse
	RequestId string
}
