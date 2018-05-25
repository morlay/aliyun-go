package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListTagPhotosRequest struct {
	requests.RpcRequest
	TagName string `position:"Query" name:"TagName"`
	MaxKeys string `position:"Query" name:"MaxKeys"`
	Marker  string `position:"Query" name:"Marker"`
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
}

func (req *ListTagPhotosRequest) Invoke(client *sdk.Client) (resp *ListTagPhotosResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "ListTagPhotos", "imm", "")
	resp = &ListTagPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListTagPhotosResponse struct {
	responses.BaseResponse
	RequestId  common.String
	NextMarker common.String
	Photos     ListTagPhotosPhotosItemList
}

type ListTagPhotosPhotosItem struct {
	SrcUri   common.String
	TagScore common.Float
}

type ListTagPhotosPhotosItemList []ListTagPhotosPhotosItem

func (list *ListTagPhotosPhotosItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagPhotosPhotosItem)
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
