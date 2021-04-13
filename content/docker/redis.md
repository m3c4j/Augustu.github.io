---
title: "Redis"
date: 2021-04-13T21:31:35+08:00
draft: false
---

### Redis

```bash
docker pull redis

docker run --name redis-bench \
	-v /home/Develop/data/redis/conf:/data \
	-p 127.0.0.1:6379:6379 \
	-d redis

docker run --name redis-bench \
	-p 127.0.0.1:6379:6379 \
	-d redis


```



#### Redis 集合(Set)

| 命令                                           | 描述                                                |
| ---------------------------------------------- | --------------------------------------------------- |
| SADD key member1 [member2]                     | 向集合添加一个或多个成员                            |
| SCARD key                                      | 获取集合的成员数                                    |
| SDIFF key1 [key2]                              | 返回第一个集合与其他集合之间的差异                  |
| SDIFFSTORE destination key1 [key2]             | 返回给定所有集合的差集并存储在 destination 中       |
| SINTER key1 [key2]                             | 返回给定所有集合的交集                              |
| SINTERSTORE destination key1 [key2]            | 返回给定所有集合的交集并存储在 destination 中       |
| SISMEMBER key member                           | 判断 member 元素是否是集合 key 的成员               |
| SMEMBERS key                                   | 返回集合中的所有成员                                |
| SMOVE source destination member                | 将 member 元素从 source 集合移动到 destination 集合 |
| SPOP key                                       | 移除并返回集合中的一个随机元素                      |
| SRANDMEMBER key [count]                        | 返回集合中一个或多个随机数                          |
| SREM key member1 [member2]                     | 移除集合中一个或多个成员                            |
| SUNION key1 [key2]                             | 返回所有给定集合的并集                              |
| SUNIONSTORE destination key1 [key2]            | 所有给定集合的并集存储在 destination 集合中         |
| SSCAN key cursor [MATCH pattern] [COUNT count] | 迭代集合中的元素                                    |

ref: https://www.runoob.com/redis/redis-sets.html



#### Redis 有序集合(sorted set)

| 序号 | 命令                                           | 描述                                                         |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ |
| 1    | ZADD key score1 member1 [score2 member2]       | 向有序集合添加一个或多个成员，或者更新已存在成员的分数       |
| 2    | ZCARD key                                      | 获取有序集合的成员数                                         |
| 3    | ZCOUNT key min max                             | 计算在有序集合中指定区间分数的成员数                         |
| 4    | ZINCRBY key increment member                   | 有序集合中对指定成员的分数加上增量 increment                 |
| 5    | ZINTERSTORE destination numkeys key [key ...]  | 计算给定的一个或多个有序集的交集并将结果集存储在新的有序集合 destination 中 |
| 6    | ZLEXCOUNT key min max                          | 在有序集合中计算指定字典区间内成员数量                       |
| 7    | ZRANGE key start stop [WITHSCORES]             | 通过索引区间返回有序集合指定区间内的成员                     |
| 8    | ZRANGEBYLEX key min max [LIMIT offset count]   | 通过字典区间返回有序集合的成员                               |
| 9    | ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT] | 通过分数返回有序集合指定区间内的成员                         |
| 10   | ZRANK key member                               | 返回有序集合中指定成员的索引                                 |
| 11   | ZREM key member [member ...]                   | 移除有序集合中的一个或多个成员                               |
| 12   | ZREMRANGEBYLEX key min max                     | 移除有序集合中给定的字典区间的所有成员                       |
| 13   | ZREMRANGEBYRANK key start stop                 | 移除有序集合中给定的排名区间的所有成员                       |
| 14   | ZREMRANGEBYSCORE key min max                   | 移除有序集合中给定的分数区间的所有成员                       |
| 15   | ZREVRANGE key start stop [WITHSCORES]          | 返回有序集中指定区间内的成员，通过索引，分数从高到低         |
| 16   | ZREVRANGEBYSCORE key max min [WITHSCORES]      | 返回有序集中指定分数区间内的成员，分数从高到低排序           |
| 17   | ZREVRANK key member                            | 返回有序集合中指定成员的排名，有序集成员按分数值递减(从大到小)排序 |
| 18   | ZSCORE key member                              | 返回有序集中，成员的分数值                                   |
| 19   | ZUNIONSTORE destination numkeys key [key ...]  | 计算给定的一个或多个有序集的并集，并存储在新的 key 中        |
| 20   | ZSCAN key cursor [MATCH pattern] [COUNT count] | 迭代有序集合中的元素（包括元素成员和元素分值）               |

ref: https://www.runoob.com/redis/redis-sorted-sets.html







