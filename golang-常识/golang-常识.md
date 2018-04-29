#### 类型别名

```bash
rune类型是int32类型的别名，
byte类型是uint8的别名
byte // uint8 的别名 ：字符类型
rune // int32 的别名，代表一个Unicode码，用UTF-8 进行编码。

# rune的使用场景
用string存储unicode的话，如果有中文，按下标是访问不到的，因为你只能得到一个byte。 要想访问中文的话，还是要用rune切片，这样就能按下表访问。
```

在Go当中 string底层是用byte数组存的，并且是不可以改变的。