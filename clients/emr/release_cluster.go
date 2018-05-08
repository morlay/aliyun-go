package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ReleaseClusterRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ForceRelease    string `position:"Query" name:"ForceRelease"`
	Id              string `position:"Query" name:"Id"`
}

func (req *ReleaseClusterRequest) Invoke(client *sdk.Client) (resp *ReleaseClusterResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ReleaseCluster", "", "")
	resp = &ReleaseClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReleaseClusterResponse struct {
	responses.BaseResponse
	RequestId common.String
}
