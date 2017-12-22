package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPlayInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	StreamType           string `position:"Query" name:"StreamType"`
	Formats              string `position:"Query" name:"Formats"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Channel              string `position:"Query" name:"Channel"`
	VideoId              string `position:"Query" name:"VideoId"`
	PlayerVersion        string `position:"Query" name:"PlayerVersion"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Rand                 string `position:"Query" name:"Rand"`
	ReAuthInfo           string `position:"Query" name:"ReAuthInfo"`
	Definition           string `position:"Query" name:"Definition"`
	AuthTimeout          int64  `position:"Query" name:"AuthTimeout"`
	AuthInfo             string `position:"Query" name:"AuthInfo"`
}

func (req *GetPlayInfoRequest) Invoke(client *sdk.Client) (resp *GetPlayInfoResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "GetPlayInfo", "vod", "")
	resp = &GetPlayInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPlayInfoResponse struct {
	responses.BaseResponse
	RequestId    string
	PlayInfoList GetPlayInfoPlayInfoList
	VideoBase    GetPlayInfoVideoBase
}

type GetPlayInfoPlayInfo struct {
	Width      int64
	Height     int64
	Size       int64
	PlayURL    string
	Bitrate    string
	Definition string
	Duration   string
	Format     string
	Fps        string
	Encrypt    int64
	Plaintext  string
	Complexity string
	StreamType string
	Rand       string
	JobId      string
}

type GetPlayInfoVideoBase struct {
	CoverURL     string
	Duration     string
	Status       string
	Title        string
	VideoId      string
	MediaType    string
	CreationTime string
}

type GetPlayInfoPlayInfoList []GetPlayInfoPlayInfo

func (list *GetPlayInfoPlayInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPlayInfoPlayInfo)
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
