package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseClusterForInternalRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	UserId          string `position:"Query" name:"UserId"`
}

func (req *ReleaseClusterForInternalRequest) Invoke(client *sdk.Client) (resp *ReleaseClusterForInternalResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ReleaseClusterForInternal", "", "")
	resp = &ReleaseClusterForInternalResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReleaseClusterForInternalResponse struct {
	responses.BaseResponse
	RequestId string
}
