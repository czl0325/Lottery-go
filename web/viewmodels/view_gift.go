package viewmodels

type ViewGift struct {
	Id           int    `form:"id"`
	Title        string `form:"title"`
	PrizeNum     int    `form:"prize_num"`
	PrizeCode    string `form:"prize_code"`
	PrizeTime    int    `form:"prize_time"`
	Img          string `form:"img"`
	DisplayOrder int    `form:"displayorder"`
	Gtype        int    `form:"gtype"`
	Gdata        string `form:"gdata"`
	TimeBegin    string `form:"time_begin"`
	TimeEnd      string `form:"time_end"`
}
