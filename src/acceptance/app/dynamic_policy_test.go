package app_test

import (
	"acceptance"
	. "acceptance/helpers"
	"fmt"
	"math"

	"time"

	cfh "github.com/KevinJCross/cf-test-helpers/v2/helpers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AutoScaler dynamic policy", func() {
	var (
		policy string

		doneChan       chan bool
		doneAcceptChan chan bool
		ticker         *time.Ticker
	)

	JustBeforeEach(func() {
		appName = CreateTestApp(cfg, "dynamic-policy", initialInstanceCount)
		appGUID = GetAppGuid(cfg, appName)
		StartApp(appName, cfg.CfPushTimeoutDuration())
		instanceName = CreatePolicy(cfg, appName, appGUID, policy)
	})

	AfterEach(AppAfterEach)

	Context("when scaling by memoryused", func() {

		Context("There is a scale out and scale in policy", func() {
			var heapToUse int64
			BeforeEach(func() {
				heapToUse = int64(math.Min(float64(cfg.NodeMemoryLimit-20), 200))

				if AppResidentSize+20 >= heapToUse {
					Fail("There is not enough app memory in the app to run this test.\n - app resident size %d\n - app memory limit: %d\n - heap to use: %d", AppResidentSize, cfg.NodeMemoryLimit, int(heapToUse))
				}

				policy = GenerateDynamicScaleOutAndInPolicy(1, 2, "memoryused", AppResidentSize+20, AppResidentSize+20)
				initialInstanceCount = 1
			})

			It("should scale out and then back in.", Label(acceptance.LabelSmokeTests), func() {
				url := fmt.Sprintf("/memory/%d/5", heapToUse)
				By(fmt.Sprintf("Use heap %d mb of heap on app", heapToUse))
				CurlAppInstance(cfg, appName, 0, url)
				By("wait for scale to 2")
				WaitForNInstancesRunning(appGUID, 2, 5*time.Minute)

				By("Drop memory used by app")
				CurlAppInstance(cfg, appName, 0, "/memory/close")

				By("Wait for scale to minimum instances")
				WaitForNInstancesRunning(appGUID, 1, 5*time.Minute)
			})
		})
	})

	Context("when scaling by memoryutil", func() {

		Context("when memoryutil", func() {
			BeforeEach(func() {
				//current app resident size is 66mb so 66/128mb is 55%
				policy = GenerateDynamicScaleOutAndInPolicy(1, 2, "memoryutil", 58, 65)
				initialInstanceCount = 1
			})

			It("should scale out and back in", func() {
				By("use 90% of memory in app")
				heapToUse := int(float32(cfg.NodeMemoryLimit) * 0.8)
				CurlAppInstance(cfg, appName, 0, fmt.Sprintf("/memory/%d/5", heapToUse))
				By("Wait for scale to 2 instances")
				WaitForNInstancesRunning(appGUID, 2, 5*time.Minute)

				By("drop memory used")
				CurlAppInstance(cfg, appName, 0, "/memory/close")

				By("Wait for scale down to 1 instance")
				WaitForNInstancesRunning(appGUID, 1, 5*time.Minute)
			})
		})
	})

	Context("when scaling by responsetime", func() {

		JustBeforeEach(func() {
			doneChan = make(chan bool)
			doneAcceptChan = make(chan bool)
		})

		AfterEach(func() {
			close(doneChan)
			Eventually(doneAcceptChan, 10*time.Second).Should(Receive())
		})

		Context("when responsetime is greater than scaling out threshold", func() {

			BeforeEach(func() {
				policy = GenerateDynamicScaleOutPolicy(1, 2, "responsetime", 2000)
				initialInstanceCount = 1
			})

			JustBeforeEach(func() {
				ticker = time.NewTicker(10 * time.Second)
				go func(chan bool) {
					defer GinkgoRecover()
					for {
						select {
						case <-doneChan:
							ticker.Stop()
							doneAcceptChan <- true
							return
						case <-ticker.C:
							cfh.CurlApp(cfg, appName, "/slow/3000", "-f")
						}
					}
				}(doneChan)
			})

			It("should scale out", Label(acceptance.LabelSmokeTests), func() {
				WaitForNInstancesRunning(appGUID, 2, 5*time.Minute)
			})
		})

		Context("when responsetime is less than scaling in threshold", func() {

			BeforeEach(func() {
				policy = GenerateDynamicScaleInPolicy(1, 2, "responsetime", 1000)
				initialInstanceCount = 2
			})

			JustBeforeEach(func() {
				ticker = time.NewTicker(2 * time.Second)
				go func(chan bool) {
					defer GinkgoRecover()
					for {
						select {
						case <-doneChan:
							ticker.Stop()
							doneAcceptChan <- true
							return
						case <-ticker.C:
							cfh.CurlApp(cfg, appName, "/fast", "-f")
						}
					}
				}(doneChan)
			})

			It("should scale in", func() {
				WaitForNInstancesRunning(appGUID, 1, 5*time.Minute)
			})
		})

	})

	Context("when scaling by throughput", func() {

		JustBeforeEach(func() {
			doneChan = make(chan bool)
			doneAcceptChan = make(chan bool)
		})

		AfterEach(func() {
			close(doneChan)
			Eventually(doneAcceptChan, 10*time.Second).Should(Receive())
		})

		Context("when throughput is greater than scaling out threshold", func() {

			BeforeEach(func() {
				policy = GenerateDynamicScaleOutPolicy(1, 2, "throughput", 1)
				initialInstanceCount = 1
			})

			JustBeforeEach(func() {
				ticker = time.NewTicker(25 * time.Millisecond)
				go func(chan bool) {
					defer GinkgoRecover()
					for {
						select {
						case <-doneChan:
							ticker.Stop()
							doneAcceptChan <- true
							return
						case <-ticker.C:
							cfh.CurlApp(cfg, appName, "/fast", "-f")
						}
					}
				}(doneChan)
			})

			It("should scale out", func() {
				WaitForNInstancesRunning(appGUID, 2, 5*time.Minute)
			})

		})

		Context("when throughput is less than scaling in threshold", func() {

			BeforeEach(func() {
				policy = GenerateDynamicScaleInPolicy(1, 2, "throughput", 100)
				initialInstanceCount = 2
			})

			JustBeforeEach(func() {
				ticker = time.NewTicker(10 * time.Second)
				go func(chan bool) {
					defer GinkgoRecover()
					for {
						select {
						case <-doneChan:
							ticker.Stop()
							doneAcceptChan <- true
							return
						case <-ticker.C:
							cfh.CurlApp(cfg, appName, "/fast", "-f")
						}
					}
				}(doneChan)
			})
			It("should scale in", func() {
				WaitForNInstancesRunning(appGUID, 1, 5*time.Minute)
			})
		})

	})

	// To check existing aggregated cpu metrics do: cf asm APP_NAME cpu
	// See: https://www.ibm.com/docs/de/cloud-private/3.2.0?topic=SSBS6K_3.2.0/cloud_foundry/integrating/cfee_autoscaler.html
	Context("when scaling by cpu", func() {

		BeforeEach(func() {
			policy = GenerateDynamicScaleOutAndInPolicy(1, 2, "cpu", int64(float64(cfg.CPUUpperThreshold)*0.2), int64(float64(cfg.CPUUpperThreshold)*0.4))
			initialInstanceCount = 1
		})

		It("when cpu is greater than scaling out threshold", func() {
			By("should scale out to 2 instances")
			AppSetCpuUsage(cfg, appName, int(float64(cfg.CPUUpperThreshold)*0.9), 5)
			WaitForNInstancesRunning(appGUID, 2, 5*time.Minute)

			By("should scale in to 1 instance after cpu usage is reduced")
			//only hit the one instance that was asked to run hot.
			AppEndCpuTest(cfg, appName, 0)

			WaitForNInstancesRunning(appGUID, 1, 10*time.Minute)
		})
	})
})
