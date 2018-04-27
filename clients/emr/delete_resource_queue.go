package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteResourceQueueRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceQueueId string `position:"Query" name:"ResourceQueueId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *DeleteResourceQueueRequest) Invoke(client *sdk.Client) (resp *DeleteResourceQueueResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteResourceQueue", "", "")
	resp = &DeleteResourceQueueResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteResourceQueueResponse struct {
	responses.BaseResponse
	RequestId string
}
