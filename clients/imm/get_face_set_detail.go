package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetFaceSetDetailRequest struct {
	requests.RpcRequest
	Marker          string `position:"Query" name:"Marker"`
	Project         string `position:"Query" name:"Project"`
	SetId           string `position:"Query" name:"SetId"`
	ReturnAttribute string `position:"Query" name:"ReturnAttribute"`
}

func (req *GetFaceSetDetailRequest) Invoke(client *sdk.Client) (resp *GetFaceSetDetailResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "GetFaceSetDetail", "imm", "")
	resp = &GetFaceSetDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetFaceSetDetailResponse struct {
	responses.BaseResponse
	RequestId   common.String
	SetId       common.String
	NextMarker  common.String
	FaceDetails GetFaceSetDetailFaceDetailsItemList
}

type GetFaceSetDetailFaceDetailsItem struct {
	FaceId        common.String
	SrcUri        common.String
	PhotoId       common.String
	GroupId       common.String
	UnGroupReason common.String
	FaceRectangle GetFaceSetDetailFaceRectangle
	FaceAttribute GetFaceSetDetailFaceAttribute
}

type GetFaceSetDetailFaceRectangle struct {
	Top    common.Integer
	Left   common.Integer
	Width  common.Integer
	Height common.Integer
}

type GetFaceSetDetailFaceAttribute struct {
	Gender      GetFaceSetDetailGender
	Age         GetFaceSetDetailAge
	Headpose    GetFaceSetDetailHeadpose
	Eyestatus   GetFaceSetDetailEyestatus
	Blur        GetFaceSetDetailBlur
	Facequality GetFaceSetDetailFacequality
}

type GetFaceSetDetailGender struct {
	Value common.String
}

type GetFaceSetDetailAge struct {
	Value common.String
}

type GetFaceSetDetailHeadpose struct {
	Pitch_angle common.Float
	Roll_angle  common.Float
	Yaw_angle   common.Float
}

type GetFaceSetDetailEyestatus struct {
	Left_eye_status  GetFaceSetDetailLeft_eye_status
	Right_eye_status GetFaceSetDetailRight_eye_status
}

type GetFaceSetDetailLeft_eye_status struct {
	Normal_glass_eye_open  common.Float
	No_glass_eye_close     common.Float
	Occlusion              common.Float
	No_glass_eye_open      common.Float
	Normal_glass_eye_close common.Float
	Dark_glasses           common.Float
}

type GetFaceSetDetailRight_eye_status struct {
	Normal_glass_eye_open  common.Float
	No_glass_eye_close     common.Float
	Occlusion              common.Float
	No_glass_eye_open      common.Float
	Normal_glass_eye_close common.Float
	Dark_glasses           common.Float
}

type GetFaceSetDetailBlur struct {
	Blurness GetFaceSetDetailBlurness
}

type GetFaceSetDetailBlurness struct {
	Value     common.Float
	Threshold common.Float
}

type GetFaceSetDetailFacequality struct {
	Value     common.Float
	Threshold common.Float
}

type GetFaceSetDetailFaceDetailsItemList []GetFaceSetDetailFaceDetailsItem

func (list *GetFaceSetDetailFaceDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetFaceSetDetailFaceDetailsItem)
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
