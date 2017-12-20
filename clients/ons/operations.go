package ons

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *OnsClient) OnsRegionList(req *OnsRegionListArgs) (resp *OnsRegionListResponse, err error) {
	resp = &OnsRegionListResponse{}
	err = c.Request("OnsRegionList", req, resp)
	return
}

type OnsRegionListRegionDo struct {
	Id          int64
	OnsRegionId string
	RegionName  string
	ChannelId   int
	ChannelName string
	CreateTime  int64
	UpdateTime  int64
}
type OnsRegionListArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
}
type OnsRegionListResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsRegionListRegionDoList
}

type OnsRegionListRegionDoList []OnsRegionListRegionDo

func (list *OnsRegionListRegionDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsRegionListRegionDo)
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

func (c *OnsClient) OnsTraceQueryByMsgId(req *OnsTraceQueryByMsgIdArgs) (resp *OnsTraceQueryByMsgIdResponse, err error) {
	resp = &OnsTraceQueryByMsgIdResponse{}
	err = c.Request("OnsTraceQueryByMsgId", req, resp)
	return
}

type OnsTraceQueryByMsgIdArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Topic        string
	MsgId        string
	EndTime      int64
	BeginTime    int64
}
type OnsTraceQueryByMsgIdResponse struct {
	RequestId string
	HelpUrl   string
	QueryId   string
}

func (c *OnsClient) OnsMessagePush(req *OnsMessagePushArgs) (resp *OnsMessagePushResponse, err error) {
	resp = &OnsMessagePushResponse{}
	err = c.Request("OnsMessagePush", req, resp)
	return
}

type OnsMessagePushArgs struct {
	PreventCache int64
	OnsRegionId  string
	ClientId     string
	OnsPlatform  string
	ConsumerId   string
	MsgId        string
	Topic        string
}
type OnsMessagePushResponse struct {
	RequestId string
	HelpUrl   string
}

func (c *OnsClient) OnsTopicCreate(req *OnsTopicCreateArgs) (resp *OnsTopicCreateResponse, err error) {
	resp = &OnsTopicCreateResponse{}
	err = c.Request("OnsTopicCreate", req, resp)
	return
}

type OnsTopicCreateArgs struct {
	PreventCache int64
	Cluster      string
	QueueNum     int
	OnsRegionId  string
	OnsPlatform  string
	AppName      string
	Qps          int64
	Topic        string
	Remark       string
	Appkey       string
	Order        core.Bool
	Status       int
}
type OnsTopicCreateResponse struct {
	RequestId string
	HelpUrl   string
}

func (c *OnsClient) OnsTopicStatus(req *OnsTopicStatusArgs) (resp *OnsTopicStatusResponse, err error) {
	resp = &OnsTopicStatusResponse{}
	err = c.Request("OnsTopicStatus", req, resp)
	return
}

type OnsTopicStatusData struct {
	TotalCount    int64
	LastTimeStamp int64
}
type OnsTopicStatusArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Topic        string
}
type OnsTopicStatusResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsTopicStatusData
}

func (c *OnsClient) OnsTopicDelete(req *OnsTopicDeleteArgs) (resp *OnsTopicDeleteResponse, err error) {
	resp = &OnsTopicDeleteResponse{}
	err = c.Request("OnsTopicDelete", req, resp)
	return
}

type OnsTopicDeleteArgs struct {
	PreventCache int64
	Cluster      string
	OnsRegionId  string
	OnsPlatform  string
	Topic        string
}
type OnsTopicDeleteResponse struct {
	RequestId string
	HelpUrl   string
}

func (c *OnsClient) OnsMqttQueryMsgTransTrend(req *OnsMqttQueryMsgTransTrendArgs) (resp *OnsMqttQueryMsgTransTrendResponse, err error) {
	resp = &OnsMqttQueryMsgTransTrendResponse{}
	err = c.Request("OnsMqttQueryMsgTransTrend", req, resp)
	return
}

type OnsMqttQueryMsgTransTrendData struct {
	Title   string
	XUnit   string
	YUnit   string
	Records OnsMqttQueryMsgTransTrendStatsDataDoList
}

type OnsMqttQueryMsgTransTrendStatsDataDo struct {
	X int64
	Y float32
}
type OnsMqttQueryMsgTransTrendArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Qos          int
	TransType    string
	EndTime      int64
	BeginTime    int64
	TpsType      string
	ParentTopic  string
	MsgType      string
	SubTopic     string
}
type OnsMqttQueryMsgTransTrendResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsMqttQueryMsgTransTrendData
}

type OnsMqttQueryMsgTransTrendStatsDataDoList []OnsMqttQueryMsgTransTrendStatsDataDo

func (list *OnsMqttQueryMsgTransTrendStatsDataDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMqttQueryMsgTransTrendStatsDataDo)
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

func (c *OnsClient) OnsTopicGet(req *OnsTopicGetArgs) (resp *OnsTopicGetResponse, err error) {
	resp = &OnsTopicGetResponse{}
	err = c.Request("OnsTopicGet", req, resp)
	return
}

type OnsTopicGetPublishInfoDo struct {
	Id           int64
	ChannelId    int
	ChannelName  string
	OnsRegionId  string
	RegionName   string
	Topic        string
	Owner        string
	Relation     int
	RelationName string
	Status       int
	StatusName   string
	Appkey       string
	CreateTime   int64
	UpdateTime   int64
	Remark       string
}
type OnsTopicGetArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Topic        string
}
type OnsTopicGetResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsTopicGetPublishInfoDoList
}

type OnsTopicGetPublishInfoDoList []OnsTopicGetPublishInfoDo

func (list *OnsTopicGetPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTopicGetPublishInfoDo)
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

