package comm

type CodeType int

const (
	Success      CodeType = 0
	NoLogin      CodeType = 1001
	NowLottery   CodeType = 1002 //正在抽奖
	NoLotteryNum CodeType = 1003
)
