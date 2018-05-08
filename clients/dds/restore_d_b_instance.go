package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RestoreDBInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BackupId             int    `position:"Query" name:"BackupId"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *RestoreDBInstanceRequest) Invoke(client *sdk.Client) (resp *RestoreDBInstanceResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "RestoreDBInstance", "dds", "")
	resp = &RestoreDBInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RestoreDBInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
