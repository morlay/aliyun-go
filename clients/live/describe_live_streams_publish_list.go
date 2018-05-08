package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamsPublishListRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int    `position:"Query" name:"PageSize"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	PageNumber    int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeLiveStreamsPublishListRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamsPublishListResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamsPublishList", "live", "")
	resp = &DescribeLiveStreamsPublishListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamsPublishListResponse struct {
	responses.BaseResponse
	RequestId   common.String
	PageNum     common.Integer
	PageSize    common.Integer
	TotalNum    common.Integer
	TotalPage   common.Integer
	PublishInfo DescribeLiveStreamsPublishListLiveStreamPublishInfoList
}

type DescribeLiveStreamsPublishListLiveStreamPublishInfo struct {
	DomainName    common.String
	AppName       common.String
	StreamName    common.String
	StreamUrl     common.String
	PublishTime   common.String
	StopTime      common.String
	PublishUrl    common.String
	ClientAddr    common.String
	EdgeNodeAddr  common.String
	PublishDomain common.String
}

type DescribeLiveStreamsPublishListLiveStreamPublishInfoList []DescribeLiveStreamsPublishListLiveStreamPublishInfo

func (list *DescribeLiveStreamsPublishListLiveStreamPublishInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamsPublishListLiveStreamPublishInfo)
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
