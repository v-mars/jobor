package redis

/*// SET key value EX 10 NX
set, err := rdb.SetNX(ctx, "key", "value", 10*time.Second).Result()

// SET key value keepttl NX
set, err := rdb.SetNX(ctx, "key", "value", redis.KeepTTL).Result()

// SORT list LIMIT 0 2 ASC
vals, err := rdb.Sort(ctx, "list", &redis.Sort{Offset: 0, Count: 2, Order: "ASC"}).Result()

// ZRANGEBYSCORE zset -inf +inf WITHSCORES LIMIT 0 2
vals, err := rdb.ZRangeByScoreWithScores(ctx, "zset", &redis.ZRangeBy{
Min: "-inf",
Max: "+inf",
Offset: 0,
Count: 2,
}).Result()


// ZINTERSTORE out 2 zset1 zset2 WEIGHTS 2 3 AGGREGATE SUM
vals, err := rdb.ZInterStore(ctx, "out", &redis.ZStore{
Keys: []string{"zset1", "zset2"},
Weights: []int64{2, 3}
}).Result()

// EVAL "return {KEYS[1],ARGV[1]}" 1 "key" "hello"
vals, err := rdb.Eval(ctx, "return {KEYS[1],ARGV[1]}", []string{"key"}, "hello").Result()

// custom command
res, err := rdb.Do(ctx, "set", "key", "value").Result()*/