package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamsControlHistoryRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveStreamsControlHistoryRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamsControlHistoryResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamsControlHistory", "live", "")
	resp = &DescribeLiveStreamsControlHistoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamsControlHistoryResponse struct {
	responses.BaseResponse
	RequestId   common.String
	ControlInfo DescribeLiveStreamsControlHistoryLiveStreamControlInfoList
}

type DescribeLiveStreamsControlHistoryLiveStreamControlInfo struct {
	StreamName common.String
	ClientIP   common.String
	Action     common.String
	TimeStamp  common.String
}

type DescribeLiveStreamsControlHistoryLiveStreamControlInfoList []DescribeLiveStreamsControlHistoryLiveStreamControlInfo

func (list *DescribeLiveStreamsControlHistoryLiveStreamControlInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamsControlHistoryLiveStreamControlInfo)
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
