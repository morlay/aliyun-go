package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PlayInfoRequest struct {
	PlayDomain           string `position:"Query" name:"PlayDomain"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	Formats              string `position:"Query" name:"Formats"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	HlsUriToken          string `position:"Query" name:"HlsUriToken"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	Rand                 string `position:"Query" name:"Rand"`
	AuthTimeout          int64  `position:"Query" name:"AuthTimeout"`
	AuthInfo             string `position:"Query" name:"AuthInfo"`
}

func (r PlayInfoRequest) Invoke(client *sdk.Client) (response *PlayInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		PlayInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "PlayInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		PlayInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PlayInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type PlayInfoResponse struct {
	RequestId         string
	PlayInfoList      PlayInfoPlayInfoList
	NotFoundCDNDomain PlayInfoNotFoundCDNDomainList
}

type PlayInfoPlayInfo struct {
	Url          string
	Duration     string
	Size         string
	Width        string
	Height       string
	Bitrate      string
	Fps          string
	Format       string
	Definition   string
	Encryption   string
	Rand         string
	Plaintext    string
	Complexity   string
	ActivityName string
}

type PlayInfoPlayInfoList []PlayInfoPlayInfo

func (list *PlayInfoPlayInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]PlayInfoPlayInfo)
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

type PlayInfoNotFoundCDNDomainList []string

func (list *PlayInfoNotFoundCDNDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