func (c *OnsClient) OnsConsumerTimeSpan(req *OnsConsumerTimeSpanArgs) (resp *OnsConsumerTimeSpanResponse, err error) {
	resp = &OnsConsumerTimeSpanResponse{}
	err = c.Request("OnsConsumerTimeSpan", req, resp)
	return
}

type OnsConsumerTimeSpanData struct {
	Topic            string
	MinTimeStamp     int64
	MaxTimeStamp     int64
	ConsumeTimeStamp int64
}
type OnsConsumerTimeSpanArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	ConsumerId   string
	Topic        string
}
type OnsConsumerTimeSpanResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsConsumerTimeSpanData
}

func (c *OnsClient) OnsBuyOrdersProduce(req *OnsBuyOrdersProduceArgs) (resp *OnsBuyOrdersProduceResponse, err error) {
	resp = &OnsBuyOrdersProduceResponse{}
	err = c.Request("OnsBuyOrdersProduce", req, resp)
	return
}

type OnsBuyOrdersProduceArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Data         string
}
type OnsBuyOrdersProduceResponse struct {
	Success   core.Bool
	RequestId string
	Code      string
	Message   string
	Data      string
}

func (c *OnsClient) OnsSubscriptionCreate(req *OnsSubscriptionCreateArgs) (resp *OnsSubscriptionCreateResponse, err error) {
	resp = &OnsSubscriptionCreateResponse{}
	err = c.Request("OnsSubscriptionCreate", req, resp)
	return
}

type OnsSubscriptionCreateArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	ConsumerId   string
	Topic        string
}
type OnsSubscriptionCreateResponse struct {
	RequestId string
	HelpUrl   string
}

func (c *OnsClient) OnsMqttQueryClientByTopic(req *OnsMqttQueryClientByTopicArgs) (resp *OnsMqttQueryClientByTopicResponse, err error) {
	resp = &OnsMqttQueryClientByTopicResponse{}
	err = c.Request("OnsMqttQueryClientByTopic", req, resp)
	return
}

type OnsMqttQueryClientByTopicMqttClientSetDo struct {
	OnlineCount  int64
	PersistCount int64
}
type OnsMqttQueryClientByTopicArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	ParentTopic  string
	SubTopic     string
}
type OnsMqttQueryClientByTopicResponse struct {
	RequestId       string
	HelpUrl         string
	MqttClientSetDo OnsMqttQueryClientByTopicMqttClientSetDo
}

func (c *OnsClient) OnsMqttQueryTraceByTraceId(req *OnsMqttQueryTraceByTraceIdArgs) (resp *OnsMqttQueryTraceByTraceIdResponse, err error) {
	resp = &OnsMqttQueryTraceByTraceIdResponse{}
	err = c.Request("OnsMqttQueryTraceByTraceId", req, resp)
	return
}

type OnsMqttQueryTraceByTraceIdArgs struct {
	TraceId      string
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Topic        string
}
type OnsMqttQueryTraceByTraceIdResponse struct {
	RequestId string
	HelpUrl   string
	PushCount int64
}

func (c *OnsClient) OnsSubscriptionDelete(req *OnsSubscriptionDeleteArgs) (resp *OnsSubscriptionDeleteResponse, err error) {
	resp = &OnsSubscriptionDeleteResponse{}
	err = c.Request("OnsSubscriptionDelete", req, resp)
	return
}

type OnsSubscriptionDeleteArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	ConsumerId   string
	Topic        string
}
type OnsSubscriptionDeleteResponse struct {
	RequestId string
	HelpUrl   string
}

func (c *OnsClient) OnsMqttQueryMsgByPubInfo(req *OnsMqttQueryMsgByPubInfoArgs) (resp *OnsMqttQueryMsgByPubInfoResponse, err error) {
	resp = &OnsMqttQueryMsgByPubInfoResponse{}
	err = c.Request("OnsMqttQueryMsgByPubInfo", req, resp)
	return
}

type OnsMqttQueryMsgByPubInfoMqttMessageDo struct {
	TraceId string
	PubTime int64
}
type OnsMqttQueryMsgByPubInfoArgs struct {
	PreventCache int64
	OnsRegionId  string
	ClientId     string
	OnsPlatform  string
	Topic        string
	EndTime      int64
	BeginTime    int64
}
type OnsMqttQueryMsgByPubInfoResponse struct {
	RequestId      string
	HelpUrl        string
	MqttMessageDos OnsMqttQueryMsgByPubInfoMqttMessageDoList
}

type OnsMqttQueryMsgByPubInfoMqttMessageDoList []OnsMqttQueryMsgByPubInfoMqttMessageDo

func (list *OnsMqttQueryMsgByPubInfoMqttMessageDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMqttQueryMsgByPubInfoMqttMessageDo)
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

func (c *OnsClient) OnsSubscriptionUpdate(req *OnsSubscriptionUpdateArgs) (resp *OnsSubscriptionUpdateResponse, err error) {
	resp = &OnsSubscriptionUpdateResponse{}
	err = c.Request("OnsSubscriptionUpdate", req, resp)
	return
}

type OnsSubscriptionUpdateArgs struct {
	PreventCache int64
	OnsRegionId  string
	ReadEnable   core.Bool
	OnsPlatform  string
	ConsumerId   string
}
type OnsSubscriptionUpdateResponse struct {
	RequestId string
	HelpUrl   string
}

func (c *OnsClient) OnsConsumerResetOffset(req *OnsConsumerResetOffsetArgs) (resp *OnsConsumerResetOffsetResponse, err error) {
	resp = &OnsConsumerResetOffsetResponse{}
	err = c.Request("OnsConsumerResetOffset", req, resp)
	return
}

type OnsConsumerResetOffsetArgs struct {
	PreventCache   int64
	OnsRegionId    string
	OnsPlatform    string
	ConsumerId     string
	Topic          string
	ResetTimestamp int64
	Type           int
}
type OnsConsumerResetOffsetResponse struct {
	RequestId string
	HelpUrl   string
}

