package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	Name            string `position:"Query" name:"Name"`
	Type            string `position:"Query" name:"Type"`
	RunParameter    string `position:"Query" name:"RunParameter"`
	FailAct         string `position:"Query" name:"FailAct"`
}

func (req *ModifyJobRequest) Invoke(client *sdk.Client) (resp *ModifyJobResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyJob", "", "")
	resp = &ModifyJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyJobResponse struct {
	responses.BaseResponse
	RequestId string
}
