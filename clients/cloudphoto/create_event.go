package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateEventRequest struct {
	requests.RpcRequest
	BannerPhotoId    string `position:"Query" name:"BannerPhotoId"`
	WatermarkPhotoId string `position:"Query" name:"WatermarkPhotoId"`
	Identity         string `position:"Query" name:"Identity"`
	SplashPhotoId    string `position:"Query" name:"SplashPhotoId"`
	LibraryId        string `position:"Query" name:"LibraryId"`
	WeixinTitle      string `position:"Query" name:"WeixinTitle"`
	StoreName        string `position:"Query" name:"StoreName"`
	Remark           string `position:"Query" name:"Remark"`
	Title            string `position:"Query" name:"Title"`
	EndAt            int64  `position:"Query" name:"EndAt"`
	StartAt          int64  `position:"Query" name:"StartAt"`
}

func (req *CreateEventRequest) Invoke(client *sdk.Client) (resp *CreateEventResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "CreateEvent", "cloudphoto", "")
	resp = &CreateEventResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateEventResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
	Event     CreateEventEvent
}

type CreateEventEvent struct {
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
