package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamRoomUserNumberRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	HlsSwitch     string `position:"Query" name:"HlsSwitch"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (r DescribeLiveStreamRoomUserNumberRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamRoomUserNumberResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamRoomUserNumberRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRoomUserNumber", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamRoomUserNumberResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamRoomUserNumberResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamRoomUserNumberResponse struct {
	RequestId       string
	TotalUserNumber int64
	OnlineUserInfo  DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfoList
}

type DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfo struct {
	StreamUrl  string
	UserNumber int64
	Time       string
}

type DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfoList []DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfo

func (list *DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfo)
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
