package core

import (
	"fmt"
	"gin/global"
	"gin/initialize"
	"go.uber.org/zap"
)

type server interface {
	// ListenAndServe
	// 不定义ListenAndServe方法就程序就不阻塞了
	ListenAndServe() error

}

func RunWindowsServer() {

	Router:=initialize.Routers()
	address:=fmt.Sprintf(":%d",global.GVA_CONFIG.System.Addr)
	s:=initServer(address,Router)

	global.GVA_LOG.Info("server run success on ",zap.String("address",address))

	fmt.Printf("欢迎使用 gin 随便玩项目")

	global.GVA_LOG.Error(s.ListenAndServe().Error())

}