func (c *OnsClient) OnsSubscriptionSearch(req *OnsSubscriptionSearchArgs) (resp *OnsSubscriptionSearchResponse, err error) {
	resp = &OnsSubscriptionSearchResponse{}
	err = c.Request("OnsSubscriptionSearch", req, resp)
	return
}

type OnsSubscriptionSearchSubscribeInfoDo struct {
	Id          int64
	ChannelId   int
	ChannelName string
	OnsRegionId string
	RegionName  string
	Owner       string
	ConsumerId  string
	Topic       string
	Status      int
	StatusName  string
	CreateTime  int64
	UpdateTime  int64
}
type OnsSubscriptionSearchArgs struct {
	PreventCache int64
	Search       string
	OnsRegionId  string
	OnsPlatform  string
}
type OnsSubscriptionSearchResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsSubscriptionSearchSubscribeInfoDoList
}

type OnsSubscriptionSearchSubscribeInfoDoList []OnsSubscriptionSearchSubscribeInfoDo

func (list *OnsSubscriptionSearchSubscribeInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsSubscriptionSearchSubscribeInfoDo)
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

func (c *OnsClient) OnsMqttGroupIdList(req *OnsMqttGroupIdListArgs) (resp *OnsMqttGroupIdListResponse, err error) {
	resp = &OnsMqttGroupIdListResponse{}
	err = c.Request("OnsMqttGroupIdList", req, resp)
	return
}

type OnsMqttGroupIdListMqttGroupIdDo struct {
	Id          int64
	ChannelId   int
	OnsRegionId string
	RegionName  string
	Owner       string
	GroupId     string
	Topic       string
	Status      int
	CreateTime  int64
	UpdateTime  int64
}
type OnsMqttGroupIdListArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
}
type OnsMqttGroupIdListResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsMqttGroupIdListMqttGroupIdDoList
}

type OnsMqttGroupIdListMqttGroupIdDoList []OnsMqttGroupIdListMqttGroupIdDo

func (list *OnsMqttGroupIdListMqttGroupIdDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMqttGroupIdListMqttGroupIdDo)
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

func (c *OnsClient) OnsSubscriptionGet(req *OnsSubscriptionGetArgs) (resp *OnsSubscriptionGetResponse, err error) {
	resp = &OnsSubscriptionGetResponse{}
	err = c.Request("OnsSubscriptionGet", req, resp)
	return
}

type OnsSubscriptionGetSubscribeInfoDo struct {
	Id          int64
	ChannelId   int
	ChannelName string
	OnsRegionId string
	RegionName  string
	Owner       string
	ConsumerId  string
	Topic       string
	Status      int
	StatusName  string
	CreateTime  int64
	UpdateTime  int64
}
type OnsSubscriptionGetArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	ConsumerId   string
	Topic        string
}
type OnsSubscriptionGetResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsSubscriptionGetSubscribeInfoDoList
}

type OnsSubscriptionGetSubscribeInfoDoList []OnsSubscriptionGetSubscribeInfoDo

func (list *OnsSubscriptionGetSubscribeInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsSubscriptionGetSubscribeInfoDo)
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

func (c *OnsClient) OnsMessageSend(req *OnsMessageSendArgs) (resp *OnsMessageSendResponse, err error) {
	resp = &OnsMessageSendResponse{}
	err = c.Request("OnsMessageSend", req, resp)
	return
}

type OnsMessageSendArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Topic        string
	ProducerId   string
	Tag          string
	Message      string
	Key          string
}
type OnsMessageSendResponse struct {
	RequestId string
	HelpUrl   string
	Data      string
}

func (c *OnsClient) OnsTopicUpdate(req *OnsTopicUpdateArgs) (resp *OnsTopicUpdateResponse, err error) {
	resp = &OnsTopicUpdateResponse{}
	err = c.Request("OnsTopicUpdate", req, resp)
	return
}

type OnsTopicUpdateArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Perm         int
	Topic        string
}
type OnsTopicUpdateResponse struct {
	RequestId string
	HelpUrl   string
}

func (c *OnsClient) OnsMessagePageQueryByTopic(req *OnsMessagePageQueryByTopicArgs) (resp *OnsMessagePageQueryByTopicResponse, err error) {
	resp = &OnsMessagePageQueryByTopicResponse{}
	err = c.Request("OnsMessagePageQueryByTopic", req, resp)
	return
}

type OnsMessagePageQueryByTopicMsgFoundDo struct {
	TaskId       string
	MaxPageCount int64
	CurrentPage  int64
	MsgFoundList OnsMessagePageQueryByTopicOnsRestMessageDoList
}

type OnsMessagePageQueryByTopicOnsRestMessageDo struct {
	Topic          string
	Flag           int
	Body           string
	StoreSize      int
	BornTimestamp  int64
	BornHost       string
	StoreTimestamp int64
	StoreHost      string
	MsgId          string
	OffsetId       string
	BodyCRC        int
	ReconsumeTimes int
	PropertyList   OnsMessagePageQueryByTopicMessagePropertyList
}

type OnsMessagePageQueryByTopicMessageProperty struct {
	Name  string
	Value string
}
type OnsMessagePageQueryByTopicArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	PageSize     int
	Topic        string
	EndTime      int64
	BeginTime    int64
	CurrentPage  int
	TaskId       string
}
type OnsMessagePageQueryByTopicResponse struct {
	RequestId  string
	HelpUrl    string
	MsgFoundDo OnsMessagePageQueryByTopicMsgFoundDo
}

type OnsMessagePageQueryByTopicOnsRestMessageDoList []OnsMessagePageQueryByTopicOnsRestMessageDo

func (list *OnsMessagePageQueryByTopicOnsRestMessageDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessagePageQueryByTopicOnsRestMessageDo)
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

