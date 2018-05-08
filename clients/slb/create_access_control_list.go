package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateAccessControlListRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AclName              string `position:"Query" name:"AclName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *CreateAccessControlListRequest) Invoke(client *sdk.Client) (resp *CreateAccessControlListResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "CreateAccessControlList", "slb", "")
	resp = &CreateAccessControlListResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateAccessControlListResponse struct {
	responses.BaseResponse
	RequestId common.String
	AclId     common.String
}
