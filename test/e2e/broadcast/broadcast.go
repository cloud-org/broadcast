/*
 * MIT License
 *
 * Copyright (c) 2021 ashing
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package broadcast

import (
	"bytes"
	"os/exec"
	"syscall"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("CLI", func() {
	ginkgo.Context("test add agent and then notify", func() {
		ginkgo.It("should return logs", func() {
			// start main
			var out bytes.Buffer
			cmd := exec.Command("main")
			cmd.Stdout = &out
			cmd.Stderr = &out
			err := cmd.Start()
			gomega.Expect(err).Should(gomega.BeNil())
			pid := cmd.Process.Pid
			gomega.Expect(pid).Should(gomega.BeNumerically(">", 0))
			go func() {
				time.Sleep(3 * time.Second)
				err = cmd.Process.Signal(syscall.SIGINT)
				gomega.Expect(err).Should(gomega.BeNil())
				err = cmd.Wait()
				gomega.Expect(err).Should(gomega.BeNil())
			}()
			// start notify
			var notifyOut bytes.Buffer
			notifyCmd := exec.Command("notify")
			notifyCmd.Stdout = &notifyOut
			notifyCmd.Stderr = &notifyOut
			err = notifyCmd.Run()
			gomega.Expect(err).Should(gomega.BeNil())
			// check log
			time.Sleep(5 * time.Second)
			gomega.Expect(out.String()).Should(gomega.ContainSubstring("add agent agent_1 region_123456"))
			gomega.Expect(out.String()).Should(gomega.ContainSubstring("add agent agent_2 region_123456"))
			gomega.Expect(out.String()).Should(gomega.ContainSubstring("key: /config/region_123456 value: value change"))
			gomega.Expect(out.String()).Should(gomega.ContainSubstring("agents count 2"))
			gomega.Expect(out.String()).Should(gomega.ContainSubstring("agent_1 接收到新的配置，value change"))
			gomega.Expect(out.String()).Should(gomega.ContainSubstring("agent_2 接收到新的配置，value change"))

			gomega.Expect(notifyOut.String()).Should(gomega.ContainSubstring("value will change"))
		})
	})
})
