# ZK
## docker
## 集群
## 数据模型
层次化结构,类似于unix目录树
每个节点称为节点ZNode,每个节点都会保存自己的数据信息(1MB)和节点信息
节点可以拥有子节点
## 节点分类

    1.持久节点(PERSISTENT)
    持久节点，创建后一直存在，直到主动删除此节点。

    2.持久顺序节点(PERSISTENT_SEQUENTIAL)
    持久顺序节点，创建后一直存在，直到主动删除此节点。在ZK中，每个父节点会为它的第一级子节点维护一份时序，记录每个子节点创建的先后顺序。
    -s

    3.临时节点(EPHEMERAL)
    临时节点在客户端会话失效后节点自动清除。临时节点下面不能创建子节点。
    -e

    4.顺序临时节点(EPHEMERAL_SEQUENTIAL)
    临时节点在客户端会话失效后节点自动清除。临时节点下面不能创建子节点。父节点getChildren会获得顺序的节点列表。
    -e -s

    -t ttl
## 服务端命令

    start
    stop
    status
    restart
## 客户端命令
```
ZooKeeper -server host:port -client-configuration properties-file cmd args
	addWatch [-m mode] path # optional mode is one of [PERSISTENT, PERSISTENT_RECURSIVE] - default is PERSISTENT_RECURSIVE
	addauth scheme auth
	close 
	config [-c] [-w] [-s]
	connect host:port
	create [-s] [-e] [-c] [-t ttl] path [data] [acl]
	delete [-v version] path
	deleteall path [-b batch size]
	delquota [-n|-b|-N|-B] path
	get [-s] [-w] path
	getAcl [-s] path
	getAllChildrenNumber path
	getEphemerals path
	history 
	listquota path
	ls [-s] [-w] [-R] path
	printwatches on|off
	quit 
	reconfig [-s] [-v version] [[-file path] | [-members serverID=host:port1:port2;port3[,...]*]] | [-add serverId=host:port1:port2;port3[,...]]* [-remove serverId[,...]*]
	redo cmdno
	removewatches path [-c|-d|-a] [-l]
	set [-s] [-v version] path data
	setAcl [-s] [-v version] [-R] path acl
	setquota -n|-b|-N|-B val path
	stat [-w] path
	sync path
	version 
	whoami 
```
stat 命令解析
```
cZxid	         数据节点创建时的事务ID
ctime	         数据节点创建时的时间
mZxid	         数据节点最后一次更新时的事务ID
mtime	         数据节点最后一次更新时的时间
pZxid	         数据节点的子节点列表最后一次被修改（是子节点列表变更，而不是子节点内容变更）时的事务ID
cversion	     子节点的版本号(创建,删除有效,更新子节点的数据信息不影响)
dataVersion	     数据节点的版本号
aclVersion	     数据节点的ACL版本号 
ephemeralOwner 	 如果节点是临时节点，则表示创建该节点的会话的SessionID如果节点是持久节点，则该属性值为0
dataLength	     数据内容的长度
numChildren	     数据节点当前的子节点个数
```