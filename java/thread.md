#Thread
两种方式实现多线程，继承Thread,实现接口Runnable

##线程状态
- 新建(new)：新创建了一个线程对象。
- 可运行(runnable)：线程对象创建后，其他线程(比如main线程）调用了该对象的start()方法。该状态的线程位于可运行线程池中，等待被线程调度选中，获取cpu 的使用权 。
- 运行(running)：可运行状态(runnable)的线程获得了cpu 时间片（timeslice） ，执行程序代码。
- 阻塞(block)：阻塞状态是指线程因为某种原因放弃了cpu 使用权，也即让出了cpu timeslice，暂时停止运行。直到线程进入可运行(runnable)状态，才有机会再次获得cpu timeslice 转到运行(running)状态。阻塞的情况分三种：   
(一). 等待阻塞：运行(running)的线程执行o.wait()方法，JVM会把该线程放入等待队列(waitting queue)中。  
(二). 同步阻塞：运行(running)的线程在获取对象的同步锁时，若该同步锁被别的线程占用，则JVM会把该线程放入锁池(lock pool)中。  
(三). 其他阻塞：运行(running)的线程执行Thread.sleep(long ms)或t.join()方法，或者发出了I/O请求时，JVM会把该线程置为阻塞状态。当sleep()状态超时、join()等待线程终止或者超时、或者I/O处理完毕时，线程重新转入可运行(runnable)状态。  
- 死亡(dead)：线程run()、main() 方法执行结束，或者因异常退出了run()方法，则该线程结束生命周期。死亡的线程不可再次复生。

##方法
- start()开始  
- join()线程合并  
入参空：当前线程等待加入该线程后面，等待该线程中止  
入参long m:当前线程等待该线程终止的时间最长为 m 毫秒。 如果在m时间内，该线程没有执行完，那么当前线程进入就绪状态，重新等待cpu调度  
入参long m,int n:等待该线程终止的时间最长为 m 毫秒 + n 纳秒  
- sleep()睡眠  
sleep是静态方法，最好不要用Thread的实例对象调用它，因为它睡眠的始终是当前正在运行的线程，而不是调用它的线程对象，它只对正在运行状态的线程对象有效  
- yield()暂停  
yield()方法和sleep()方法有点相似，它也是Thread类提供的一个静态的方法，它也可以让当前正在执行的线程暂停，让出cpu资源给其他的线程。但是和sleep()方法不同的是，它不会进入到阻塞状态，而是进入到就绪状态。yield()方法只是让当前线程暂停一下，重新进入就绪的线程池中，让系统的线程调度器重新调度器重新调度一次  
- setDaemon(boolean on)将该线程标记为守护线程或用户线程。
- setPriority(int newPriority)更改线程的优先级
- interrupt()中止线程

##线程优先级
每个线程执行时都有一个优先级的属性，优先级高的线程可以获得较多的执行机会，而优先级低的线程则获得较少的执行机会  
每个线程默认的优先级都与创建它的父线程具有相同的优先级，在默认情况下，main线程具有普通优先级

##守护进程和用户进程
守护线程使用的情况较少，但并非无用，举例来说，JVM的垃圾回收、内存管理等线程都是守护线程。还有就是在做数据库应用时候，使用的数据库连接池，连接池本身也包含着很多后台线程，监控连接个数、超时时间、状态等等  

方法
setDaemon(boolean on)  
参数：  
on - 如果为 true，则将该线程标记为守护线程  
抛出：  
IllegalThreadStateException - 如果该线程处于活动状态  
SecurityException - 如果当前线程无法修改该线程  

该方法必须在启动线程前调用。 该方法首先调用该线程的 checkAccess 方法，且不带任何参数。这可能抛出 SecurityException（在当前线程中）  

##常用的中止方式
Thread中的stop()和suspend()方法，由于固有的不安全性，已经建议不再使用！  
###interrupt(),中断本线程  
- 本线程中断自己是被允许的；其它线程调用本线程的interrupt()方法时，会通过checkAccess()检查权限。这有可能抛出SecurityException异常。  
- 如果本线程是处于阻塞状态：调用interrupt()会立即将线程的中断标记设为“true”，但是由于线程处于阻塞状态，所以该“中断标记”会立即被清除为“false”，同时，会产生一个InterruptedException的异常。  
- 如果线程被阻塞在一个Selector选择器中，那么通过interrupt()中断它时；线程的中断标记会被设置为true，并且它会立即从选择操作中返回。
- 如果不属于前面所说的情况，那么通过interrupt()中断线程时，它的中断标记会被设置为“true”

###终结阻塞的线程
interrupt使用：  
<pre><code>
@Override
public void run() {
    try {
        while (true) {
            // 执行任务...
        }
    } catch (InterruptedException ie) {  
        // 若线程处于堵塞状态，产生InterruptedException异常，退出while(true)循环，线程终止！
    }
}</code></pre>

###终结运行状态线程
<pre><code>
@Override
public void run() {
    while (!isInterrupted()) {
        // 执行任务...
    }
}
</code></pre>
isInterrupted()是判断线程的中断标记是不是为true。当线程处于运行状态，并且我们需要终止它时；可以调用线程的interrupt()方法，使用线程的中断标记为true，即isInterrupted()会返回true。此时，就会退出while循环。  
注意：interrupt()并不会终止处于“运行状态”的线程！它会将线程的中断标记设为true。  

###通用中断线程方法
<pre><code>
@Override
public void run() {
    try {
        // 1. isInterrupted()保证，只要中断标记为true就终止线程。
        while (!isInterrupted()) {
            // 执行任务...
        }
    } catch (InterruptedException ie) {  
        // 2. InterruptedException异常保证，当InterruptedException异常产生时，线程被终止。
    }
}
</code></pre>

###检测线程状态，interrupted()和isInterrupted()
interrupted() 和 isInterrupted()都能够用于检测对象的“中断标记”。  
区别是，interrupted()除了返回中断标记之外，它还会清除中断标记(即将中断标记设为false)；而isInterrupted()仅仅返回中断标记。
