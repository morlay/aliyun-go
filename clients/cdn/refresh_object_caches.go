package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RefreshObjectCachesRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ObjectPath    string `position:"Query" name:"ObjectPath"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	ObjectType    string `position:"Query" name:"ObjectType"`
}

func (req *RefreshObjectCachesRequest) Invoke(client *sdk.Client) (resp *RefreshObjectCachesResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "RefreshObjectCaches", "", "")
	resp = &RefreshObjectCachesResponse{}
	err = client.DoAction(req, resp)
	return
}

type RefreshObjectCachesResponse struct {
	responses.BaseResponse
	RequestId     common.String
	RefreshTaskId common.String
}
