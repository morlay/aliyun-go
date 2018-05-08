package imagesearch

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteItemRequest struct {
	requests.RoaRequest
	ItemId       string                     `name:"ItemId"`
	PictureList  *DeleteItemPictureListList `type:"Repeated" name:"PictureList"`
	InstanceName string                     `position:"Query" name:"InstanceName"`
}

func (req *DeleteItemRequest) Invoke(client *sdk.Client) (resp *DeleteItemResponse, err error) {
	req.InitWithApiInfo("ImageSearch", "2018-01-20", "DeleteItem", "/item/delete", "", "")
	req.Method = "POST"

	resp = &DeleteItemResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteItemResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Code      common.Integer
}

type DeleteItemPictureListList []string

func (list *DeleteItemPictureListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