type OnsMessagePageQueryByTopicMessagePropertyList []OnsMessagePageQueryByTopicMessageProperty

func (list *OnsMessagePageQueryByTopicMessagePropertyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessagePageQueryByTopicMessageProperty)
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

func (c *OnsClient) OnsWarnCreate(req *OnsWarnCreateArgs) (resp *OnsWarnCreateResponse, err error) {
	resp = &OnsWarnCreateResponse{}
	err = c.Request("OnsWarnCreate", req, resp)
	return
}

type OnsWarnCreateArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	BlockTime    string
	Level        string
	ConsumerId   string
	DelayTime    string
	Topic        string
	Threshold    string
	AlertTime    string
	Contacts     string
}
type OnsWarnCreateResponse struct {
	RequestId string
	HelpUrl   string
}

func (c *OnsClient) OnsMessageGetByMsgId(req *OnsMessageGetByMsgIdArgs) (resp *OnsMessageGetByMsgIdResponse, err error) {
	resp = &OnsMessageGetByMsgIdResponse{}
	err = c.Request("OnsMessageGetByMsgId", req, resp)
	return
}

type OnsMessageGetByMsgIdData struct {
	Topic          string
	Flag           int
	Body           string
	StoreSize      int
	BornTimestamp  int64
	BornHost       string
	StoreTimestamp int64
	StoreHost      string
	MsgId          string
	OffsetId       string
	BodyCRC        int
	ReconsumeTimes int
	PropertyList   OnsMessageGetByMsgIdMessagePropertyList
}

type OnsMessageGetByMsgIdMessageProperty struct {
	Name  string
	Value string
}
type OnsMessageGetByMsgIdArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	MsgId        string
	Topic        string
}
type OnsMessageGetByMsgIdResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsMessageGetByMsgIdData
}

type OnsMessageGetByMsgIdMessagePropertyList []OnsMessageGetByMsgIdMessageProperty

func (list *OnsMessageGetByMsgIdMessagePropertyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessageGetByMsgIdMessageProperty)
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

func (c *OnsClient) OnsTopicSearch(req *OnsTopicSearchArgs) (resp *OnsTopicSearchResponse, err error) {
	resp = &OnsTopicSearchResponse{}
	err = c.Request("OnsTopicSearch", req, resp)
	return
}

type OnsTopicSearchPublishInfoDo struct {
	Id           int64
	ChannelId    int
	ChannelName  string
	OnsRegionId  string
	RegionName   string
	Topic        string
	Owner        string
	Relation     int
	RelationName string
	Status       int
	StatusName   string
	Appkey       string
	CreateTime   int64
	UpdateTime   int64
	Remark       string
}
type OnsTopicSearchArgs struct {
	PreventCache int64
	Search       string
	OnsRegionId  string
	OnsPlatform  string
}
type OnsTopicSearchResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsTopicSearchPublishInfoDoList
}

type OnsTopicSearchPublishInfoDoList []OnsTopicSearchPublishInfoDo

func (list *OnsTopicSearchPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTopicSearchPublishInfoDo)
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

func (c *OnsClient) OnsMessageGetByKey(req *OnsMessageGetByKeyArgs) (resp *OnsMessageGetByKeyResponse, err error) {
	resp = &OnsMessageGetByKeyResponse{}
	err = c.Request("OnsMessageGetByKey", req, resp)
	return
}

type OnsMessageGetByKeyOnsRestMessageDo struct {
	Topic          string
	Flag           int
	Body           string
	StoreSize      int
	BornTimestamp  int64
	BornHost       string
	StoreTimestamp int64
	StoreHost      string
	MsgId          string
	OffsetId       string
	BodyCRC        int
	ReconsumeTimes int
	PropertyList   OnsMessageGetByKeyMessagePropertyList
}

type OnsMessageGetByKeyMessageProperty struct {
	Name  string
	Value string
}
type OnsMessageGetByKeyArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Topic        string
	Key          string
}
type OnsMessageGetByKeyResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsMessageGetByKeyOnsRestMessageDoList
}

type OnsMessageGetByKeyMessagePropertyList []OnsMessageGetByKeyMessageProperty

func (list *OnsMessageGetByKeyMessagePropertyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessageGetByKeyMessageProperty)
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

type OnsMessageGetByKeyOnsRestMessageDoList []OnsMessageGetByKeyOnsRestMessageDo

func (list *OnsMessageGetByKeyOnsRestMessageDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessageGetByKeyOnsRestMessageDo)
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

func (c *OnsClient) OnsMqttBuyRefund(req *OnsMqttBuyRefundArgs) (resp *OnsMqttBuyRefundResponse, err error) {
	resp = &OnsMqttBuyRefundResponse{}
	err = c.Request("OnsMqttBuyRefund", req, resp)
	return
}

type OnsMqttBuyRefundArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Data         string
}
type OnsMqttBuyRefundResponse struct {
	Success   core.Bool
	RequestId string
	Code      string
	Message   string
	Data      string
}

func (c *OnsClient) OnsWarnDelete(req *OnsWarnDeleteArgs) (resp *OnsWarnDeleteResponse, err error) {
	resp = &OnsWarnDeleteResponse{}
	err = c.Request("OnsWarnDelete", req, resp)
	return
}

type OnsWarnDeleteArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	ConsumerId   string
	Topic        string
}
type OnsWarnDeleteResponse struct {
	RequestId string
	HelpUrl   string
}

func (c *OnsClient) OnsMessageTrace(req *OnsMessageTraceArgs) (resp *OnsMessageTraceResponse, err error) {
	resp = &OnsMessageTraceResponse{}
	err = c.Request("OnsMessageTrace", req, resp)
	return
}

