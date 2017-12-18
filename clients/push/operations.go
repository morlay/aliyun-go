package push

import "encoding/json"

func (c *PushClient) PushNoticeToAndroid(req *PushNoticeToAndroidArgs) (resp *PushNoticeToAndroidResponse, err error) {
	resp = &PushNoticeToAndroidResponse{}
	err = c.Request("PushNoticeToAndroid", req, resp)
	return
}

type PushNoticeToAndroidArgs struct {
	ExtParameters string
	AppKey        int64
	TargetValue   string
	Title         string
	Body          string
	JobKey        string
	Target        string
}
type PushNoticeToAndroidResponse struct {
	RequestId string
	MessageId string
}

func (c *PushClient) UnbindAlias(req *UnbindAliasArgs) (resp *UnbindAliasResponse, err error) {
	resp = &UnbindAliasResponse{}
	err = c.Request("UnbindAlias", req, resp)
	return
}

type UnbindAliasArgs struct {
	AliasName string
	AppKey    int64
	DeviceId  string
	UnbindAll bool
}
type UnbindAliasResponse struct {
	RequestId string
}

func (c *PushClient) PushMessageToAndroid(req *PushMessageToAndroidArgs) (resp *PushMessageToAndroidResponse, err error) {
	resp = &PushMessageToAndroidResponse{}
	err = c.Request("PushMessageToAndroid", req, resp)
	return
}

type PushMessageToAndroidArgs struct {
	AppKey      int64
	TargetValue string
	Title       string
	Body        string
	JobKey      string
	Target      string
}
type PushMessageToAndroidResponse struct {
	RequestId string
	MessageId string
}

func (c *PushClient) QueryDeviceInfo(req *QueryDeviceInfoArgs) (resp *QueryDeviceInfoResponse, err error) {
	resp = &QueryDeviceInfoResponse{}
	err = c.Request("QueryDeviceInfo", req, resp)
	return
}

type QueryDeviceInfoDeviceInfo struct {
	DeviceId       string
	DeviceType     string
	Account        string
	DeviceToken    string
	Tags           string
	Alias          string
	LastOnlineTime string
	Online         bool
}
type QueryDeviceInfoArgs struct {
	AppKey   int64
	DeviceId string
}
type QueryDeviceInfoResponse struct {
	RequestId  string
	DeviceInfo QueryDeviceInfoDeviceInfo
}

func (c *PushClient) CheckDevice(req *CheckDeviceArgs) (resp *CheckDeviceResponse, err error) {
	resp = &CheckDeviceResponse{}
	err = c.Request("CheckDevice", req, resp)
	return
}

type CheckDeviceArgs struct {
	AppKey   int64
	DeviceId string
}
type CheckDeviceResponse struct {
	RequestId string
	Available bool
}

func (c *PushClient) QueryPushDetail(req *QueryPushDetailArgs) (resp *QueryPushDetailResponse, err error) {
	resp = &QueryPushDetailResponse{}
	err = c.Request("QueryPushDetail", req, resp)
	return
}

