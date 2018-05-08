package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AttachClusterForNoteRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *AttachClusterForNoteRequest) Invoke(client *sdk.Client) (resp *AttachClusterForNoteResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "AttachClusterForNote", "", "")
	resp = &AttachClusterForNoteResponse{}
	err = client.DoAction(req, resp)
	return
}

type AttachClusterForNoteResponse struct {
	responses.BaseResponse
	RequestId common.String
}
