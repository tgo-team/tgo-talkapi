package tcp

import (
	"github.com/tgo-team/tgo-core/tgo"
	"net"
	"os"
	"runtime"
	"strings"
)

func init() {
	tgo.RegistryServer(func(context *tgo.Context) tgo.Server {
		return NewServer(context)

	})
}

type Server struct {
	tcpListener        net.Listener
	exitChan           chan int
	waitGroup          tgo.WaitGroupWrapper
	acceptPacketChan   chan *tgo.PacketContext
	acceptConnChan     chan tgo.Conn
	acceptConnExitChan chan tgo.Conn
	storage            tgo.Storage
	opts               *tgo.Options
	pro                tgo.Protocol
	ctx                *tgo.Context
}

func NewServer(ctx *tgo.Context) *Server {
	s := &Server{
		exitChan:           make(chan int, 0),
		acceptConnExitChan: ctx.TGO.AcceptConnExitChan,
		acceptPacketChan:   ctx.TGO.AcceptPacketChan,
		acceptConnChan:     ctx.TGO.AcceptConnChan,
		opts:               ctx.TGO.GetOpts(),
		pro:                ctx.TGO.GetOpts().Pro,
		ctx:                ctx,
	}
	var err error
	s.tcpListener, err = net.Listen("tcp", s.opts.TCPAddress)
	if err != nil {
		s.Fatal("listen (%s) failed - %s", s.opts.TCPAddress, err)
		os.Exit(1)
	}
	return s
}

func (s *Server) GetOpts() *tgo.Options {
	return s.opts
}

func (s *Server) Start() error {
	s.waitGroup.Wrap(s.connLoop)
	return nil
}

func (s *Server) Stop() error {
	if s.tcpListener != nil {
		err := s.tcpListener.Close()
		if err != nil {
			return err
		}
	}
	close(s.exitChan)
	s.waitGroup.Wait()
	s.Info("Server -> 退出")
	return nil
}

func (s *Server) connLoop() {
	s.Info("开始监听 -> %s", s.tcpListener.Addr())
	for {
		select {
		case <-s.exitChan:
			goto exit
		default:
			cn, err := s.tcpListener.Accept()
			if err != nil {
				if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
					s.Error("temporary Accept() failure - %s", err)
					runtime.Gosched()
					continue
				}
				if !strings.Contains(err.Error(), "use of closed network connection") {
					s.Error("listener.Accept() - %s", err)
				}
				break
			}
			println(cn.RemoteAddr().String())
			s.Debug("客户端[%s] -> 请求连接", cn.RemoteAddr())
			s.waitGroup.Wrap(func() {
				s.generateConn(cn)
			})
		}
	}
exit:
	s.Debug("退出监听")
}

func (s *Server) generateConn(conn net.Conn) {
	cn := NewConn(conn, NewConnChan(s.acceptPacketChan, s.acceptConnExitChan), s.ctx)
	cn.StartIOLoop()
}

func (s *Server) RealTCPAddr() *net.TCPAddr {
	return s.tcpListener.Addr().(*net.TCPAddr)
}