type QueryPushDetailArgs struct {
	MessageId string
	AppKey    int64
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

func (c *PushClient) Push(req *PushArgs) (resp *PushResponse, err error) {
	resp = &PushResponse{}
	err = c.Request("Push", req, resp)
	return
}

type PushArgs struct {
	AndroidNotificationBarType     int
	SmsSendPolicy                  int
	AndroidExtParameters           string
	IOSBadge                       int
	IOSBadgeAutoIncrement          bool
	AndroidOpenType                string
	Title                          string
	Body                           string
	DeviceType                     string
	PushTime                       string
	SmsDelaySecs                   int
	SendSpeed                      int
	AndroidPopupActivity           string
	IOSRemindBody                  string
	BatchNumber                    string
	IOSExtParameters               string
	AndroidNotifyType              string
	AndroidPopupTitle              string
	IOSMusic                       string
	IOSApnsEnv                     string
	IOSMutableContent              bool
	AndroidNotificationBarPriority int
	ExpireTime                     string
	SmsTemplateName                string
	AndroidPopupBody               string
	IOSNotificationCategory        string
	StoreOffline                   bool
	IOSSilentNotification          bool
	SmsParams                      string
	JobKey                         string
	Target                         string
	AndroidOpenUrl                 string
	AndroidRemind                  bool
	AndroidActivity                string
	AndroidXiaoMiNotifyBody        string
	IOSSubtitle                    string
	SmsSignName                    string
	IOSRemind                      bool
	AppKey                         int64
	TargetValue                    string
	AndroidMusic                   string
	AndroidXiaoMiActivity          string
	AndroidXiaoMiNotifyTitle       string
	PushType                       string
}
type PushResponse struct {
	RequestId string
	MessageId string
}

func (c *PushClient) RemoveTag(req *RemoveTagArgs) (resp *RemoveTagResponse, err error) {
	resp = &RemoveTagResponse{}
	err = c.Request("RemoveTag", req, resp)
	return
}

type RemoveTagArgs struct {
	TagName string
	AppKey  int64
}
type RemoveTagResponse struct {
	RequestId string
}

func (c *PushClient) PushNoticeToiOS(req *PushNoticeToiOSArgs) (resp *PushNoticeToiOSResponse, err error) {
	resp = &PushNoticeToiOSResponse{}
	err = c.Request("PushNoticeToiOS", req, resp)
	return
}

type PushNoticeToiOSArgs struct {
	ExtParameters string
	ApnsEnv       string
	AppKey        int64
	TargetValue   string
	Title         string
	Body          string
	JobKey        string
	Target        string
}
type PushNoticeToiOSResponse struct {
	RequestId string
	MessageId string
}

func (c *PushClient) BindAlias(req *BindAliasArgs) (resp *BindAliasResponse, err error) {
	resp = &BindAliasResponse{}
	err = c.Request("BindAlias", req, resp)
	return
}

type BindAliasArgs struct {
	AliasName string
	AppKey    int64
	DeviceId  string
}
type BindAliasResponse struct {
	RequestId string
}

func (c *PushClient) ListTags(req *ListTagsArgs) (resp *ListTagsResponse, err error) {
	resp = &ListTagsResponse{}
	err = c.Request("ListTags", req, resp)
	return
}

type ListTagsTagInfo struct {
	TagName string
}
type ListTagsArgs struct {
	AppKey int64
}
type ListTagsResponse struct {
	RequestId string
	TagInfos  ListTagsTagInfoList
}

type ListTagsTagInfoList []ListTagsTagInfo

func (list *ListTagsTagInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagsTagInfo)
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

func (c *PushClient) CancelPush(req *CancelPushArgs) (resp *CancelPushResponse, err error) {
	resp = &CancelPushResponse{}
	err = c.Request("CancelPush", req, resp)
	return
}

type CancelPushArgs struct {
	MessageId string
	AppKey    int64
}
type CancelPushResponse struct {
	RequestId string
}

func (c *PushClient) QueryPushStatByMsg(req *QueryPushStatByMsgArgs) (resp *QueryPushStatByMsgResponse, err error) {
	resp = &QueryPushStatByMsgResponse{}
	err = c.Request("QueryPushStatByMsg", req, resp)
	return
}

type QueryPushStatByMsgPushStat struct {
	MessageId              string
	AcceptCount            int64
	SentCount              int64
	ReceivedCount          int64
	OpenedCount            int64
	DeletedCount           int64
	SmsSentCount           int64
	SmsSkipCount           int64
	SmsFailedCount         int64
	SmsReceiveSuccessCount int64
	SmsReceiveFailedCount  int64
}
type QueryPushStatByMsgArgs struct {
	MessageId string
	AppKey    int64
}
type QueryPushStatByMsgResponse struct {
	RequestId string
	PushStats QueryPushStatByMsgPushStatList
}

type QueryPushStatByMsgPushStatList []QueryPushStatByMsgPushStat

func (list *QueryPushStatByMsgPushStatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPushStatByMsgPushStat)
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

func (c *PushClient) BindTag(req *BindTagArgs) (resp *BindTagResponse, err error) {
	resp = &BindTagResponse{}
	err = c.Request("BindTag", req, resp)
	return
}

type BindTagArgs struct {
	TagName   string
	ClientKey string
	AppKey    int64
	KeyType   string
}
type BindTagResponse struct {
	RequestId string
}

func (c *PushClient) ListSummaryApps(req *ListSummaryAppsArgs) (resp *ListSummaryAppsResponse, err error) {
	resp = &ListSummaryAppsResponse{}
	err = c.Request("ListSummaryApps", req, resp)
	return
}

type ListSummaryAppsSummaryAppInfo struct {
	AppName string
	AppKey  int64
}
type ListSummaryAppsArgs struct {
}
type ListSummaryAppsResponse struct {
	RequestId       string
	SummaryAppInfos ListSummaryAppsSummaryAppInfoList
}

type ListSummaryAppsSummaryAppInfoList []ListSummaryAppsSummaryAppInfo

func (list *ListSummaryAppsSummaryAppInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSummaryAppsSummaryAppInfo)
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

func (c *PushClient) QueryUniqueDeviceStat(req *QueryUniqueDeviceStatArgs) (resp *QueryUniqueDeviceStatResponse, err error) {
	resp = &QueryUniqueDeviceStatResponse{}
	err = c.Request("QueryUniqueDeviceStat", req, resp)
	return
}

type QueryUniqueDeviceStatAppDeviceStat struct {
	Time  string
	Count int64
}
type QueryUniqueDeviceStatArgs struct {
	Granularity string
	EndTime     string
	AppKey      int64
	StartTime   string
}
type QueryUniqueDeviceStatResponse struct {
	RequestId      string
	AppDeviceStats QueryUniqueDeviceStatAppDeviceStatList
}

type QueryUniqueDeviceStatAppDeviceStatList []QueryUniqueDeviceStatAppDeviceStat

func (list *QueryUniqueDeviceStatAppDeviceStatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryUniqueDeviceStatAppDeviceStat)
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

func (c *PushClient) QueryDeviceStat(req *QueryDeviceStatArgs) (resp *QueryDeviceStatResponse, err error) {
	resp = &QueryDeviceStatResponse{}
	err = c.Request("QueryDeviceStat", req, resp)
	return
}

type QueryDeviceStatAppDeviceStat struct {
	Time       string
	Count      int64
	DeviceType string
}
type QueryDeviceStatArgs struct {
	EndTime    string
	AppKey     int64
	StartTime  string
	DeviceType string
	QueryType  string
}
type QueryDeviceStatResponse struct {
	RequestId      string
	AppDeviceStats QueryDeviceStatAppDeviceStatList
}

type QueryDeviceStatAppDeviceStatList []QueryDeviceStatAppDeviceStat

func (list *QueryDeviceStatAppDeviceStatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDeviceStatAppDeviceStat)
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

func (c *PushClient) QueryAliases(req *QueryAliasesArgs) (resp *QueryAliasesResponse, err error) {
	resp = &QueryAliasesResponse{}
	err = c.Request("QueryAliases", req, resp)
	return
}

type QueryAliasesAliasInfo struct {
	AliasName string
}
type QueryAliasesArgs struct {
	AppKey   int64
	DeviceId string
}
type QueryAliasesResponse struct {
	RequestId  string
	AliasInfos QueryAliasesAliasInfoList
}

type QueryAliasesAliasInfoList []QueryAliasesAliasInfo

func (list *QueryAliasesAliasInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAliasesAliasInfo)
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

func (c *PushClient) PushMessageToiOS(req *PushMessageToiOSArgs) (resp *PushMessageToiOSResponse, err error) {
	resp = &PushMessageToiOSResponse{}
	err = c.Request("PushMessageToiOS", req, resp)
	return
}

type PushMessageToiOSArgs struct {
	AppKey      int64
	TargetValue string
	Title       string
	Body        string
	JobKey      string
	Target      string
}
type PushMessageToiOSResponse struct {
	RequestId string
	MessageId string
}

func (c *PushClient) ListPushRecords(req *ListPushRecordsArgs) (resp *ListPushRecordsResponse, err error) {
	resp = &ListPushRecordsResponse{}
	err = c.Request("ListPushRecords", req, resp)
	return
}

type ListPushRecordsPushMessageInfo struct {
	AppKey     int64
	AppName    string
	MessageId  string
	Type       string
	DeviceType string
	PushTime   string
	Title      string
	Body       string
}
type ListPushRecordsArgs struct {
	PageSize  int
	EndTime   string
	AppKey    int64
	StartTime string
	Page      int
	PushType  string
}
type ListPushRecordsResponse struct {
	RequestId        string
	Total            int
	Page             int
	PageSize         int
	PushMessageInfos ListPushRecordsPushMessageInfoList
}

type ListPushRecordsPushMessageInfoList []ListPushRecordsPushMessageInfo

func (list *ListPushRecordsPushMessageInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPushRecordsPushMessageInfo)
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

func (c *PushClient) QueryPushStatByApp(req *QueryPushStatByAppArgs) (resp *QueryPushStatByAppResponse, err error) {
	resp = &QueryPushStatByAppResponse{}
	err = c.Request("QueryPushStatByApp", req, resp)
	return
}

type QueryPushStatByAppAppPushStat struct {
	Time                   string
	AcceptCount            int64
	SentCount              int64
	ReceivedCount          int64
	OpenedCount            int64
	DeletedCount           int64
	SmsSentCount           int64
	SmsSkipCount           int64
	SmsFailedCount         int64
	SmsReceiveSuccessCount int64
	SmsReceiveFailedCount  int64
}
type QueryPushStatByAppArgs struct {
	Granularity string
	EndTime     string
	AppKey      int64
	StartTime   string
}
type QueryPushStatByAppResponse struct {
	RequestId    string
	AppPushStats QueryPushStatByAppAppPushStatList
}

type QueryPushStatByAppAppPushStatList []QueryPushStatByAppAppPushStat

func (list *QueryPushStatByAppAppPushStatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPushStatByAppAppPushStat)
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

func (c *PushClient) UnbindPhone(req *UnbindPhoneArgs) (resp *UnbindPhoneResponse, err error) {
	resp = &UnbindPhoneResponse{}
	err = c.Request("UnbindPhone", req, resp)
	return
}

type UnbindPhoneArgs struct {
	AppKey   int64
	DeviceId string
}
type UnbindPhoneResponse struct {
	RequestId string
}

func (c *PushClient) QueryTags(req *QueryTagsArgs) (resp *QueryTagsResponse, err error) {
	resp = &QueryTagsResponse{}
	err = c.Request("QueryTags", req, resp)
	return
}

type QueryTagsTagInfo struct {
	TagName string
}
type QueryTagsArgs struct {
	ClientKey string
	AppKey    int64
	KeyType   string
}
type QueryTagsResponse struct {
	RequestId string
	TagInfos  QueryTagsTagInfoList
}

type QueryTagsTagInfoList []QueryTagsTagInfo

func (list *QueryTagsTagInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTagsTagInfo)
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

func (c *PushClient) UnbindTag(req *UnbindTagArgs) (resp *UnbindTagResponse, err error) {
	resp = &UnbindTagResponse{}
	err = c.Request("UnbindTag", req, resp)
	return
}

type UnbindTagArgs struct {
	TagName   string
	ClientKey string
	AppKey    int64
	KeyType   string
}
type UnbindTagResponse struct {
	RequestId string
}

func (c *PushClient) CheckDevices(req *CheckDevicesArgs) (resp *CheckDevicesResponse, err error) {
	resp = &CheckDevicesResponse{}
	err = c.Request("CheckDevices", req, resp)
	return
}

type CheckDevicesDeviceCheckInfo struct {
	DeviceId  string
	Available bool
}
type CheckDevicesArgs struct {
	DeviceIds string
	AppKey    int64
}
type CheckDevicesResponse struct {
	RequestId        string
	DeviceCheckInfos CheckDevicesDeviceCheckInfoList
}

type CheckDevicesDeviceCheckInfoList []CheckDevicesDeviceCheckInfo

func (list *CheckDevicesDeviceCheckInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CheckDevicesDeviceCheckInfo)
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

func (c *PushClient) BindPhone(req *BindPhoneArgs) (resp *BindPhoneResponse, err error) {
	resp = &BindPhoneResponse{}
	err = c.Request("BindPhone", req, resp)
	return
}

type BindPhoneArgs struct {
	PhoneNumber string
	AppKey      int64
	DeviceId    string
}
type BindPhoneResponse struct {
	RequestId string
}
