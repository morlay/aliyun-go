package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLogsDownloadAttributeRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RoleName             string `position:"Query" name:"RoleName"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DeleteLogsDownloadAttributeRequest) Invoke(client *sdk.Client) (resp *DeleteLogsDownloadAttributeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DeleteLogsDownloadAttribute", "slb", "")
	resp = &DeleteLogsDownloadAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteLogsDownloadAttributeResponse struct {
	responses.BaseResponse
	RequestId string
}
