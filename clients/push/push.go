package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PushRequest struct {
	requests.RpcRequest
	AndroidNotificationBarType     int    `position:"Query" name:"AndroidNotificationBarType"`
	SmsSendPolicy                  int    `position:"Query" name:"SmsSendPolicy"`
	AndroidExtParameters           string `position:"Query" name:"AndroidExtParameters"`
	IOSBadge                       int    `position:"Query" name:"IOSBadge"`
	IOSBadgeAutoIncrement          string `position:"Query" name:"IOSBadgeAutoIncrement"`
	AndroidOpenType                string `position:"Query" name:"AndroidOpenType"`
	Title                          string `position:"Query" name:"Title"`
	Body                           string `position:"Query" name:"Body"`
	DeviceType                     string `position:"Query" name:"DeviceType"`
	PushTime                       string `position:"Query" name:"PushTime"`
	SmsDelaySecs                   int    `position:"Query" name:"SmsDelaySecs"`
	SendSpeed                      int    `position:"Query" name:"SendSpeed"`
	AndroidPopupActivity           string `position:"Query" name:"AndroidPopupActivity"`
	IOSRemindBody                  string `position:"Query" name:"IOSRemindBody"`
	BatchNumber                    string `position:"Query" name:"BatchNumber"`
	IOSExtParameters               string `position:"Query" name:"IOSExtParameters"`
	AndroidNotifyType              string `position:"Query" name:"AndroidNotifyType"`
	AndroidPopupTitle              string `position:"Query" name:"AndroidPopupTitle"`
	IOSMusic                       string `position:"Query" name:"IOSMusic"`
	IOSApnsEnv                     string `position:"Query" name:"IOSApnsEnv"`
	IOSMutableContent              string `position:"Query" name:"IOSMutableContent"`
	AndroidNotificationBarPriority int    `position:"Query" name:"AndroidNotificationBarPriority"`
	ExpireTime                     string `position:"Query" name:"ExpireTime"`
	SmsTemplateName                string `position:"Query" name:"SmsTemplateName"`
	AndroidPopupBody               string `position:"Query" name:"AndroidPopupBody"`
	IOSNotificationCategory        string `position:"Query" name:"IOSNotificationCategory"`
	StoreOffline                   string `position:"Query" name:"StoreOffline"`
	IOSSilentNotification          string `position:"Query" name:"IOSSilentNotification"`
	SmsParams                      string `position:"Query" name:"SmsParams"`
	JobKey                         string `position:"Query" name:"JobKey"`
	Target                         string `position:"Query" name:"Target"`
	AndroidOpenUrl                 string `position:"Query" name:"AndroidOpenUrl"`
	AndroidNotificationChannel     string `position:"Query" name:"AndroidNotificationChannel"`
	AndroidRemind                  string `position:"Query" name:"AndroidRemind"`
	AndroidActivity                string `position:"Query" name:"AndroidActivity"`
	AndroidXiaoMiNotifyBody        string `position:"Query" name:"AndroidXiaoMiNotifyBody"`
	IOSSubtitle                    string `position:"Query" name:"IOSSubtitle"`
	SmsSignName                    string `position:"Query" name:"SmsSignName"`
	IOSRemind                      string `position:"Query" name:"IOSRemind"`
	AppKey                         int64  `position:"Query" name:"AppKey"`
	TargetValue                    string `position:"Query" name:"TargetValue"`
	AndroidMusic                   string `position:"Query" name:"AndroidMusic"`
	AndroidXiaoMiActivity          string `position:"Query" name:"AndroidXiaoMiActivity"`
	AndroidXiaoMiNotifyTitle       string `position:"Query" name:"AndroidXiaoMiNotifyTitle"`
	PushType                       string `position:"Query" name:"PushType"`
}

func (req *PushRequest) Invoke(client *sdk.Client) (resp *PushResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "Push", "", "")
	resp = &PushResponse{}
	err = client.DoAction(req, resp)
	return
}

type PushResponse struct {
	responses.BaseResponse
	RequestId string
	MessageId string
}
