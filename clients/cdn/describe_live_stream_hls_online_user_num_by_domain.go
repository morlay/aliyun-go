package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamHlsOnlineUserNumByDomainRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	HlsSwitch     string `position:"Query" name:"HlsSwitch"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNumber    int64  `position:"Query" name:"PageNumber"`
}

func (req *DescribeLiveStreamHlsOnlineUserNumByDomainRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamHlsOnlineUserNumByDomainResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamHlsOnlineUserNumByDomain", "", "")
	resp = &DescribeLiveStreamHlsOnlineUserNumByDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamHlsOnlineUserNumByDomainResponse struct {
	responses.BaseResponse
	RequestId       string
	TotalUserNumber int64
	Count           int64
	PageNumber      int64
	PageSize        int64
	OnlineUserInfo  DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList
}

type DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfo struct {
	StreamUrl  string
	UserNumber int64
	Time       string
}

type DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList []DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfo

func (list *DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfo)
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
