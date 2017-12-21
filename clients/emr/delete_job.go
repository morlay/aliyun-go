package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DeleteJobRequest) Invoke(client *sdk.Client) (resp *DeleteJobResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteJob", "", "")
	resp = &DeleteJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteJobResponse struct {
	responses.BaseResponse
	RequestId string
}
