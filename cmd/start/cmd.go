package start

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Duke1616/alertmanager-wechat-robot/conf"
	"github.com/Duke1616/alertmanager-wechat-robot/protocol"
	app "github.com/Duke1616/alertmanager-wechat-robot/register"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"

	// 注册所有服务
	_ "github.com/Duke1616/alertmanager-wechat-robot/apps/all"
)

// Cmd startCmd represents the start command
var Cmd = &cobra.Command{
	Use:   "start",
	Short: "alertmanager-wechat-robot API服务",
	Long:  "alertmanager-wechat-robot API服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		conf := conf.C()
		// 启动服务
		ch := make(chan os.Signal, 1)
		defer close(ch)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

		// 初始化服务
		svr, err := newService(conf)
		if err != nil {
			return err
		}

		// 等待信号处理
		go svr.waitSign(ch)

		// 启动服务
		if err := svr.start(); err != nil {
			if !strings.Contains(err.Error(), "http: Server closed") {
				return err
			}
		}

		return nil
	},
}

func newService(cnf *conf.Config) (*service, error) {
	http := protocol.NewHTTPService()
	grpc := protocol.NewGRPCService()
	svr := &service{
		http: http,
		grpc: grpc,
		log:  zap.L().Named("CLI"),
	}

	return svr, nil
}

type service struct {
	http *protocol.HTTPService
	grpc *protocol.GRPCService

	log logger.Logger
}

func (s *service) start() error {
	s.log.Infof("loaded grpc app: %s", app.LoadedGrpcApp())
	s.log.Infof("loaded http app: %s", app.LoadedRestfulApp())
	s.log.Infof("loaded inappnal app: %s", app.LoadedInternalApp())

	go s.grpc.Start()
	return s.http.Start()
}

func (s *service) waitSign(sign chan os.Signal) {
	for sg := range sign {
		switch v := sg.(type) {
		default:

			s.log.Infof("receive signal '%v', start graceful shutdown", v.String())
			//
			//if err := s.grpc.Stop(); err != nil {
			//	s.log.Errorf("grpc graceful shutdown err: %s, force exit", err)
			//} else {
			//	s.log.Info("grpc service stop complete")
			//}

			if err := s.http.Stop(); err != nil {
				s.log.Errorf("http graceful shutdown err: %s, force exit", err)
			} else {
				s.log.Infof("http service stop complete")
			}
			return
		}
	}
}
