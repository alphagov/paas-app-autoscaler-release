package collector_test

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"code.cloudfoundry.org/app-autoscaler/src/autoscaler/db"
	"code.cloudfoundry.org/app-autoscaler/src/autoscaler/fakes"
	"code.cloudfoundry.org/app-autoscaler/src/autoscaler/metricsserver/collector"
	"code.cloudfoundry.org/app-autoscaler/src/autoscaler/metricsserver/config"
	"code.cloudfoundry.org/app-autoscaler/src/autoscaler/models"

	"code.cloudfoundry.org/lager"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/ginkgomon"

	"testing"
)

var (
	serverProcess ifrit.Process
	serverUrl     *url.URL
)

const (
	TestCollectInterval = 1 * time.Second
	TestRefreshInterval = 2 * time.Second
	TestSaveInterval    = 2 * time.Second
)

func TestCollector(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Collector Suite")
}

var _ = BeforeSuite(func() {

	port := 1111 + GinkgoParallelProcess()
	serverConf := &collector.ServerConfig{
		Port:      port,
		NodeAddrs: []string{fmt.Sprintf("%s:%d", "localhost", port)},
		NodeIndex: 0,
	}

	conf := &config.Config{}

	queryFunc := func(appID string, instanceIndex int, name string, start, end int64, order db.OrderType) ([]*models.AppInstanceMetric, error) {
		return nil, nil
	}

	httpStatusCollector := &fakes.FakeHTTPStatusCollector{}
	httpServer, err := collector.NewServer(lager.NewLogger("test"), serverConf, conf, queryFunc, httpStatusCollector)
	Expect(err).NotTo(HaveOccurred())

	serverUrl, err = url.Parse("http://127.0.0.1:" + strconv.Itoa(port))
	Expect(err).ToNot(HaveOccurred())

	serverProcess = ginkgomon.Invoke(httpServer)
})

var _ = AfterSuite(func() {
	ginkgomon.Interrupt(serverProcess)
})
