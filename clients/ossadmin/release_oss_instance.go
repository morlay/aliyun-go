package ossadmin

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseOssInstanceRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Region               string `position:"Query" name:"Region"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
}

func (req *ReleaseOssInstanceRequest) Invoke(client *sdk.Client) (resp *ReleaseOssInstanceResponse, err error) {
	req.InitWithApiInfo("OssAdmin", "2015-03-02", "ReleaseOssInstance", "", "")
	resp = &ReleaseOssInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReleaseOssInstanceResponse struct {
	responses.BaseResponse
	RequestId string
}
