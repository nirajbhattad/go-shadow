package kafka

// import (
// 	"fmt"
// 	"strconv"

// 	"github.aexp.com/amex-eng/eventcodes"
// 	"github.aexp.com/amex-eng/go-fish/protocols/tcp"
// 	"github.aexp.com/amex-eng/kenny"
// 	"github.aexp.com/amex-eng/tcp-proxy/client"
// 	"github.aexp.com/amex-eng/tcp-proxy/utils"
// 	"golang.org/x/sys/unix"
// )

// var pingListener *tcp.Server

// // loadClients provides definition for all binding client sockets for inbound connection.
// func loadClients() {
// 	cl := GetAllClientCfg()
// 	//Map profiles and port's
// 	cl.Range(func(key, value interface{}) bool {
// 		var c *client.Client
// 		c, ok := value.(*client.Client)
// 		if !ok {
// 			logger.Loggings(logger.Error, nil, kenny.New(eventcodes.InvalidConfigSyntaxError+"unable to cast client."), nil)
// 			return true
// 		}
// 		if c.TargetProfileIdentifier() != "" {
// 			cc, ok := GetClientCfg(strconv.Itoa(c.Port()) + c.TargetPrefix() + c.TargetProfileIdentifier())
// 			if !ok {
// 				logger.Loggings(logger.Info, nil, kenny.New(eventcodes.GeneralInfo+" throttling/pass though configuration enabled for client["+c.Name()+"] but no mapping profile["+c.TargetProfileIdentifier()+"found."), nil)
// 				return true
// 			}
// 			cc.SetParentPortAndProfile(strconv.Itoa(c.Port()) + c.Prefix() + c.CprofileClientID())
// 			logger.Loggings(logger.Info, nil, kenny.New(eventcodes.GeneralInfo+"parent profile for client["+cc.Name()+"] is ["+c.CprofileClientID()+strconv.Itoa(c.Port())+"]"), nil)
// 		}
// 		return true
// 	})

// 	cl.Range(func(key, value interface{}) bool {
// 		var c *client.Client
// 		c, ok := value.(*client.Client)
// 		if !ok {
// 			logger.Loggings(logger.Error, nil, kenny.New(eventcodes.InvalidConfigSyntaxError+"unable to cast client."), nil)
// 			return true
// 		}
// 		go tcpIn{netListener: tcpIn{}, netDialer: tcpIn{}}.HandleConn(c)
// 		return true
// 	})

// 	// Start Ping Port
// 	pingListener = tcp.NewServer(tcp.Config{
// 		Handler:       pingReader,
// 		SocketOptions: []int{unix.SO_REUSEADDR, unix.SO_REUSEPORT},
// 	})
// 	go func() {
// 		logger.Loggings(logger.Info, nil, fmt.Errorf("Starting Ping Port Listener on %s", utils.Cfg.Custom(utils.PingPortKey)), nil)
// 		err := pingListener.Listen("0.0.0.0:" + utils.Cfg.Custom(utils.PingPortKey))
// 		if err != nil && err != tcp.ErrShutdown {
// 			logger.Loggings(logger.Error, nil, fmt.Errorf("%s - Failed to bind Ping Listener on %s - %s", eventcodes.ListenError, utils.Cfg.Custom(utils.PingPortKey), err),
// 				err)
// 		}
// 	}()

// 	logger.Loggings(logger.Info, nil, kenny.New(eventcodes.GeneralInfo+"successfully opened client sockets. Started ping listener."), nil)
// }
