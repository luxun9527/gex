package etcd

import (
	"github.com/spf13/cast"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

//基于交易对负载均衡
//场景 撮合，订单，行情 分交易对，api 服务对这些交易对建立连接后

//key klineRpc/IKUN_USDT/xxxxx

func init() {
	balancer.Register(newSymbolBalancerBuilder())
}

var (
	NotAvailableConn = status.Error(codes.Unavailable, "no available connection")
)

const SymbolLB = "symbol_lb"

// 自定义 Picker
type symbolPicker struct {
	subConns map[string][]balancer.SubConn // 连接列表
}

func (p *symbolPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	// 如果没有可用的连接，返回错误
	if len(p.subConns) == 0 {
		return balancer.PickResult{}, NotAvailableConn
	}
	md, ok := metadata.FromIncomingContext(info.Ctx)
	if !ok {
		return balancer.PickResult{}, NotAvailableConn
	}
	symbol := md.Get("symbol")[0]
	conns, ok := p.subConns[symbol]
	if !ok || len(conns) == 0 {
		return balancer.PickResult{}, NotAvailableConn
	}
	index := time.Now().UnixNano() % int64(len(conns))
	return balancer.PickResult{SubConn: conns[index]}, nil
}

// 负载均衡器构建器
type symbolPickerBuilder struct {
	weightConfig map[string]int32
}

func (wp *symbolPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}
	var p = map[string][]balancer.SubConn{}
	for sc, addr := range info.ReadySCs {
		symbolData, ok := addr.Address.Metadata.(map[string]interface{})
		if !ok {
			continue
		}
		symbol := cast.ToString(symbolData["symbol"])
		if symbol == "" {
			continue
		}
		conns, ok := p[symbol]
		if !ok {
			conns = make([]balancer.SubConn, 0, 1)
		}
		conns = append(conns, sc)
		p[symbol] = conns
	}

	return &symbolPicker{
		subConns: p,
	}
}

// 自定义负载均衡
func newSymbolBalancerBuilder() balancer.Builder {
	return base.NewBalancerBuilder(SymbolLB, &symbolPickerBuilder{}, base.Config{HealthCheck: true})
}
