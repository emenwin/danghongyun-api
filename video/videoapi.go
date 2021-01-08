package video

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/emenwin/danghongyun-api/auth"
	"github.com/emenwin/danghongyun-api/client"
)

const (
	//LiveRestAPIURL rest url
	LiveRestAPIURL = "http://api.danghongyun.com/rest"
	//Version version
	Version = "2.0"
)

// DHLiveManager 提供了对直播进行管理的操作
type DHLiveManager struct {
	Client      *client.Client
	Credentials *auth.Credentials
}

// NewLiveManager 用来构建一个新的直播管理对象
func NewLiveManager(cred *auth.Credentials) *DHLiveManager {

	return NewLiveManagerEx(cred, nil)
}

// NewLiveManagerEx 用来构建一个新的直播管理对象
func NewLiveManagerEx(cred *auth.Credentials, clt *client.Client) *DHLiveManager {

	if clt == nil {
		clt = &client.DefaultClient
	}

	return &DHLiveManager{
		Client:      &client.DefaultClient,
		Credentials: cred,
	}
}

//GetUserInfo 获取用户信息
func (ctl *DHLiveManager) GetUserInfo() (*UserInfoRespParam, error) {
	params := ctl.Credentials.NewSignParams("getUser", Version, "")
	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte UserInfoRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//GetToken 获取token
func (ctl *DHLiveManager) GetToken() (*UserInfoRespParam, error) {
	params := ctl.Credentials.NewSignParams("getToken", Version, "")
	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte UserInfoRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//GetVideos 视频列表
func (ctl *DHLiveManager) GetVideos(searchParams SearchVideoParam) (*VideoListRespParam, error) {
	params := ctl.Credentials.NewSignParams("queryVideo", Version, "")
	params["start"] = searchParams.Start
	params["sort"] = searchParams.Sort
	params["number"] = searchParams.Number
	params["order"] = searchParams.Order
	if searchParams.SearchType != "" {
		params["searchType"] = searchParams.SearchType
	}
	if searchParams.Keyword != "" {
		params["keyword"] = searchParams.Keyword
	}

	if searchParams.CategoryID != -1 {
		params["categoryId"] = strconv.Itoa(searchParams.CategoryID)
	}

	if searchParams.VideoType != -1 {
		params["videoType"] = strconv.Itoa(searchParams.VideoType)
	}

	if searchParams.Status != -1 {
		params["status"] = strconv.Itoa(searchParams.Status)
	}

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte VideoListRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//EditVideo 修改视频信息
func (ctl *DHLiveManager) EditVideo(param Video) (*VideotypeInfoRespParam, error) {

	params := ctl.Credentials.NewSignParams("updateVideoInfo", Version, "")
	params["videoId"] = param.ID
	params["fileName"] = param.FileName
	params["thumbnail"] = param.Thumbnail
	params["cover"] = param.Cover
	params["title"] = param.Title
	params["playNum"] = strconv.Itoa(param.PlayNum)
	params["categoryId"] = strconv.Itoa(param.CategoryID)
	params["status"] = strconv.Itoa(param.Status)
	params["tag"] = param.Tag
	params["description"] = param.Description
	params["width"] = strconv.Itoa(param.Width)
	params["height"] = strconv.Itoa(param.Height)
	params["bitrate"] = strconv.Itoa(param.Bitrate)
	params["audioBitrate"] = strconv.Itoa(param.AudioBitrate)
	params["duration"] = strconv.Itoa(param.Duration)

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte VideotypeInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//DeleteVideo 删除视频
//分类id，多个用逗号隔开，如"1,2,3"
func (ctl *DHLiveManager) DeleteVideo(videoIds string) (*VideotypeInfoRespParam, error) {

	params := ctl.Credentials.NewSignParams("deleteVideo", Version, "")
	params["videoIds"] = videoIds

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte VideotypeInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//GetVideoInfo 视频信息
func (ctl *DHLiveManager) GetVideoInfo(videoID string) (*VideoInfoRespParam, error) {
	params := ctl.Credentials.NewSignParams("getVideoInfo", Version, "")
	params["videoId"] = videoID

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte VideoInfoRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//GetVideoStatus 视频状态
func (ctl *DHLiveManager) GetVideoStatus(videoID string) (*VideoInfoRespParam, error) {
	params := ctl.Credentials.NewSignParams("getVideoStatus", Version, "")
	params["videoId"] = videoID

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte VideoInfoRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//GetVideoCount 视频数量
func (ctl *DHLiveManager) GetVideoCount(searchParams SearchVideoParam) (*VideoInfoRespParam, error) {
	params := ctl.Credentials.NewSignParams("getVideoCount", Version, "")
	params["categoryID"] = strconv.Itoa(searchParams.CategoryID)
	params["videoType"] = strconv.Itoa(searchParams.VideoType)
	params["status"] = strconv.Itoa(searchParams.Status)
	params["keyword"] = searchParams.Keyword

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte VideoInfoRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//GetVideoPlayCode 视频播放码
func (ctl *DHLiveManager) GetVideoPlayCode(videoID string) (*VideoInfoRespParam, error) {
	params := ctl.Credentials.NewSignParams("getVideoPlaycode", Version, "")
	params["videoId"] = videoID

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte VideoInfoRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//GetVideoThumbnailList 视频缩略图列表
func (ctl *DHLiveManager) GetVideoThumbnailList(videoID string) (*VideoListRespParam, error) {
	params := ctl.Credentials.NewSignParams("getThumbnailList", Version, "")
	params["videoId"] = videoID
	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte VideoListRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//GetVideoCoverList 视频封面图列表
func (ctl *DHLiveManager) GetVideoCoverList(videoID string) (*VideoListRespParam, error) {
	params := ctl.Credentials.NewSignParams("getCoverList", Version, "")
	params["videoId"] = videoID
	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte VideoListRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//GetTranscodedVideos 获取转码后的视频点播地址
//视频id, 多个id以逗号隔开
func (ctl *DHLiveManager) GetTranscodedVideos(ids string) (*VideoListRespParam, error) {
	params := ctl.Credentials.NewSignParams("getTranscodedVideos", Version, "")
	params["ids"] = ids
	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte VideoListRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//GetDownloadAddress 获取转码后的视频下载地址
//视频id, 多个id以逗号隔开
//type 0 普通， 1 试看，默认值0
func (ctl *DHLiveManager) GetDownloadAddress(ids string, videotype int) (*VideoListRespParam, error) {
	params := ctl.Credentials.NewSignParams("getDownloadAddress", Version, "")
	params["ids"] = ids
	params["type"] = strconv.Itoa(videotype)
	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte VideoListRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//GetVideocategory 视频分类列表
func (ctl *DHLiveManager) GetVideocategory() (*VideotypeListRespParam, error) {
	params := ctl.Credentials.NewSignParams("getVideoCategoryList", Version, "")
	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte VideotypeListRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//EditVideocategory 修改视频分类
func (ctl *DHLiveManager) EditVideocategory(param Videocategory) (*VideotypeInfoRespParam, error) {

	params := ctl.Credentials.NewSignParams("updateVideoCategory", Version, "")
	params["id"] = strconv.Itoa(param.ID)
	params["category"] = param.Category
	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte VideotypeInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//CreateVideocategory 新增视频分类
func (ctl *DHLiveManager) CreateVideocategory(param Videocategory) (*VideotypeInfoRespParam, error) {

	params := ctl.Credentials.NewSignParams("addVideoCategory", Version, "")
	params["category"] = param.Category
	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte VideotypeInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//DeleteVideocategory 删除视频分类
//分类id，多个用逗号隔开，如"1,2,3"
func (ctl *DHLiveManager) DeleteVideocategory(ids string) (*VideotypeInfoRespParam, error) {

	params := ctl.Credentials.NewSignParams("deleteVideoCategory", Version, "")
	params["ids"] = ids

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte VideotypeInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//GetTemplateGroupList 获取转码模板组列表
func (ctl *DHLiveManager) GetTemplateGroupList(templateGroupType int) (*TemplateGroupListRespParam, error) {
	params := ctl.Credentials.NewSignParams("templateGroupList", Version, "")
	params["type"] = strconv.Itoa(templateGroupType)
	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte TemplateGroupListRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//CreateOrEditTemplateGroup 新建 修改转码模板组
func (ctl *DHLiveManager) CreateOrEditTemplateGroup(param TemplateGroup) (*TemplateGroupInfoRespParam, error) {
	params := ctl.Credentials.NewSignParams("templateGroup", Version, "")
	if param.ID != 0 {
		params["id"] = strconv.Itoa(param.ID)
	}

	params["transcodeTemplates"] = param.TranscodeTemplates
	params["name"] = param.Name
	params["type"] = strconv.Itoa(param.Type)
	params["trailor"] = strconv.FormatBool(param.Trailor)
	params["trailorDuration"] = strconv.Itoa(param.TrailorDuration)
	params["logoGroupId"] = strconv.Itoa(param.LogoGroupID)
	fmt.Print(params)
	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte TemplateGroupInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//DeleteTemplateGroup 删除码模板组
//分类id，多个用逗号隔开，如"1,2,3"
func (ctl *DHLiveManager) DeleteTemplateGroup(ids string) (*TemplateGroupInfoRespParam, error) {

	params := ctl.Credentials.NewSignParams("deleteTemplateGroup", Version, "")
	params["ids"] = ids

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte TemplateGroupInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//GetLogoGroupList 获取台标组列表
func (ctl *DHLiveManager) GetLogoGroupList() (*LogoGroupListRespParam, error) {
	params := ctl.Credentials.NewSignParams("getLogoGroupList", Version, "")

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte LogoGroupListRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//CreateOrEditLogoGroup 新建 修改台标组
func (ctl *DHLiveManager) CreateOrEditLogoGroup(param LogoGroup) (*LogoGroupInfoRespParam, error) {
	params := ctl.Credentials.NewSignParams("logoGroup", Version, "")
	if param.ID != 0 {
		params["id"] = strconv.Itoa(param.ID)
	}

	params["name"] = param.Name
	params["logos"] = param.Logos

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte LogoGroupInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//DeleteLogoGroup 删除台标组
//台标组id，多个用逗号隔开，如“1,2,3”
func (ctl *DHLiveManager) DeleteLogoGroup(ids string) (*LogoGroupInfoRespParam, error) {
	params := ctl.Credentials.NewSignParams("deleteLogoGroup", Version, "")
	params["ids"] = ids

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte LogoGroupInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//SecondTranscode 二次转码
func (ctl *DHLiveManager) SecondTranscode(param SecondTranscode) (*LogoGroupInfoRespParam, error) {
	params := ctl.Credentials.NewSignParams("secondTranscode", Version, "")
	params["videoIds"] = param.VideoIds
	params["videoType"] = strconv.Itoa(param.VideoType)
	params["encrypt"] = strconv.FormatBool(param.Encrypt)
	params["replace"] = strconv.FormatBool(param.Replace)
	params["transcodeTemplates"] = param.TranscodeTemplates

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte LogoGroupInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//SecondTranscodeV2 二次转码
func (ctl *DHLiveManager) SecondTranscodeV2(param SecondTranscodeV2) (*LogoGroupInfoRespParam, error) {
	params := ctl.Credentials.NewSignParams("secondTranscodeV2", Version, "")
	params["videoIds"] = param.VideoIds
	params["videoType"] = strconv.Itoa(param.VideoType)
	params["trailorDuration"] = strconv.Itoa(param.TrailorDuration)
	params["transcodeTemplateGroupId"] = strconv.Itoa(param.TranscodeTemplateGroupID)
	params["trailor"] = strconv.FormatBool(param.Trailor)
	params["replace"] = strconv.FormatBool(param.Replace)
	params["transcodeTemplates"] = param.TranscodeTemplates

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte LogoGroupInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//UploadVideo 文件上传
func (ctl *DHLiveManager) UploadVideo() string {

	var buff bytes.Buffer
	writer := multipart.NewWriter(&buff)
	writer.WriteField("token", "b7265bc6b9cdddd71b31949806336b36")
	w, _ := writer.CreateFormFile("file", "http://vjs.zencdn.net/v/oceans.mp4")
	w.Write([]byte("this is a file"))
	writer.Close()
	var client http.Client
	targetUrl := "http://upload.danghongyun.com/multipart/upload?"
	resp, err := client.Post(targetUrl, writer.FormDataContentType(), &buff)
	if err != nil {
		fmt.Println(err)

	}

	defer resp.Body.Close()

	fmt.Println("==========jsons=======", resp)
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(data)
	return string(data)

}

//InitMultipartUpload 多线程上传初始化
func (ctl *DHLiveManager) InitMultipartUpload(param UploadInitMultipart) (*UploadInitMultipartRespParam, error) {
	params := make(map[string]string)
	params["token"] = "b7265bc6b9cdddd71b31949806336b36"
	params["fileName"] = param.FileName
	params["fileSize"] = strconv.Itoa(param.FileSize)
	if param.FileMD5 != "" {
		params["fileMD5"] = param.FileMD5
	}

	params["type"] = param.Type

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := "http://upload.danghongyun.com/multipart/initMultipartUpload?" + queryparam
	var respTempalte UploadInitMultipartRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//MultipartUpload 上传分块文件
func (ctl *DHLiveManager) MultipartUpload(param UploadMultipart) (*UploadInitMultipartRespParam, error) {
	params := make(map[string]string)
	params["token"] = "b7265bc6b9cdddd71b31949806336b36"
	params["uploadId"] = param.UploadID
	params["chunkIndex"] = strconv.Itoa(param.ChunkIndex)
	params["chunkMD5"] = param.ChunkMD5
	params["chunkSize"] = strconv.Itoa(param.ChunkSize)

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := "http://upload.danghongyun.com/multipart/multipartUpload?" + queryparam
	var respTempalte UploadInitMultipartRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//SetCallbackURLl 设置上传回调地址
func (ctl *DHLiveManager) SetCallbackURLl(callbackURL string) (*LogoGroupInfoRespParam, error) {
	params := ctl.Credentials.NewSignParams("setCallbackUrl", Version, "")
	params["callbackUrl"] = callbackURL
	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte LogoGroupInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}
