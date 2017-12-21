package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveApRadioSsidConfigRequest struct {
	Hidden             int    `position:"Query" name:"Hidden"`
	DynamicVlan        int    `position:"Query" name:"DynamicVlan"`
	Ssid               string `position:"Query" name:"Ssid"`
	Cir                int    `position:"Query" name:"Cir"`
	Mac                string `position:"Query" name:"Mac"`
	Ieee80211w         int    `position:"Query" name:"Ieee.80211.w"`
	Network            int    `position:"Query" name:"Network"`
	Isolate            int    `position:"Query" name:"Isolate"`
	ApAssetId          int64  `position:"Query" name:"ApAssetId"`
	EncKey             string `position:"Query" name:"EncKey"`
	MulticastForward   int    `position:"Query" name:"MulticastForward"`
	Encryption         string `position:"Query" name:"Encryption"`
	Wmm                int    `position:"Query" name:"Wmm"`
	AuthCache          int    `position:"Query" name:"AuthCache"`
	Disabled           int    `position:"Query" name:"Disabled"`
	Id                 int64  `position:"Query" name:"Id"`
	RadioIndex         int    `position:"Query" name:"RadioIndex"`
	IgnoreWeakProbe    int    `position:"Query" name:"IgnoreWeakProbe"`
	Maxassoc           int    `position:"Query" name:"Maxassoc"`
	DisassocLowAck     int    `position:"Query" name:"DisassocLowAck"`
	DisassocWeakRssi   int    `position:"Query" name:"DisassocWeakRssi"`
	SsidLb             int    `position:"Query" name:"SsidLb"`
	MaxInactivity      int    `position:"Query" name:"MaxInactivity"`
	VlanDhcp           int    `position:"Query" name:"VlanDhcp"`
	InstantlyEffective int    `position:"Query" name:"InstantlyEffective"`
	ShortPreamble      int    `position:"Query" name:"ShortPreamble"`
}

func (r SaveApRadioSsidConfigRequest) Invoke(client *sdk.Client) (response *SaveApRadioSsidConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveApRadioSsidConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveApRadioSsidConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveApRadioSsidConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveApRadioSsidConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveApRadioSsidConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}
