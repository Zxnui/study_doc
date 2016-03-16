##服务器发送更新通知

###前端页面

var source=new EventSource("demo_sse.php");
source.onmessage=function(event){
  document.getElementById("result").innerHTML+=event.data + "<br />";
};

例子解释：
- 创建一个新的 EventSource 对象，然后规定发送更新的页面的 URL（本例中是 "demo_sse.php"）
- 每接收到一次更新，就会发生 onmessage 事件
- 当 onmessage 事件发生时，把已接收的数据推入 id 为 "result" 的元素中

###服务器端
把 "Content-Type" 报头设置为 "text/event-stream"

onmessage事件获取消息，还可以通以下事件来获取消息

onopen		当通往服务器的连接被打开
onmessage	当接收到消息
onerror		当错误发生