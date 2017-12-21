package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamOnlineUserNumByDomainRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	HlsSwitch     string `position:"Query" name:"HlsSwitch"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNumber    int64  `position:"Query" name:"PageNumber"`
}

func (req *DescribeLiveStreamOnlineUserNumByDomainRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamOnlineUserNumByDomainResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamOnlineUserNumByDomain", "", "")
	resp = &DescribeLiveStreamOnlineUserNumByDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamOnlineUserNumByDomainResponse struct {
	responses.BaseResponse
	RequestId       string
	TotalUserNumber int64
	Count           int64
	PageNumber      int64
	PageSize        int64
	OnlineUserInfo  DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList
}

type DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfo struct {
	StreamUrl  string
	UserNumber int64
	Time       string
}

type DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList []DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfo

func (list *DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfo)
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
