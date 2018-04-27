package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyEmrAlarmStatusForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	UniqueKey       string `position:"Query" name:"UniqueKey"`
	Status          string `position:"Query" name:"Status"`
}

func (req *ModifyEmrAlarmStatusForAdminRequest) Invoke(client *sdk.Client) (resp *ModifyEmrAlarmStatusForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyEmrAlarmStatusForAdmin", "", "")
	resp = &ModifyEmrAlarmStatusForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyEmrAlarmStatusForAdminResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
}
