package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Events     ListEventsEventList
}

type ListEventsEvent struct {
	Id               int64
	IdStr            string
	Title            string
	BannerPhotoId    string
	Identity         string
	SplashPhotoId    string
	State            string
	WeixinTitle      string
	WatermarkPhotoId string
	StartAt          int64
	EndAt            int64
	Ctime            int64
	Mtime            int64
	ViewsCount       int64
	LibraryId        string
	IdStr1           string
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
