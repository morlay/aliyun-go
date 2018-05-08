package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type MigrateInstanceForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *MigrateInstanceForAdminRequest) Invoke(client *sdk.Client) (resp *MigrateInstanceForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "MigrateInstanceForAdmin", "", "")
	resp = &MigrateInstanceForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type MigrateInstanceForAdminResponse struct {
	responses.BaseResponse
	RequestId common.String
	WfId      common.String
}
