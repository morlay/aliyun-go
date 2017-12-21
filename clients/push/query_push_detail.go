package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryPushDetailRequest struct {
	MessageId string `position:"Query" name:"MessageId"`
	AppKey    int64  `position:"Query" name:"AppKey"`
}

func (r QueryPushDetailRequest) Invoke(client *sdk.Client) (response *QueryPushDetailResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryPushDetailRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "QueryPushDetail", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryPushDetailResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryPushDetailResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryPushDetailResponse struct {
	RequestId                      string
	AppKey                         int64
	Target                         string
	TargetValue                    string
	PushType                       string
	DeviceType                     string
	Title                          string
	Body                           string
	PushTime                       string
	ExpireTime                     string
	AntiHarassStartTime            int
	AntiHarassDuration             int
	StoreOffline                   bool
	BatchNumber                    string
	ProvinceId                     string
	IOSApnsEnv                     string
	IOSRemind                      bool
	IOSRemindBody                  string
	IOSBadge                       int
	IOSMusic                       string
	IOSSubtitle                    string
	IOSNotificationCategory        string
	IOSMutableContent              bool
	IOSExtParameters               string
	AndroidNotifyType              string
	AndroidOpenType                string
	AndroidActivity                string
	AndroidMusic                   string
	AndroidOpenUrl                 string
	AndroidXiaoMiActivity          string
	AndroidXiaoMiNotifyTitle       string
	AndroidXiaoMiNotifyBody        string
	AndroidPopupActivity           string
	AndroidPopupTitle              string
	AndroidPopupBody               string
	AndroidNotificationBarType     int
	AndroidNotificationBarPriority int
	AndroidExtParameters           string
	SmsTemplateName                string
	SmsSignName                    string
	SmsParams                      string
	SmsDelaySecs                   int
	SmsSendPolicy                  int
}
