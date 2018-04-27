package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveAccessControlListEntryRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	AclId                string `position:"Query" name:"AclId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AclEntrys            string `position:"Query" name:"AclEntrys"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *RemoveAccessControlListEntryRequest) Invoke(client *sdk.Client) (resp *RemoveAccessControlListEntryResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "RemoveAccessControlListEntry", "slb", "")
	resp = &RemoveAccessControlListEntryResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveAccessControlListEntryResponse struct {
	responses.BaseResponse
	RequestId string
}
