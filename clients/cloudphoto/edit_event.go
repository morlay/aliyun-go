package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type EditEventRequest struct {
	requests.RpcRequest
	EventId          string `position:"Query" name:"EventId"`
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

func (req *EditEventRequest) Invoke(client *sdk.Client) (resp *EditEventResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "EditEvent", "cloudphoto", "")
	resp = &EditEventResponse{}
	err = client.DoAction(req, resp)
	return
}

type EditEventResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Event     EditEventEvent
}

type EditEventEvent struct {
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