type OnsMessageTraceMessageTrack struct {
	ConsumerGroup string
	TrackType     string
	ExceptionDesc string
}
type OnsMessageTraceArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Topic        string
	MsgId        string
}
type OnsMessageTraceResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsMessageTraceMessageTrackList
}

type OnsMessageTraceMessageTrackList []OnsMessageTraceMessageTrack

func (list *OnsMessageTraceMessageTrackList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessageTraceMessageTrack)
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

func (c *OnsClient) OnsMqttManualUpdateRule(req *OnsMqttManualUpdateRuleArgs) (resp *OnsMqttManualUpdateRuleResponse, err error) {
	resp = &OnsMqttManualUpdateRuleResponse{}
	err = c.Request("OnsMqttManualUpdateRule", req, resp)
	return
}

type OnsMqttManualUpdateRuleArgs struct {
	PreventCache int64
	OnsRegionId  string
	InstanceId   string
	OnsPlatform  string
	OwnerId      string
	AdminKey     string
}
type OnsMqttManualUpdateRuleResponse struct {
	RequestId string
	HelpUrl   string
}

func (c *OnsClient) OnsMqttQueryHistoryOnline(req *OnsMqttQueryHistoryOnlineArgs) (resp *OnsMqttQueryHistoryOnlineResponse, err error) {
	resp = &OnsMqttQueryHistoryOnlineResponse{}
	err = c.Request("OnsMqttQueryHistoryOnline", req, resp)
	return
}

type OnsMqttQueryHistoryOnlineData struct {
	Title   string
	XUnit   string
	YUnit   string
	Records OnsMqttQueryHistoryOnlineStatsDataDoList
}

type OnsMqttQueryHistoryOnlineStatsDataDo struct {
	X int64
	Y float32
}
type OnsMqttQueryHistoryOnlineArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	GroupId      string
	EndTime      int64
	BeginTime    int64
}
type OnsMqttQueryHistoryOnlineResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsMqttQueryHistoryOnlineData
}

type OnsMqttQueryHistoryOnlineStatsDataDoList []OnsMqttQueryHistoryOnlineStatsDataDo

func (list *OnsMqttQueryHistoryOnlineStatsDataDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMqttQueryHistoryOnlineStatsDataDo)
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

func (c *OnsClient) OnsTraceGetResult(req *OnsTraceGetResultArgs) (resp *OnsTraceGetResultResponse, err error) {
	resp = &OnsTraceGetResultResponse{}
	err = c.Request("OnsTraceGetResult", req, resp)
	return
}

type OnsTraceGetResultTraceData struct {
	QueryId    string
	UserId     string
	Topic      string
	MsgId      string
	MsgKey     string
	Status     string
	CreateTime int64
	UpdateTime int64
	TraceList  OnsTraceGetResultTraceMapDoList
}

type OnsTraceGetResultTraceMapDo struct {
	PubTime      int64
	Topic        string
	PubGroupName string
	MsgId        string
	Tag          string
	MsgKey       string
	BornHost     string
	CostTime     int
	Status       string
	SubList      OnsTraceGetResultSubMapDoList
}

type OnsTraceGetResultSubMapDo struct {
	SubGroupName string
	SuccessCount int
	FailCount    int
	ClientList   OnsTraceGetResultSubClientInfoDoList
}

type OnsTraceGetResultSubClientInfoDo struct {
	SubGroupName   string
	SubTime        int64
	ClientHost     string
	ReconsumeTimes int
	CostTime       int
	Status         string
}
type OnsTraceGetResultArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	QueryId      string
}
type OnsTraceGetResultResponse struct {
	RequestId string
	HelpUrl   string
	TraceData OnsTraceGetResultTraceData
}

type OnsTraceGetResultTraceMapDoList []OnsTraceGetResultTraceMapDo

func (list *OnsTraceGetResultTraceMapDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTraceGetResultTraceMapDo)
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

type OnsTraceGetResultSubMapDoList []OnsTraceGetResultSubMapDo

func (list *OnsTraceGetResultSubMapDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTraceGetResultSubMapDo)
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

type OnsTraceGetResultSubClientInfoDoList []OnsTraceGetResultSubClientInfoDo

func (list *OnsTraceGetResultSubClientInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTraceGetResultSubClientInfoDo)
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

func (c *OnsClient) OnsMqttQueryClientByClientId(req *OnsMqttQueryClientByClientIdArgs) (resp *OnsMqttQueryClientByClientIdResponse, err error) {
	resp = &OnsMqttQueryClientByClientIdResponse{}
	err = c.Request("OnsMqttQueryClientByClientId", req, resp)
	return
}

type OnsMqttQueryClientByClientIdMqttClientInfoDo struct {
	Online          core.Bool
	ClientId        string
	SocketChannel   string
	LastTouch       int64
	SubScriptonData OnsMqttQueryClientByClientIdSubscriptionDoList
}

type OnsMqttQueryClientByClientIdSubscriptionDo struct {
	ParentTopic string
	SubTopic    string
	Qos         int
}
type OnsMqttQueryClientByClientIdArgs struct {
	PreventCache int64
	OnsRegionId  string
	ClientId     string
	OnsPlatform  string
}
type OnsMqttQueryClientByClientIdResponse struct {
	RequestId        string
	HelpUrl          string
	MqttClientInfoDo OnsMqttQueryClientByClientIdMqttClientInfoDo
}

type OnsMqttQueryClientByClientIdSubscriptionDoList []OnsMqttQueryClientByClientIdSubscriptionDo

func (list *OnsMqttQueryClientByClientIdSubscriptionDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMqttQueryClientByClientIdSubscriptionDo)
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

func (c *OnsClient) OnsPublishSearch(req *OnsPublishSearchArgs) (resp *OnsPublishSearchResponse, err error) {
	resp = &OnsPublishSearchResponse{}
	err = c.Request("OnsPublishSearch", req, resp)
	return
}

