package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamOnlineUserNumRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	HlsSwitch     string `position:"Query" name:"HlsSwitch"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveStreamOnlineUserNumRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamOnlineUserNumResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamOnlineUserNum", "live", "")
	resp = &DescribeLiveStreamOnlineUserNumResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamOnlineUserNumResponse struct {
	responses.BaseResponse
	RequestId       common.String
	TotalUserNumber common.Long
	OnlineUserInfo  DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfoList
}

type DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfo struct {
	StreamUrl  common.String
	UserNumber common.Long
	Time       common.String
}

type DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfoList []DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfo

func (list *DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfo)
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
