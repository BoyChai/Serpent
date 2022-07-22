# Serpent
## 前言
网络测压工具，本工具仅允许合法压力测试，禁止非法攻击，相关法律责任由使用者承担，作者不承担任何法律责任！
## 使用
Serpent version 查看版本  
Serpent <模式> {Options}  
模式 sniff  
-ip 指定主机  
-p  指定端口(1,2,3,4||1-8000,20)  
-c  指定缓存通道数量,默认为2000,太低会导致太慢，太高会导致输出结果的准确性，根据自己主机性能来确定  
ps: 嗅探端口只能判断端口是否开启目前只支持tcp扫描，并且只能扫描端口是否开放不能扫除任何相关端口的信息，扫描出之后建议使用nmap进行相信的扫描  

模式: http  
-u 指定URL地址  
-c 指定线程  
-t 指定请求类型 get 或者 post

模式: mail  
-s 指定smtp地址  
-p 指定端口,默认25  
-m 指定对方邮箱  
-t 指定线程,默认10(不建议太高)  