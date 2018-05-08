
package imagesearch

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type AddItemRequest struct {
requests.RoaRequest
ItemId string `name:"ItemId"`
CatId string `name:"CatId"`
CustContent string `name:"CustContent"`
PicMap AddItemMap<String,String> `name:"PicMap"`
InstanceName string `position:"Query" name:"InstanceName"`
}

func (req *AddItemRequest) Invoke(client *sdk.Client) (resp *AddItemResponse, err error) {
	req.InitWithApiInfo("ImageSearch", "2018-01-20", "AddItem", "/item/add","", "")
    req.Method = "POST"

    resp = &AddItemResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddItemResponse struct {
responses.BaseResponse
RequestId common.String
Success bool
Message common.String
Code common.Integer
}

