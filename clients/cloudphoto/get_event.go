package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Event     GetEventEvent
}

type GetEventEvent struct {
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
