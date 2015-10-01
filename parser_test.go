package lcddigits_test

import (
	. "github.com/135yshr/lcddigits"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("#Parser", func() {
	Describe("数字を文字列に変換する", func() {
		Context("数字の８が渡されたとき", func() {
			It("LCDに表示する８が返ってくるべき", func() {
				Expect(Parse(8)).To(Equal("._.\n|_|\n|_|"))
			})
		})
		Context("数字の０が渡されたとき", func() {
			It("LCDに表示する０が返ってくるべき", func() {
				Expect(Parse(0)).To(Equal("._.\n|.|\n|_|"))
			})
		})
		Context("数字の６が渡されたとき", func() {
			It("LCDに表示する６が返ってくるべき", func() {
				Expect(Parse(6)).To(Equal("._.\n|_.\n|_|"))
			})
		})
	})
})
