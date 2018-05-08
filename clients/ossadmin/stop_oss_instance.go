package ossadmin

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type StopOssInstanceRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Region               string `position:"Query" name:"Region"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
}

func (req *StopOssInstanceRequest) Invoke(client *sdk.Client) (resp *StopOssInstanceResponse, err error) {
	req.InitWithApiInfo("OssAdmin", "2015-03-02", "StopOssInstance", "", "")
	resp = &StopOssInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopOssInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
