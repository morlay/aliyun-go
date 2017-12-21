package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RefreshObjectCachesRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ObjectPath    string `position:"Query" name:"ObjectPath"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	ObjectType    string `position:"Query" name:"ObjectType"`
}

func (r RefreshObjectCachesRequest) Invoke(client *sdk.Client) (response *RefreshObjectCachesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RefreshObjectCachesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "RefreshObjectCaches", "", "")

	resp := struct {
		*responses.BaseResponse
		RefreshObjectCachesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RefreshObjectCachesResponse

	err = client.DoAction(&req, &resp)
	return
}

type RefreshObjectCachesResponse struct {
	RequestId     string
	RefreshTaskId string
}
