package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamOnlineUserNumRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	HlsSwitch     string `position:"Query" name:"HlsSwitch"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (r DescribeLiveStreamOnlineUserNumRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamOnlineUserNumResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamOnlineUserNumRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamOnlineUserNum", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamOnlineUserNumResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamOnlineUserNumResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamOnlineUserNumResponse struct {
	RequestId       string
	TotalUserNumber int64
	OnlineUserInfo  DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfoList
}

type DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfo struct {
	StreamUrl  string
	UserNumber int64
	Time       string
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
