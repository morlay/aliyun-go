package ossadmin

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateOssInstanceRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Region               string `position:"Query" name:"Region"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
}

func (req *CreateOssInstanceRequest) Invoke(client *sdk.Client) (resp *CreateOssInstanceResponse, err error) {
	req.InitWithApiInfo("OssAdmin", "2015-03-02", "CreateOssInstance", "", "")
	resp = &CreateOssInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateOssInstanceResponse struct {
	responses.BaseResponse
	RequestId      common.String
	AliUid         common.Long
	InstanceId     common.String
	InstacneStatus common.String
	InstanceName   common.String
	StartTime      common.String
	EndTime        common.String
}
