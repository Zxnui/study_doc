##web work
一种后台运行的js，不会影响页面性能

- 检测是否支持web work
if(typeof(Worker)!=="undefined"){
	// Yes! Web worker support!
	// Some code.....
}else{
	// Sorry! No Web Worker support..
}

- 创建web work文件
定义一个js文件，deam.js
里面的内容：
var i=0;
function timedCount(){
	i=i+1;
	postMessage(i);
	setTimeout("timedCount()",500);
}
timedCount();

postMessage()方法，向html传回一段消息

- 创建实例

if(typeof(w)=="undefined"){
  w=new Worker("demo_workers.js");
}

向 web worker 添加一个 "onmessage" 事件监听器：
w.onmessage=function(event){
	document.getElementById("result").innerHTML=event.data;
};
当 web worker 传递消息时，会执行事件监听器中的代码。event.data 中存有来自 event.data 的数据。

- 终止实例
w.terminate();