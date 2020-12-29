package video

import (
	"mime/multipart"
)

//Video 视频Video
type Video struct {
	ID                  string            `json:"id"`                  //视频ID
	ParentID            interface{}       `json:"parentId"`            //父视频ID
	Playrate            string            `json:"playrate"`            //播放速率，目前为倍数:0.8,1.2,1.4,1.6;缩时:10,30,60,180,单位为秒
	FileName            string            `json:"fileName"`            //文件名
	Thumbnail           string            `json:"thumbnail"`           //默认缩略图
	Cover               string            `json:"cover"`               //默认封面图
	Title               string            `json:"title"`               //标题
	Width               int               `json:"width"`               //源视频文件宽度
	Height              int               `json:"height"`              //源视频文件高度
	Bitrate             int               `json:"bitrate"`             //源文件比特率
	AudioBitrate        int               `json:"audioBitrate"`        //源文件音频码率
	Duration            int               `json:"duration"`            //视频时长，毫秒
	PlayNum             int               `json:"playNum"`             //播放次数
	Size                int               `json:"size"`                //视频大小，KB
	UserID              int               `json:"userId"`              // 用户id
	CategoryID          int               `json:"categoryId"`          //类别id
	Category            interface{}       `json:"category"`            //类别
	Status              int               `json:"status"`              //视频转码状态，1 转码中，2 转码完成，3 转码出错，4 上传失败，5 未知文件类型，6 二次转码中，7 排队中
	Date                string            `json:"date"`                //视频上传日期
	UploadTime          int               `json:"uploadTime"`          //上传时间
	Tag                 string            `json:"tag"`                 //视频标签
	Description         string            `json:"description"`         //描述
	ContentID           interface{}       `json:"contentId"`           //content id，用于视频加解密
	CustomID            interface{}       `json:"customId"`            //custom id，用于视频加解密
	SourceContentID     interface{}       `json:"sourceContentId"`     //视频源本身为加密时所使用的key之一
	SourceCustomID      interface{}       `json:"sourceCustomId"`      //视频源本身为加密时所使用的key之一
	TranscodedVideoList []TranscodedVideo `json:"transcodedVideoList"` //转码后的视频地址列表
	SubPlayrateList     []SubPlayrate     `json:"subPlayrateList"`     //子视频列表
	PlayCode            PlayCode          `json:"playCode"`            // 播放码
	Type                int               `json:"type"`                //源文件类型，0 普通文件，1 hls文件，3 直播录制，4 快编，5 直播录制二次转码，6 快编二次转码，7 视频叠加，8 视频叠加二次转码
	URI                 interface{}       `json:"uri"`                 //源文件uri
	FastEditor          int               `json:"fastEditor"`          //是否快编
	VideoType           int               `json:"videoType"`           //视频类型 0 普通，1 VR
	Audio               interface{}       `json:"audio"`               //VR转出的音频地址
	AudioStatus         interface{}       `json:"audioStatus"`         //VR转出的音频转码状态
	AudioTaskID         interface{}       `json:"audioTaskId"`         //VR转出的音频转码任务id
	Subtitle            string            `json:"subtitle"`            //字幕
	SubVideoType        int               `json:"subVideoType"`        //如果该视频为子视频，则表示子视频类型，0 倍数转码，1 缩时转码
	ShareStatus         int               `json:"shareStatus"`         //分享状态，目前无效
	ShareToUserID       interface{}       `json:"shareToUserId"`       //分享给的用户id，目前无效
	ShareCategoryID     int               `json:"shareCategoryId"`     //分享的视频类别，目前无效
	LogoGroupID         interface{}       `json:"logoGroupId"`         //转码时所用的台标组id
}

//PlayCode 播放码
type PlayCode struct {
	PreviewURL    string `json:"previewUrl"`    //预览地址
	FlashURL      string `json:"flashUrl"`      //flash播放地址
	HTMLEmbedCode string `json:"htmlEmbedCode"` //html嵌入式代码
}

//SubPlayrate 子视频列表
type SubPlayrate struct {
	VideoID      string `json:"videoId"`      //子视频id
	Playrate     string `json:"playrate"`     //子视频播放速率，含义同
	SubVideoType int    `json:"subVideoType"` //子视频类型，0 倍数转码，1 缩时转码
}

