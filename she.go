package main

import (
	"Clib"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const WIDE int = 20  // 设置布局
const HIGH int = 20

var food Food  // 定义一个全局变量，存食物信息

type Position struct{
	X int  // 存储位置
	Y int
}

// 定义蛇类
type Snake struct{
	size int
	dir byte
	pos [WIDE*HIGH]Position
}

// 食物基本信息
type Food struct{
	Position
}

// 蛇信息初始化
func(s *Snake) SnakeInit(){
	s.size = 2  // 一头一尾
	s.dir = 'R'  //初始化方向 用UDLR做上下左右
	s.pos[0].X = WIDE/2
	s.pos[0].Y = HIGH/2
	s.pos[1].X = WIDE/2 - 1
	s.pos[1].Y = HIGH/2
	//绘制蛇
	for i := 0; i < s.size; i++{
		var ui byte
		if i == 0{
			ui = '@'
		} else{
			ui = '*'
		}
		ShowUI(s.pos[i].X, s.pos[i].Y, ui)
	}

	// 接受键盘按键信息
	// go添加一个独立函数，非阻塞信息
	// 需要查找c语言的ASCII表
	go func() {
		for {
			switch Clib.Direction() {
			//方向上  W|w|↑
			case 83, 115, 80:
				s.dir = 'U'
				//方向左
			case 65, 97, 75:
				s.dir = 'L'
				//方向右
			case 100, 68, 77:
				s.dir = 'R'
				//方向下
			case 72, 87, 119:
				s.dir = 'D'
				//暂停  空格键
			case 32:
				s.dir = 'P'
			}
		}
	}()  // 特殊的独立函数
}

// 游戏逻辑
func (s *Snake) PlayGame(){
	for{
		// 延迟执行,贪吃蛇移动速度可能太快，需要为其降低速度
		time.Sleep(time.Second/4)
		// 设置一个蛇的坐标移动量
		var nx, ny int = 0, 0
		// 更新蛇的位置
		switch s.dir {
		case 'U':
			nx = 0
			ny = 1
		case 'D':
			nx = 0
			ny = -1
		case 'L':
			nx = -1
			ny = 0
		case 'R':
			nx = 1
			ny = 0
		}

		// 蛇头和墙体碰撞判断
		if s.pos[0].X < 1 || s.pos[0].Y < 1 || s.pos[0].X >= WIDE+1 || s.pos[0].Y >= HIGH+1{
			return
		}
		// 蛇头和身体判断
		for i := 1; i < s.size; i++{
			if s.pos[0].X == s.pos[i].X && s.pos[0].Y == s.pos[i].Y{
				return
			}
		}
		// 蛇跟食物判断
		if s.pos[0].X == food.X && s.pos[0].Y == food.Y{
			// 身体增长
			s.size++
			// 刷新食物位置
			RandomFood()
			// 分数变量
		}

		//获取蛇尾坐标,在之后隐去蛇尾
		wx := s.pos[s.size-1].X
		wy := s.pos[s.size-1].Y

		// 从尾部开始更新蛇的身体坐标
		for i := s.size-1; i > 0; i-- {
			s.pos[i].X = s.pos[i-1].X
			s.pos[i].Y = s.pos[i-1].Y
		}

		// 更新蛇头坐标
		s.pos[0].X += nx
		s.pos[0].Y += ny

		//绘制蛇
		for i := 0; i < s.size; i++{
			var ui byte
			if i == 0{
				ui = '@'
			} else{
				ui = '*'
			}
			ShowUI(s.pos[i].X, s.pos[i].Y, ui)
		}

		ShowUI(wx, wy, ' ')
	}
}  // 注意方法的定义方式

// 随机产生食物
func RandomFood(){
	food.X = rand.Intn(WIDE)+1  // 直接写起不到随机的效果,这是go的一个缺点
	food.Y = rand.Intn(HIGH)+1
	ShowUI(food.X, food.Y, 's')
}

// 初始化地图信息
func MapInit(){
	fmt.Fprintln(os.Stderr, `
  #-----------------------------------------#
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  #-----------------------------------------#
`)  // 采用默认格式将其参数格式化并写入 w 。总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符。返回写入的字节数和遇到的任何错误。
}
func ShowUI(X int, Y int, ui byte){
	// 找到对应坐标点光标位置
	Clib.GotoPosition(X*2+2,Y+2)  // 避免蛇移动影响棋盘边界
	// 绘制图形
	fmt.Fprintf(os.Stderr, "%c", ui)
}

func main() {
	// 设置一个随机种子，用作混淆
	rand.Seed(time.Now().UnixNano())
	// 隐藏光标
	Clib.HideCursor()
	// 初始化地图
	MapInit()
	// 生成随机食物
	RandomFood()
	// 绘制食物
	//ShowUI(food.X, food.Y, 's')

	var s Snake
	s.SnakeInit()

	s.PlayGame()

}