package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetEventRequest struct {
	requests.RpcRequest
	EventId   int64  `position:"Query" name:"EventId"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (req *GetEventRequest) Invoke(client *sdk.Client) (resp *GetEventResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetEvent", "cloudphoto", "")
	resp = &GetEventResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetEventResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
	Event     GetEventEvent
}

type GetEventEvent struct {
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
