package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamHistoryUserNumRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveStreamHistoryUserNumRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamHistoryUserNumResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamHistoryUserNum", "live", "")
	resp = &DescribeLiveStreamHistoryUserNumResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamHistoryUserNumResponse struct {
	responses.BaseResponse
	RequestId              string
	LiveStreamUserNumInfos DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfoList
}

type DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfo struct {
	StreamTime string
	UserNum    string
}

type DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfoList []DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfo

func (list *DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfo)
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
