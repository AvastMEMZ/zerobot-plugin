package guaguale

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/FloatTech/AnimeAPI/wallet"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func random(min, max int) int {
	return globalRand.Intn(max-min) + min
}

func init() {
	engine := control.Register("example", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "小游戏",
		Help: "猜大小" +
			"输入 猜大/猜小 10/200/1000/10000/100000 来进行小游戏",
		PublicDataFolder: "Example",
		OnEnable: func(ctx *zero.Ctx) {
			ctx.Send("猜大小已启用")
		},
		// 自定义插件关闭时的回复
		OnDisable: func(ctx *zero.Ctx) {
			ctx.Send("猜大小已禁用")
		},
	})
	engine.OnRegex(`^(猜小|猜大)\s*(.+)$`).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		number := ctx.State["regex_matched"].([]string)[1]
		num, err := strconv.Atoi(number)
		if err != nil {
			ctx.SendChain(message.Text("ERROR:", err))
			return
		}
		uid := ctx.Event.UserID
		if num == 0 {
			ctx.SendChain(message.Text("输入 猜大/猜小 10/200/1000/10000/100000 来进行小游戏"))
		} else if num == 10 {
			myrand := random(1, 1000)
			var err error
			if myrand <= 300 {
				add := -10
				err = wallet.InsertWalletOf(uid, add)
				ctx.SendChain(message.Text("你获得了", add, "金币"))
				if err != nil {
					ctx.SendChain(message.Text("ERROR: ", err))
					return
				}
			} else if myrand <= 600 {
				add := random(11, 20)
				err = wallet.InsertWalletOf(uid, add)
				ctx.SendChain(message.Text("你获得了", add, "金币"))
				if err != nil {
					ctx.SendChain(message.Text("ERROR: ", err))
					return
				}
			} else if myrand <= 800 {
				add := random(21, 50)
				err = wallet.InsertWalletOf(uid, add)
				ctx.SendChain(message.Text("你获得了", add, "金币"))
				if err != nil {
					ctx.SendChain(message.Text("ERROR: ", err))
					return
				}
			} else if myrand <= 950 {
				add := random(51, 80)
				err = wallet.InsertWalletOf(uid, add)
				ctx.SendChain(message.Text("你获得了", add, "金币"))
				if err != nil {
					ctx.SendChain(message.Text("ERROR: ", err))
					return
				}
			} else if myrand <= 1000 {
				add := random(81, 100)
				err = wallet.InsertWalletOf(uid, add)
				ctx.SendChain(message.Text("你获得了", add, "金币"))
				if err != nil {
					ctx.SendChain(message.Text("ERROR: ", err))
					return
				}
			}
		}
	})
}
