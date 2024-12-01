package etcd

import (
	"github.com/spf13/cast"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/metadata"
	"sync"
	"time"
)

//基于交易对负载均衡
//场景 撮合，订单，行情 分交易对，api 服务对这些交易对建立连接后

//key kline/IKUN_USDT/xxxxx

var (
	_lock = &sync.RWMutex{}
)

const SymbolLB = "symbol_lb"

// 自定义负载均衡，加权轮询
type weightConf struct {
	addr   string
	weight int32
}

// 自定义 Picker
type symbolPicker struct {
	subConns map[string][]balancer.SubConn // 连接列表
	weights  []*weightConf                 // 权重列表
}

func (p *symbolPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	// 如果没有可用的连接，返回错误
	if len(p.subConns) == 0 {
		return balancer.PickResult{}, nil
	}
	md, ok := metadata.FromIncomingContext(info.Ctx)
	if !ok {
		return balancer.PickResult{}, nil
	}
	symbol := md.Get("symbol")[0]
	conns, ok := p.subConns[symbol]
	if !ok || len(conns) == 0 {
		return balancer.PickResult{}, nil
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
		symbol := cast.ToString(addr.Address.Attributes.Value("symbol"))
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
func newWeightBalancerBuilder() balancer.Builder {
	return base.NewBalancerBuilder(SymbolLB, &symbolPickerBuilder{}, base.Config{HealthCheck: true})
}