//TranscodedVideo 视频地址
type TranscodedVideo struct {
	Name           string      `json:"name"`           //名称
	HLSURL         string      `json:"hlsUrl"`         //转码后hls地址
	HLSSize        int         `json:"hlsSize"`        //转码后hls文件大小，Byte
	Mp4URL         string      `json:"mp4Url"`         //转码后mp4地址
	Mp4Size        int         `json:"mp4Size"`        //转码后mp4文件大小，Byte
	FlvURL         string      `json:"flvUrl"`         //转码后flv地址
	FlvSize        int         `json:"flvSize"`        //转码后flv文件大小，Byte
	TsURL          string      `json:"tsUrl"`          //转码后ts地址
	TsSize         int         `json:"tsSize"`         //转码后ts文件大小，Byte
	VideoWidth     int         `json:"videoWidth"`     //转码后视频宽度
	VideoHeight    int         `json:"videoHeight"`    //转码后视频高度
	VideoBitrate   int         `json:"videoBitrate"`   //视频码率
	AudioBitrate   int         `json:"audioBitrate"`   //转码后音频码率
	DefaultVideo   int         `json:"defaultVideo"`   //是否为默认播放，1 是， 0 否
	TaskID         string      `json:"taskId"`         //视频转码任务id
	Status         int         `json:"status"`         //转码任务状态
	Idx            int         `json:"idx"`            //同一个文件的不同输出地址索引
	Type           int         `json:"type"`           //转码后的文件类型，0 普通，1 试看片段
	HLSEncryptType int         `json:"hlsEncryptType"` //转码后hls加密类型，0 不加密，6 AES加密，7 当虹加密，目前AES加密只对hls起作用
	Mp4EncryptType int         `json:"mp4EncryptType"` //转码后mp4加密类型，0 不加密，7当虹加密
	FlvEncryptType int         `json:"flvEncryptType"` //转码后flv加密类型，0 不加密，7 当虹加密
	TsEncryptType  int         `json:"tsEncryptType"`  //转码后ts加密类型，0 不加密，7 当虹加密
	ContentID      string      `json:"contentId"`      //当虹加密所用的key之一
	CustomID       string      `json:"customId"`       //当虹加密所用的key之一
	AESKey         interface{} `json:"aesKey"`         //AES加密所用的key
}

//VideoListRespParam 视频列表返回
type VideoListRespParam struct {
	Code    int     `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string  `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool    `json:"success"` // true 或 false
	Result  []Video `json:"result"`  // 字符串，返回的结果
}

//VideoInfoRespParam 视频信息返回
type VideoInfoRespParam struct {
	Code    int    `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool   `json:"success"` // true 或 false
	Result  Video  `json:"result"`  // 字符串，返回的结果
}

//SearchVideoParam 视频列表参数
type SearchVideoParam struct {
	Start      string `json:"start"`      //开始位置
	Number     string `json:"number"`     //数量
	Sort       string `json:"sort"`       //结果按该字段排序，只能为title, size, category, uploadTime, status, tag
	Order      string `json:"order"`      //asc 顺序，desc 倒序
	SearchType string `json:"searchType"` //查找类型为title，tag，all,
	Keyword    string `json:"keyword"`    //查找的关键字
	CategoryID int    `json:"categoryId"` //视频类别id
	VideoType  int    `json:"videoType"`  // 0:普通视频，1:VR，不传该值的话默认为0
	Status     int    `json:"status"`     //视频状态，不传该值的话为返回所有状态的视频
}

// UploadVideoParams 视频上传参数
type UploadVideoParams struct {
	Token                    string         `json:"token"` //令牌
	File                     multipart.File `json:"file"`
	IsEncrypt                string         `json:"isEncrypt"`                //是否加密 0不加密，1加密。默认不加密
	CategoryID               string         `json:"categoryId"`               //视频分类id
	VideoType                int            `json:"videoType"`                // 0:普通视频，1:VR，不传该值的话默认为0
	IsTranscoder             string         `json:"isTranscoder"`             //1:视频上传后编码，0:视频上传后不编码。默认转码
	TranscodeTemplateGroupID int            `json:"transcodeTemplateGroupId"` //转码模板组id，如该值为空，则转码时使用默认转码模板组
	CapturePoints            string         `json:"capturePoints"`            //缩略图/封面图 的截图时间点（秒）最多8张
	IsExtractFrame           int            `json:"isExtractFrame"`           //视频上传成功后是否立即抽帧： 0 否，1 是
	FrameInterval            int            `json:"frameInterval"`            //抽帧的时间间隔（秒），如果抽帧，则需要设置该参数值
	Tag                      string         `json:"tag"`                      //上传视频后存储的视频标签，用于检索
	URL                      string         `json:"url"`
}

// UploadInitMultipart 视频上传参数
type UploadInitMultipart struct {
	Token    string `json:"token"`    //令牌
	Type     string `json:"type"`     //文件类型FileName
	FileMD5  string `json:"fileMD5"`  //文件的MD5值，若不提供文件的MD5值，则不会校验合并后的文件MD5值
	FileSize int    `json:"fileSize"` //文件大小
	FileName string `json:"fileName"` //文件名称
}

type UploadInitMultipartRespParam struct {
	Code    int    `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool   `json:"success"` // true 或 false
	Result  string `json:"result"`  // 字符串，返回的结果
}

