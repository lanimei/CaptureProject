# CaptureProject
This is IoT malware sandbox, include mips, mipsel, arm. You can also extend this project

#中文说明
此项目由于经常使用到命令，因此要注意的是命令处理机制的引入，一定要引入错误处理机制
1. 比如挂载一个qcow2文件后，因为某种错误原因未能及时unmount掉该目录，那么就可以使用一些命令强行卸载文件目录？
2. 再比如说恶意样本拷贝至qcow2后，无法执行， 这一点如何进行处理？
3. 如何单独启动一个恶意样本的镜像
4. 如何建立一个错误处理机制？主要是加载镜像的过程中出现错误

