package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamsOnlineListRequest struct {
	requests.RpcRequest
	StreamType    string `position:"Query" name:"StreamType"`
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int    `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNum       int    `position:"Query" name:"PageNum"`
}

func (req *DescribeLiveStreamsOnlineListRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamsOnlineListResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamsOnlineList", "live", "")
	resp = &DescribeLiveStreamsOnlineListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamsOnlineListResponse struct {
	responses.BaseResponse
	RequestId  string
	PageNum    int
	PageSize   int
	TotalNum   int
	TotalPage  int
	OnlineInfo DescribeLiveStreamsOnlineListLiveStreamOnlineInfoList
}

type DescribeLiveStreamsOnlineListLiveStreamOnlineInfo struct {
	DomainName    string
	AppName       string
	StreamName    string
	PublishTime   string
	PublishUrl    string
	PublishDomain string
}

type DescribeLiveStreamsOnlineListLiveStreamOnlineInfoList []DescribeLiveStreamsOnlineListLiveStreamOnlineInfo

func (list *DescribeLiveStreamsOnlineListLiveStreamOnlineInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamsOnlineListLiveStreamOnlineInfo)
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