type UploadMultipart struct {
	Token           string `json:"token"`           //令牌
	UploadID        string `json:"uploadId"`        // 上传唯一标识
	ChunkMD5        string `json:"chunkMD5"`        //分块文件的MD5值
	ChunkIndex      int    `json:"chunkIndex"`      //分块文件序列
	ChunkSize       int    `json:"chunkSize"`       //分块文件大小
	FileInputStream int64  `json:"fileInputStream"` //文件流

}

type SecondTranscode struct {
	VideoIds           string `json:"videoIds"`           //视频id，多个id以逗号隔开，
	TranscodeTemplates string `json:"transcodeTemplates"` //视频id，多个id以逗号隔开，
	VideoType          int    `json:"videoType"`          //视频转码类型，0 普通视频转码，1 VR转码
	Replace            bool   `json:"replace"`            //0 另存为，1 替换
	Encrypt            bool   `json:"encrypt"`            //是否加密，0 非加密，1 加密

}

type TranscodeTemplates struct {
	Name           string `json:"name"`         //模板名字
	VideoWidth     int    `json:"videoWidth"`   //视频宽度
	VideoHeight    int    `json:"videoHeight"`  //视频高度
	VideoBitrate   int    `json:"videoBitrate"` //视频比特率，例如1000为1000kbps
	AudioBitrate   int    `json:"audioBitrate"` //音频比特率，例如64000为64kbps
	HLS            bool   `json:"hls"`
	Mp4            bool   `json:"mp4"`
	Flv            bool   `json:"flv"`
	Ts             bool   `json:"ts"`
	DefaultVideo   int    `json:"defaultVideo"` //1 默认播放模板，0 非默认
	HLSEncryptType int    `json:"hlsEncryptType"`
	Mp4EncryptType int    `json:"mp4EncryptType"`
	FlvEncryptType int    `json:"flvEncryptType"`
	TsEncryptType  int    `json:"tsEncryptType"`
}

type SecondTranscodeV2 struct {
	VideoIds                 string `json:"videoIds"`                 //视频id，多个id以逗号隔开，
	TranscodeTemplates       string `json:"transcodeTemplates"`       //视频id，多个id以逗号隔开，
	VideoType                int    `json:"videoType"`                //视频转码类型，0 普通视频转码，1 VR转码
	TranscodeTemplateGroupID int    `json:"transcodeTemplateGroupId"` //视频转码类型，0 普通视频转码，1 VR转码
	TrailorDuration          int    `json:"trailorDuration"`          //视频转码类型，0 普通视频转码，1 VR转码
	Replace                  bool   `json:"replace"`                  //0 另存为，1 替换
	Trailor                  bool   `json:"trailor"`                  //是否加密，0 非加密，1 加密

}
