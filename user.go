package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

// user.go 脚本

type User struct {
	Name    string      // 昵称，默认与Addr相同
	Addr    string      // 地址
	Channel chan string // 消息管道
	conn    net.Conn    // 连接
	server  *Server     // 缓存Server的引用
}

// NewUser 创建新用户
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:    userAddr,
		Addr:    userAddr,
		Channel: make(chan string),
		conn:    conn,
		server:  server,
	}
	go user.ListenMessage()
	return user
}

// 上线
func (this *User) Online() {

	// 用户上线，将用户加入到OnlineMap中，注意加锁操作
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播当前用户上线消息
	this.server.BroadCast(this, "上线啦O(∩_∩)O")
}

// 下线
func (this *User) Offline() {

	// 用户下线，将用户从OnlineMap中删除，注意加锁
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// 广播当前用户下线消息
	this.server.BroadCast(this, "下线了o(╥﹏╥)o")
}

// user.go 脚本
func (this *User) DoMessage(buf []byte, len int) {
	//提取用户的消息(去除'\n')
	msg := string(buf[:len-1])
	// 调用Server的BroadCast方法
	this.server.BroadCast(this, msg)
}

func (this *User) ListenMessage() {
	for {
		msg := <-this.Channel
		fmt.Println("Send msg to client: ", msg, ", len: ", int16(len(msg)))
		bytebuf := bytes.NewBuffer([]byte{})
		// 前两个字节写入消息长度
		binary.Write(bytebuf, binary.BigEndian, int16(len(msg)))
		// 写入消息数据
		binary.Write(bytebuf, binary.BigEndian, []byte(msg))
		// 发送消息给客户端
		this.conn.Write(bytebuf.Bytes())
	}
}
