package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceMaintainTimeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	MaintainStartTime    string `position:"Query" name:"MaintainStartTime"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MaintainEndTime      string `position:"Query" name:"MaintainEndTime"`
}

func (r ModifyDBInstanceMaintainTimeRequest) Invoke(client *sdk.Client) (response *ModifyDBInstanceMaintainTimeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDBInstanceMaintainTimeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "ModifyDBInstanceMaintainTime", "dds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDBInstanceMaintainTimeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDBInstanceMaintainTimeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDBInstanceMaintainTimeResponse struct {
	RequestId string
}
