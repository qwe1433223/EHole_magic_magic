### 介绍

> 可以指定状态码和标题以排除不想要的数据，支持从管道符传入参数，根据cms类型来做总结输出

> 原项目地址https://github.com/lemonlove7/EHole_magic
>
> golang只能看懂一点点，所以大部分代码都是gpt辅助完成的，今天gpt炸了，搞得我都休息半天了

> 附上一个合并finger.json的脚本，合并后如果a.json和b.json里的某个cms名字相同，并且规则也相同会自动去重。+v获取
>
> 希望各位兄弟们能够分享一下finger.json，十分感谢，这里附上我的finger.json
>
> 加我vx，互相分享指纹文件

<img src="https://github.com/qwe1433223/EHole_magic_magic/blob/main\README.assets\image-20240605103031397.png" alt="image-20240605103031397" style="zoom:25%;" />

### 效果

![image-20240604204737304](https://github.com/qwe1433223/EHole_magic_magic/blob/main/README.assets/image-20240604204737304.png)



### 魔改1

> 在FinScan结构体中加了两个列表，默认的错误标题有这些，然后可以通过config.ini指定不需要的标题

![image-20240604202019107](https://github.com/qwe1433223/EHole_magic_magic/blob/main/README.assets/image-20240604202019107.png)

![image-20240604202237155](https://github.com/qwe1433223/EHole_magic_magic/blob/main/README.assets/image-20240604202237155.png)

```php
[error_StatusCode]
error_StatusCode1 = 404
error_StatusCode2 = 403
error_StatusCode3 = 500
error_StatusCode4 = 503
error_StatusCode4 = 505

[error_title]
error_message1 = "错误信息"
error_message2 = "Not Found"
error_message3 = "域名到期"
error_message4 = "未备案"
error_message5 = "未注册"
```

### 魔改2

> 最后总结的时候，根据CMS来进行分类输出，看资产的时候感觉脑子都好使了

![image-20240604202531847](https://github.com/qwe1433223/EHole_magic_magic/blob/main/README.assets/image-20240604202531847.png)

### 魔改3

> 支持从管道符里获取url，联合fofax工具(高级语法)不要太爽了，又方便又好玩又炫酷，脚本小子乐开怀

```php
代码我忘记在哪了。。应该不在finger.go文件内。。
```

![image-20240604203505328](https://github.com/qwe1433223/EHole_magic_magic/blob/main/README.assets/image-20240604203505328.png)
