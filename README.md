# chatWithUnity
实现unityUI前端和go后端进行通信，实现简单的聊天室功能

通过go语言的sync.RWMutex，在广播消息同步给在线用户时进行加锁，使用chan实现了进程间的通信，保证了广播队列的线程安全。使用Unity的UI组件简单的写了一个消息界面，实现了简单的消息同步和正常交流。两者均打包为了exe文件。


此库为go后端部分代码，exe为打包后文件