type OnsPublishSearchPublishInfoDo struct {
	Id          int64
	ChannelId   int
	ChannelName string
	OnsRegionId string
	RegionName  string
	Owner       string
	ProducerId  string
	Topic       string
	Status      int
	StatusName  string
	CreateTime  int64
	UpdateTime  int64
}
type OnsPublishSearchArgs struct {
	PreventCache int64
	Search       string
	OnsRegionId  string
	OnsPlatform  string
}
type OnsPublishSearchResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsPublishSearchPublishInfoDoList
}

type OnsPublishSearchPublishInfoDoList []OnsPublishSearchPublishInfoDo

func (list *OnsPublishSearchPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsPublishSearchPublishInfoDo)
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

func (c *OnsClient) OnsEmpowerList(req *OnsEmpowerListArgs) (resp *OnsEmpowerListResponse, err error) {
	resp = &OnsEmpowerListResponse{}
	err = c.Request("OnsEmpowerList", req, resp)
	return
}

type OnsEmpowerListAuthOwnerInfoDo struct {
	Topic    string
	Owner    int64
	Relation int
}
type OnsEmpowerListArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	EmpowerUser  string
	Topic        string
}
type OnsEmpowerListResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsEmpowerListAuthOwnerInfoDoList
}

type OnsEmpowerListAuthOwnerInfoDoList []OnsEmpowerListAuthOwnerInfoDo

func (list *OnsEmpowerListAuthOwnerInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsEmpowerListAuthOwnerInfoDo)
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

func (c *OnsClient) OnsPublishList(req *OnsPublishListArgs) (resp *OnsPublishListResponse, err error) {
	resp = &OnsPublishListResponse{}
	err = c.Request("OnsPublishList", req, resp)
	return
}

type OnsPublishListPublishInfoDo struct {
	Id          int64
	ChannelId   int
	ChannelName string
	OnsRegionId string
	RegionName  string
	Owner       string
	ProducerId  string
	Topic       string
	Status      int
	StatusName  string
	CreateTime  int64
	UpdateTime  int64
}
type OnsPublishListArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
}
type OnsPublishListResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsPublishListPublishInfoDoList
}

type OnsPublishListPublishInfoDoList []OnsPublishListPublishInfoDo

func (list *OnsPublishListPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsPublishListPublishInfoDo)
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

func (c *OnsClient) OnsTrendTopicInputTps(req *OnsTrendTopicInputTpsArgs) (resp *OnsTrendTopicInputTpsResponse, err error) {
	resp = &OnsTrendTopicInputTpsResponse{}
	err = c.Request("OnsTrendTopicInputTps", req, resp)
	return
}

type OnsTrendTopicInputTpsData struct {
	Title   string
	XUnit   string
	YUnit   string
	Records OnsTrendTopicInputTpsStatsDataDoList
}

type OnsTrendTopicInputTpsStatsDataDo struct {
	X int64
	Y float32
}
type OnsTrendTopicInputTpsArgs struct {
	PreventCache int64
	Period       int64
	OnsRegionId  string
	OnsPlatform  string
	Topic        string
	EndTime      int64
	BeginTime    int64
	Type         int
}
type OnsTrendTopicInputTpsResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsTrendTopicInputTpsData
}

type OnsTrendTopicInputTpsStatsDataDoList []OnsTrendTopicInputTpsStatsDataDo

func (list *OnsTrendTopicInputTpsStatsDataDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTrendTopicInputTpsStatsDataDo)
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

func (c *OnsClient) OnsMqttBuyProduce(req *OnsMqttBuyProduceArgs) (resp *OnsMqttBuyProduceResponse, err error) {
	resp = &OnsMqttBuyProduceResponse{}
	err = c.Request("OnsMqttBuyProduce", req, resp)
	return
}

type OnsMqttBuyProduceArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Data         string
}
type OnsMqttBuyProduceResponse struct {
	Success   core.Bool
	RequestId string
	Code      string
	Message   string
	Data      string
}

func (c *OnsClient) OnsTrendGroupOutputTps(req *OnsTrendGroupOutputTpsArgs) (resp *OnsTrendGroupOutputTpsResponse, err error) {
	resp = &OnsTrendGroupOutputTpsResponse{}
	err = c.Request("OnsTrendGroupOutputTps", req, resp)
	return
}

type OnsTrendGroupOutputTpsData struct {
	Title   string
	XUnit   string
	YUnit   string
	Records OnsTrendGroupOutputTpsStatsDataDoList
}

type OnsTrendGroupOutputTpsStatsDataDo struct {
	X int64
	Y float32
}
type OnsTrendGroupOutputTpsArgs struct {
	PreventCache int64
	Period       int64
	OnsRegionId  string
	OnsPlatform  string
	ConsumerId   string
	Topic        string
	EndTime      int64
	BeginTime    int64
	Type         int
}
type OnsTrendGroupOutputTpsResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsTrendGroupOutputTpsData
}

type OnsTrendGroupOutputTpsStatsDataDoList []OnsTrendGroupOutputTpsStatsDataDo

func (list *OnsTrendGroupOutputTpsStatsDataDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTrendGroupOutputTpsStatsDataDo)
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

func (c *OnsClient) OnsTraceQueryByMsgKey(req *OnsTraceQueryByMsgKeyArgs) (resp *OnsTraceQueryByMsgKeyResponse, err error) {
	resp = &OnsTraceQueryByMsgKeyResponse{}
	err = c.Request("OnsTraceQueryByMsgKey", req, resp)
	return
}

type OnsTraceQueryByMsgKeyArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Topic        string
	EndTime      int64
	BeginTime    int64
	MsgKey       string
}
type OnsTraceQueryByMsgKeyResponse struct {
	RequestId string
	HelpUrl   string
	QueryId   string
}

