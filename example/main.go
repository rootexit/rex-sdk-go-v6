package main

import (
	"context"
	"github.com/rootexit/rex-sdk-go-v6/rex"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexConfig"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexTypes"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

func main() {
	rexConfig := rexConfig.DefaultConfig(os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	logx.Infof("打印一下请求的accessKey :%s", rexConfig.AccessKeyID)
	logx.Infof("打印一下请求的AccessKeySecret :%s", rexConfig.AccessKeySecret)
	logx.Infof("打印一下请求的Endpoint :%s", rexConfig.Endpoint)

	sdk, err := rex.NewSdk(rexConfig)
	if err != nil {
		panic(err)
	}
	//allCodes, err := sdk.QxBaseService.Codes(context.Background(), &rexTypes.CodesReq{})
	//if err != nil {
	//	panic(err)
	//}
	//logx.Infof("%+v", allCodes)

	//queryBucketResult, err := sdk.SasService.QueryBucket(context.Background(), &rexTypes.SasQueryBucketReq{
	//	BucketKey: "default",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//logx.Infof("%+v", queryBucketResult)

	//queryBucketResult, err := sdk.UpsService.UpsTagService.Query(context.Background(), &rexTypes.TagApiFormIdReq{
	//	Id: 2,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//logx.Infof("%+v", queryBucketResult)

	//queryBucketResult, err := sdk.UpsService.IndustryService.QueryList(context.Background(), &rexTypes.IndustryCommonSearchParams{
	//	Page:       1,
	//	PageSize:   5,
	//	OnlyParent: true,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//logx.Infof("%+v", queryBucketResult)

	_, queryAccessTokenResult, err := sdk.TpasService.WechatOffiaccountService.GetAccessToken(context.Background(), &rexTypes.WechatOffiaccountGetAccessTokenReq{
		Key: "default",
	})
	if err != nil {
		panic(err)
	}
	logx.Infof("%+v", queryAccessTokenResult)
}
