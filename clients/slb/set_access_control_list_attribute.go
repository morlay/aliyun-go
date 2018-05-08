package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetAccessControlListAttributeRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	AclId                string `position:"Query" name:"AclId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AclName              string `position:"Query" name:"AclName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *SetAccessControlListAttributeRequest) Invoke(client *sdk.Client) (resp *SetAccessControlListAttributeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "SetAccessControlListAttribute", "slb", "")
	resp = &SetAccessControlListAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetAccessControlListAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
	AclId     common.String
}
