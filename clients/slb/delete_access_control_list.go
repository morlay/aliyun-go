package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAccessControlListRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	AclId                string `position:"Query" name:"AclId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DeleteAccessControlListRequest) Invoke(client *sdk.Client) (resp *DeleteAccessControlListResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DeleteAccessControlList", "slb", "")
	resp = &DeleteAccessControlListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteAccessControlListResponse struct {
	responses.BaseResponse
	RequestId string
}
