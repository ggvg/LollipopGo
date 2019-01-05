package Proto2

//  G_GameDSQ_Proto == 10    斗兽棋
const (
	INITDSQ                      = iota //  INITDSQ == 0
	DSQ2GW_ConnServerProto2             //  DSQ2GW_ConnServerProto2 == 1 DSQ主动链接 主动链接 gateway 进行注册
	GW2DSQ_ConnServerProto2             //  GW2DSQ_ConnServerProto2 == 2 选择链接
	GW2DSQ_InitGameProto2               //  GW2DSQ_InitGameProto2   == 3  初始化协议-- 相当于注册
	DSQ2GW_InitGameProto2               //  GW2DSQ_InitGameProto2   == 4
	GW2DSQ_PlayerStirChessProto2        // GW2DSQ_PlayerStirChessProto2 == 5   玩家翻棋子
	DSQ2GW_PlayerStirChessProto2        // DSQ2GW_PlayerStirChessProto2 == 6   广播同一个桌子上的,且接受到此协议后，已经移动的再无法移动棋子，对手获取操作权限

)

// 斗兽棋的棋子类型
const (
	DSQINIT_QZ = iota // DSQ_QZ == 0
	Elephant          // elephant == 1  大象
	Lion              // lion == 2 		狮子
	Tiger             // tiger == 3 	老虎
	Leopard           // leopard == 4 	豹子
	Wolf              // wolf == 5 		狼
	Dog               // dog == 6 		狗
	Cat               // cat == 7 		猫
	Mouse             // mouse == 8 	老鼠
)

// 棋子的行动的方向
const (
	FANGXIANGINIT = iota // FANGXIANGINIT == 0
	UP                   // UP 		== 1
	DOWN                 // DOWN 	== 2
	LEFT                 // LEFT 	== 3
	RIGHT                // RIGHT 	== 4
)

// 棋子的攻击方式
const (
	ITYPEINIY    = iota // ITYPEINIY == 0
	MOVE                // MOVE == 1         正常移动
	DISAPPEAR           // DISAPPEAR == 2 	 自残
	ALLDISAPPEAR        // ALLDISAPPEAR == 3 同归于尽
	BEAT                // BEAT == 4         击败对方
	TEAMMATE            // TEAMMATE == 5     队友
	MOVESUCC            // MOVESUCC == 6     移动成功
	MOVEFAIL            // MOVEFAIL == 7     移动失败
	DATAERROR           // DATAERROR == 8    数据错误    玩家的棋子已经被吃掉不存在了
	DATANOEXIT          // DATANOEXIT == 9   数据不存在  棋子的数据大于 16或者小于0
)

//------------------------------------------------------------------------------
// GW2DSQ_PlayerStirChessProto2
type GW2DSQ_PlayerStirChess struct {
	Protocol  int
	Protocol2 int
	OpenID    string
	RoomUID   int
	StirPos   string // 翻动的位置 格式: x,y
}

// DSQ2GW_PlayerStirChessProto2
type DSQ2GW_PlayerStirChess struct {
	Protocol  int
	Protocol2 int
	OpenID    string // 谁翻动了棋子
	StirPos   string // 翻动的位置  格式:x,y
	ChessNum  int    // 1 - 16 正数
}

//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// DSQ2GW_ConnServerProto2
type DSQ2GW_ConnServer struct {
	Protocol  int
	Protocol2 int
	ServerID  string //全局配置 唯一的也是
}

// GW2DSQ_ConnServerProto2
type GW2DSQ_ConnServer struct {
	Protocol  int
	Protocol2 int
	ServerID  string //全局配置 唯一的也是
}

//------------------------------------------------------------------------------
// GW2DSQ_InitGameProto2
type GW2DSQ_InitGame struct {
	Protocol  int
	Protocol2 int
	OpenID    string
	RoomID    string
}

// DSQ2GW_InitGameProto2
type DSQ2GW_InitGame struct {
	Protocol  int
	Protocol2 int
	OpenID    string
	RoomID    string
	InitData  [4][4]int // 斗兽棋的棋盘的数据
}

//------------------------------------------------------------------------------
