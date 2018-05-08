package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListEventsRequest struct {
	requests.RpcRequest
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int    `position:"Query" name:"Size"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
}

func (req *ListEventsRequest) Invoke(client *sdk.Client) (resp *ListEventsResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListEvents", "cloudphoto", "")
	resp = &ListEventsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListEventsResponse struct {
	responses.BaseResponse
	Code       common.String
	Message    common.String
	NextCursor common.String
	TotalCount common.Integer
	RequestId  common.String
	Action     common.String
	Events     ListEventsEventList
}

type ListEventsEvent struct {
	Id               common.Long
	IdStr            common.String
	Title            common.String
	BannerPhotoId    common.String
	Identity         common.String
	SplashPhotoId    common.String
	State            common.String
	WeixinTitle      common.String
	WatermarkPhotoId common.String
	StartAt          common.Long
	EndAt            common.Long
	Ctime            common.Long
	Mtime            common.Long
	ViewsCount       common.Long
	LibraryId        common.String
	IdStr1           common.String
}

type ListEventsEventList []ListEventsEvent

func (list *ListEventsEventList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEventsEvent)
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
