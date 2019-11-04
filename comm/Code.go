package comm

type CodeType int

const (
	Success      CodeType = 0
	NoLogin      CodeType = 1001
	NowLottery   CodeType = 1002 //正在抽奖
	NoLotteryNum CodeType = 1003
	IpLimit      CodeType = 1004 //相同IP参与次数太多，明天再来参与吧
	NoPrize      CodeType = 1005 //没有中奖
)
