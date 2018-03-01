package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetTagPhotosRequest struct {
	requests.RpcRequest
	TagName string `position:"Query" name:"TagName"`
	MaxKeys string `position:"Query" name:"MaxKeys"`
	Marker  string `position:"Query" name:"Marker"`
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
}

func (req *GetTagPhotosRequest) Invoke(client *sdk.Client) (resp *GetTagPhotosResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "GetTagPhotos", "imm", "")
	resp = &GetTagPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetTagPhotosResponse struct {
	responses.BaseResponse
	RequestId  string
	NextMarker string
	Photos     GetTagPhotosPhotosItemList
}

type GetTagPhotosPhotosItem struct {
	SrcUri   string
	TagScore float32
}

type GetTagPhotosPhotosItemList []GetTagPhotosPhotosItem

func (list *GetTagPhotosPhotosItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetTagPhotosPhotosItem)
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
