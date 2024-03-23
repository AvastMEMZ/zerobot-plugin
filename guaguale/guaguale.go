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

func guagualeMain(num, addmin, addmax int, uid int64) (err error, money int, add int) {
	err = wallet.InsertWalletOf(uid, num)
	add = random(addmin, addmax)
	err = wallet.InsertWalletOf(uid, add)
	money = wallet.GetWalletOf(uid)
	if err != nil {
		return err, 0, 0
	} else {
		return nil, money, add
	}
}

func init() {
	engine := control.Register("example", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "小游戏",
		Help: "刮刮乐" +
			"输入 刮刮乐/彩票 10/200/1000/10000/100000 来进行小游戏",
		PublicDataFolder: "Example",
		OnEnable: func(ctx *zero.Ctx) {
			ctx.Send("刮刮乐已启用")
		},
		// 自定义插件关闭时的回复
		OnDisable: func(ctx *zero.Ctx) {
			ctx.Send("刮刮乐已禁用")
		},
	})
	engine.OnRegex(`^(刮刮乐|彩票)\s*(.+)$`).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		number := ctx.State["regex_matched"].([]string)[2]
		num, err := strconv.Atoi(number)
		if err != nil {
			ctx.SendChain(message.Text("ERROR:", err))
			return
		}
		uid := ctx.Event.UserID
		if num == 0 {
			ctx.SendChain(message.Text("输入 刮刮乐/彩票 10/200/1000/10000/100000 来进行小游戏"))
		} else if num == 10 { //十块本金
			oldmoney := wallet.GetWalletOf(uid)
			if oldmoney < 10 {
				ctx.SendChain(message.Text("你的金币不够哦", "\n你的钱包还有", oldmoney, "Atri币"))
			} else if oldmoney >= 10 {
				myrand := random(1, 1000)
				if myrand <= 300 {
					err = wallet.InsertWalletOf(uid, -10)
					money := wallet.GetWalletOf(uid)
					ctx.SendChain(message.Text("你获得了0金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 600 {
					err, money, add := guagualeMain(-10, 11, 20, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 800 {
					err, money, add := guagualeMain(-10, 21, 50, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 950 {
					err, money, add := guagualeMain(-10, 51, 80, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 1000 {
					err, money, add := guagualeMain(-10, 81, 80, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				}
			}
		} else if num == 200 { //二百块本金
			oldmoney := wallet.GetWalletOf(uid)
			if oldmoney < 200 {
				ctx.SendChain(message.Text("你的金币不够哦", "\n你的钱包还有", oldmoney, "Atri币"))
			} else if oldmoney >= 200 {
				myrand := random(1, 1000)
				if myrand <= 350 {
					err = wallet.InsertWalletOf(uid, -200)
					money := wallet.GetWalletOf(uid)
					ctx.SendChain(message.Text("你获得了0金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 600 {
					err, money, add := guagualeMain(-200, 201, 250, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 800 {
					err, money, add := guagualeMain(-200, 251, 400, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 950 {
					err, money, add := guagualeMain(-200, 401, 500, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 1000 {
					err, money, add := guagualeMain(-200, 500, 1000, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				}
			}
		} else if num == 1000 { //一千块本金
			oldmoney := wallet.GetWalletOf(uid)
			if oldmoney < 1000 {
				ctx.SendChain(message.Text("你的金币不够哦", "\n你的钱包还有", oldmoney, "Atri币"))
			} else if oldmoney >= 1000 {
				myrand := random(1, 1000)
				if myrand <= 550 {
					err = wallet.InsertWalletOf(uid, -1000)
					money := wallet.GetWalletOf(uid)
					ctx.SendChain(message.Text("你获得了0金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 750 {
					err, money, add := guagualeMain(-1000, 1001, 1150, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 900 {
					err, money, add := guagualeMain(-1000, 1151, 1500, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 970 {
					err, money, add := guagualeMain(-1000, 1501, 2000, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 1000 {
					err, money, add := guagualeMain(-1000, 2001, 5000, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				}
			}
		} else if num == 10000 {
			oldmoney := wallet.GetWalletOf(uid)
			if oldmoney < 10000 {
				ctx.SendChain(message.Text("你的金币不够哦", "\n你的钱包还有", oldmoney, "Atri币"))
			} else if oldmoney >= 10000 {
				myrand := random(1, 1000)
				if myrand <= 600 {
					err, money, add := guagualeMain(-10000, 100, 4800, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 800 {
					err, money, add := guagualeMain(-10000, 10001, 12500, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 920 {
					err, money, add := guagualeMain(-10000, 12500, 16000, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 990 {
					err, money, add := guagualeMain(-10000, 16001, 20000, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 1000 {
					err, money, add := guagualeMain(-10000, 20001, 50000, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				}
			}
		} else if num == 100000 {
			oldmoney := wallet.GetWalletOf(uid)
			if oldmoney < 100000 {
				ctx.SendChain(message.Text("你的金币不够哦", "\n你的钱包还有", oldmoney, "Atri币"))
			} else if oldmoney >= 100000 {
				myrand := random(1, 1000)
				if myrand <= 705 {
					err, money, add := guagualeMain(-100000, 1, 78000, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 855 {
					err, money, add := guagualeMain(-100000, 100001, 120000, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 945 {
					err, money, add := guagualeMain(-100000, 120001, 150000, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 995 {
					err, money, add := guagualeMain(-100000, 150001, 200000, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				} else if myrand <= 1000 {
					err, money, add := guagualeMain(-100000, 200001, 1000000, uid)
					ctx.SendChain(message.Text("你获得了", add, "金币", "\n你的钱包之前有", oldmoney, "Atri币", "\n你的钱包现在有", money, "Atri币"))
					if err != nil {
						ctx.SendChain(message.Text("ERROR: ", err))
						return
					}
				}
			}
		} else {
			ctx.SendChain(message.Text("输入\n刮刮乐/彩票\n10/200/1000/10000/100000\n来进行小游戏\n目前只支持单次"))
		}
	})
}
