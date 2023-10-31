package pulsar

import (
	"github.com/apache/pulsar-client-go/pulsar"
	pulsarLog "github.com/apache/pulsar-client-go/pulsar/log"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type PulsarConfig struct {
	Hosts []string `json:"hosts" yaml:"Hosts"`
}

func (pc PulsarConfig) BuildClient() (pulsar.Client, error) {
	logger := logrus.StandardLogger()
	logger.Level = logrus.WarnLevel
	addr := make([]string, 0, len(pc.Hosts))
	for _, v := range pc.Hosts {
		addr = append(addr, "pulsar://"+v)
	}
	url := strings.Join(addr, ",")
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               url,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
		Logger:            pulsarLog.NewLoggerWithLogrus(logger),
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

type Topic struct {
	Tenant    string
	Namespace string
	Topic     string
}

func (t Topic) BuildTopic() string {
	return "persistent://" + t.Tenant + "/" + t.Namespace + "/" + t.Topic
}

const (
	PublicTenant          = "public"
	GexNamespace          = "trade"
	MatchSourceTopic      = "match_source"
	MatchResultTopic      = "match_result"
	MatchResultAccountSub = "MatchResultAccountSub"
	MatchSourceSub        = "match_source_sub"
	MatchResultOrderSub   = "MatchResultOrderSub"
	MatchResultKlineSub   = "MatchResultKlineSub"
	MatchResultTickerSub  = "MatchResultTickerSub"
	MatchResultMatchSub   = "MatchResultMatchSub"
)
