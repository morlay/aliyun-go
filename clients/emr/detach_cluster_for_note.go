package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DetachClusterForNoteRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DetachClusterForNoteRequest) Invoke(client *sdk.Client) (resp *DetachClusterForNoteResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DetachClusterForNote", "", "")
	resp = &DetachClusterForNoteResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetachClusterForNoteResponse struct {
	responses.BaseResponse
	RequestId common.String
}
