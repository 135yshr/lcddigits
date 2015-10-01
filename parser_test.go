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
				Expect(Parse(8)).To(Equal("._.|_||_|"))
			})
		})
		Context("数字の０が渡されたとき", func() {
			It("LCDに表示する０が返ってくるべき", func() {
				Expect(Parse(0)).To(Equal("._.|.||_|"))
			})
		})
		Context("数字の６が渡されたとき", func() {
			It("LCDに表示する６が返ってくるべき", func() {
				Expect(Parse(6)).To(Equal("._.|_.|_|"))
			})
		})
		Context("数字の５が渡されたとき", func() {
			It("LCDに表示する５が返ってくるべき", func() {
				Expect(Parse(5)).To(Equal("._.|_.._|"))
			})
		})
		Context("数字の９が渡されたとき", func() {
			It("LCDに表示する９が返ってくるべき", func() {
				Expect(Parse(9)).To(Equal("._.|_|..|"))
			})
		})
		Context("数字の７が渡されたとき", func() {
			It("LCDに表示する７が返ってくるべき", func() {
				Expect(Parse(7)).To(Equal("._.|.|..|"))
			})
		})
		Context("数字の３が渡されたとき", func() {
			It("LCDに表示する３が返ってくるべき", func() {
				Expect(Parse(3)).To(Equal("._.._|._|"))
			})
		})
		Context("数字の２が渡されたとき", func() {
			It("LCDに表示する２が返ってくるべき", func() {
				Expect(Parse(2)).To(Equal("._.._||_."))
			})
		})
		Context("数字の１が渡されたとき", func() {
			It("LCDに表示する１が返ってくるべき", func() {
				Expect(Parse(1)).To(Equal(".....|..|"))
			})
		})
		Context("数字の４が渡されたとき", func() {
			It("LCDに表示する４が返ってくるべき", func() {
				Expect(Parse(4)).To(Equal("...|_|..|"))
			})
		})
	})
})
