package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamRoomUserNumberRequest struct {
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

func (req *DescribeLiveStreamRoomUserNumberRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamRoomUserNumberResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRoomUserNumber", "", "")
	resp = &DescribeLiveStreamRoomUserNumberResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamRoomUserNumberResponse struct {
	responses.BaseResponse
	RequestId       common.String
	TotalUserNumber common.Long
	OnlineUserInfo  DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfoList
}

type DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfo struct {
	StreamUrl  common.String
	UserNumber common.Long
	Time       common.String
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
