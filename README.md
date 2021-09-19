# GeekTrainingCamp
After class practice at GeekTime CloudNative Training Camp 

## Day01

给定一个字符串数组

["I", "am", "stupid", "and", "weak"]

用for循环遍历该数组并修改为

["I", "am", "smart", "and", "strong"]

## Day02

* **基于Channel编写一个简单的单线程生产者消费者模型**
* 队列： 队列长度10，队列元素类型为int
* 生产者：每1秒往队列中放入一个类型为int的元素，队列满时生产者可以阻塞
* 消费者：每1秒钟从队列中获取一个元素并打印，队列为空时，消费者阻塞