func (c *OnsClient) OnsConsumerAccumulate(req *OnsConsumerAccumulateArgs) (resp *OnsConsumerAccumulateResponse, err error) {
	resp = &OnsConsumerAccumulateResponse{}
	err = c.Request("OnsConsumerAccumulate", req, resp)
	return
}

type OnsConsumerAccumulateData struct {
	Online            core.Bool
	TotalDiff         int64
	ConsumeTps        float32
	LastTimestamp     int64
	DelayTime         int64
	DetailInTopicList OnsConsumerAccumulateDetailInTopicDoList
}

type OnsConsumerAccumulateDetailInTopicDo struct {
	Topic         string
	TotalDiff     int64
	LastTimestamp int64
	DelayTime     int64
}
type OnsConsumerAccumulateArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	ConsumerId   string
	Detail       core.Bool
}
type OnsConsumerAccumulateResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsConsumerAccumulateData
}

type OnsConsumerAccumulateDetailInTopicDoList []OnsConsumerAccumulateDetailInTopicDo

func (list *OnsConsumerAccumulateDetailInTopicDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerAccumulateDetailInTopicDo)
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

func (c *OnsClient) OnsMqttBuyCheck(req *OnsMqttBuyCheckArgs) (resp *OnsMqttBuyCheckResponse, err error) {
	resp = &OnsMqttBuyCheckResponse{}
	err = c.Request("OnsMqttBuyCheck", req, resp)
	return
}

type OnsMqttBuyCheckArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Data         string
}
type OnsMqttBuyCheckResponse struct {
	Success   core.Bool
	RequestId string
	Code      string
	Message   string
	Data      string
}

func (c *OnsClient) OnsTopicList(req *OnsTopicListArgs) (resp *OnsTopicListResponse, err error) {
	resp = &OnsTopicListResponse{}
	err = c.Request("OnsTopicList", req, resp)
	return
}

type OnsTopicListPublishInfoDo struct {
	Id           int64
	ChannelId    int
	ChannelName  string
	OnsRegionId  string
	RegionName   string
	Topic        string
	Owner        string
	Relation     int
	RelationName string
	Status       int
	StatusName   string
	Appkey       string
	CreateTime   int64
	UpdateTime   int64
	Remark       string
}
type OnsTopicListArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Topic        string
}
type OnsTopicListResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsTopicListPublishInfoDoList
}

type OnsTopicListPublishInfoDoList []OnsTopicListPublishInfoDo

func (list *OnsTopicListPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTopicListPublishInfoDo)
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

func (c *OnsClient) OnsSubscriptionList(req *OnsSubscriptionListArgs) (resp *OnsSubscriptionListResponse, err error) {
	resp = &OnsSubscriptionListResponse{}
	err = c.Request("OnsSubscriptionList", req, resp)
	return
}

type OnsSubscriptionListSubscribeInfoDo struct {
	Id          int64
	ChannelId   int
	ChannelName string
	OnsRegionId string
	RegionName  string
	Owner       string
	ConsumerId  string
	Topic       string
	Status      int
	StatusName  string
	CreateTime  int64
	UpdateTime  int64
}
type OnsSubscriptionListArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
}
type OnsSubscriptionListResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsSubscriptionListSubscribeInfoDoList
}

type OnsSubscriptionListSubscribeInfoDoList []OnsSubscriptionListSubscribeInfoDo

func (list *OnsSubscriptionListSubscribeInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsSubscriptionListSubscribeInfoDo)
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

func (c *OnsClient) OnsConsumerStatus(req *OnsConsumerStatusArgs) (resp *OnsConsumerStatusResponse, err error) {
	resp = &OnsConsumerStatusResponse{}
	err = c.Request("OnsConsumerStatus", req, resp)
	return
}

type OnsConsumerStatusData struct {
	Online                     core.Bool
	TotalDiff                  int64
	ConsumeTps                 float32
	LastTimestamp              int64
	DelayTime                  int64
	ConsumeModel               string
	SubscriptionSame           core.Bool
	RebalanceOK                core.Bool
	ConnectionSet              OnsConsumerStatusConnectionDoList
	DetailInTopicList          OnsConsumerStatusDetailInTopicDoList
	ConsumerConnectionInfoList OnsConsumerStatusConsumerConnectionInfoDoList
}

type OnsConsumerStatusConnectionDo struct {
	ClientId   string
	ClientAddr string
	Language   string
	Version    string
}

type OnsConsumerStatusDetailInTopicDo struct {
	Topic         string
	TotalDiff     int64
	LastTimestamp int64
	DelayTime     int64
}

type OnsConsumerStatusConsumerConnectionInfoDo struct {
	ClientId        string
	Connection      string
	Language        string
	Version         string
	ConsumeModel    string
	ConsumeType     string
	ThreadCount     int
	StartTimeStamp  int64
	LastTimeStamp   int64
	SubscriptionSet OnsConsumerStatusSubscriptionDataList
	RunningDataList OnsConsumerStatusConsumerRunningDataDoList
	Jstack          OnsConsumerStatusThreadTrackDoList
}

type OnsConsumerStatusSubscriptionData struct {
	Topic      string
	SubString  string
	SubVersion int64
	TagsSet    OnsConsumerStatusTagsSetList
}

type OnsConsumerStatusConsumerRunningDataDo struct {
	ConsumerId         string
	Topic              string
	Rt                 float32
	OkTps              float32
	FailedTps          float32
	FailedCountPerHour int64
}

type OnsConsumerStatusThreadTrackDo struct {
	Thread    string
	TrackList OnsConsumerStatusTrackListList
}
type OnsConsumerStatusArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	NeedJstack   core.Bool
	ConsumerId   string
	Detail       core.Bool
}
type OnsConsumerStatusResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsConsumerStatusData
}

type OnsConsumerStatusConnectionDoList []OnsConsumerStatusConnectionDo

func (list *OnsConsumerStatusConnectionDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusConnectionDo)
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

type OnsConsumerStatusDetailInTopicDoList []OnsConsumerStatusDetailInTopicDo

func (list *OnsConsumerStatusDetailInTopicDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusDetailInTopicDo)
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

type OnsConsumerStatusConsumerConnectionInfoDoList []OnsConsumerStatusConsumerConnectionInfoDo

func (list *OnsConsumerStatusConsumerConnectionInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusConsumerConnectionInfoDo)
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

type OnsConsumerStatusSubscriptionDataList []OnsConsumerStatusSubscriptionData

func (list *OnsConsumerStatusSubscriptionDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusSubscriptionData)
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

type OnsConsumerStatusConsumerRunningDataDoList []OnsConsumerStatusConsumerRunningDataDo

func (list *OnsConsumerStatusConsumerRunningDataDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusConsumerRunningDataDo)
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

type OnsConsumerStatusThreadTrackDoList []OnsConsumerStatusThreadTrackDo

func (list *OnsConsumerStatusThreadTrackDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusThreadTrackDo)
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

type OnsConsumerStatusTagsSetList []string

func (list *OnsConsumerStatusTagsSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type OnsConsumerStatusTrackListList []string

func (list *OnsConsumerStatusTrackListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

func (c *OnsClient) OnsPublishCreate(req *OnsPublishCreateArgs) (resp *OnsPublishCreateResponse, err error) {
	resp = &OnsPublishCreateResponse{}
	err = c.Request("OnsPublishCreate", req, resp)
	return
}

type OnsPublishCreateArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	AppName      string
	Topic        string
	ProducerId   string
}
type OnsPublishCreateResponse struct {
	RequestId string
	HelpUrl   string
}

func (c *OnsClient) OnsEmpowerDelete(req *OnsEmpowerDeleteArgs) (resp *OnsEmpowerDeleteResponse, err error) {
	resp = &OnsEmpowerDeleteResponse{}
	err = c.Request("OnsEmpowerDelete", req, resp)
	return
}

type OnsEmpowerDeleteArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	EmpowerUser  string
	Topic        string
}
type OnsEmpowerDeleteResponse struct {
	RequestId string
	HelpUrl   string
}

func (c *OnsClient) OnsConsumerGetConnection(req *OnsConsumerGetConnectionArgs) (resp *OnsConsumerGetConnectionResponse, err error) {
	resp = &OnsConsumerGetConnectionResponse{}
	err = c.Request("OnsConsumerGetConnection", req, resp)
	return
}

type OnsConsumerGetConnectionData struct {
	ConnectionList OnsConsumerGetConnectionConnectionDoList
}

type OnsConsumerGetConnectionConnectionDo struct {
	ClientId   string
	ClientAddr string
	Language   string
	Version    string
}
type OnsConsumerGetConnectionArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	ConsumerId   string
}
type OnsConsumerGetConnectionResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsConsumerGetConnectionData
}

type OnsConsumerGetConnectionConnectionDoList []OnsConsumerGetConnectionConnectionDo

func (list *OnsConsumerGetConnectionConnectionDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerGetConnectionConnectionDo)
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

func (c *OnsClient) OnsMqttQueryClientByGroupId(req *OnsMqttQueryClientByGroupIdArgs) (resp *OnsMqttQueryClientByGroupIdResponse, err error) {
	resp = &OnsMqttQueryClientByGroupIdResponse{}
	err = c.Request("OnsMqttQueryClientByGroupId", req, resp)
	return
}

type OnsMqttQueryClientByGroupIdMqttClientSetDo struct {
	OnlineCount  int64
	PersistCount int64
}
type OnsMqttQueryClientByGroupIdArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	GroupId      string
}
type OnsMqttQueryClientByGroupIdResponse struct {
	RequestId       string
	HelpUrl         string
	MqttClientSetDo OnsMqttQueryClientByGroupIdMqttClientSetDo
}

func (c *OnsClient) OnsPublishGet(req *OnsPublishGetArgs) (resp *OnsPublishGetResponse, err error) {
	resp = &OnsPublishGetResponse{}
	err = c.Request("OnsPublishGet", req, resp)
	return
}

type OnsPublishGetPublishInfoDo struct {
	Id          int64
	ChannelId   int
	ChannelName string
	OnsRegionId string
	RegionName  string
	Owner       string
	ProducerId  string
	Topic       string
	Status      int
	StatusName  string
	CreateTime  int64
	UpdateTime  int64
}
type OnsPublishGetArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Topic        string
	ProducerId   string
}
type OnsPublishGetResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsPublishGetPublishInfoDoList
}

type OnsPublishGetPublishInfoDoList []OnsPublishGetPublishInfoDo

func (list *OnsPublishGetPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsPublishGetPublishInfoDo)
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

func (c *OnsClient) OnsPublishDelete(req *OnsPublishDeleteArgs) (resp *OnsPublishDeleteResponse, err error) {
	resp = &OnsPublishDeleteResponse{}
	err = c.Request("OnsPublishDelete", req, resp)
	return
}

type OnsPublishDeleteArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	Topic        string
	ProducerId   string
}
type OnsPublishDeleteResponse struct {
	RequestId string
	HelpUrl   string
}

func (c *OnsClient) OnsEmpowerCreate(req *OnsEmpowerCreateArgs) (resp *OnsEmpowerCreateResponse, err error) {
	resp = &OnsEmpowerCreateResponse{}
	err = c.Request("OnsEmpowerCreate", req, resp)
	return
}

type OnsEmpowerCreateArgs struct {
	PreventCache int64
	OnsRegionId  string
	OnsPlatform  string
	EmpowerUser  string
	Topic        string
	Relation     int
}
type OnsEmpowerCreateResponse struct {
	RequestId string
	HelpUrl   string
}
