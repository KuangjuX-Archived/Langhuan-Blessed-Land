# Lang Huan Blessed Land

## Introduction

**Mysql&&Redis&&WebSocket&&Docker&&OAuth**

Lang Huan Blessed Land(琅嬛福地) is a place in the chivalrous story  "The semi Gods and semi Devils"(天龙八部) written by Jin Yong which is located in a deep valley stone cave in the Wuliang Mountain of Dali Kingdom.  The head of the XiaoYao Pai(逍遥派) Wu Yazi(无崖子) and his junior sister QiuShui Li(李秋水) gave birth to QingLuo Li(李青萝) and lived together after then. In this place, there are hidden martial arts secrets from all walks of life in the world.

金庸名著《天龙八部》中的地名，位於大理国无量山中一深谷石洞中，「逍遥派」掌门人无崖子与师妹李秋水两人生了李青萝（即王夫人）爱女後，共居此地中，在山洞内藏有普天下各路的武林秘笈。

This project is named "Lang Huan Blessed Land" because "Lang Huan Blessed Land" is a place where martial arts classics are collected in stroy. And I want to build a community to help me or others store own articles and thoughts into this project. To be honest, although there are already various communities in the world that agree with my ideas and whose technology is much better than mine, I still want to build my own community through technology and feel truely happy and satisfied.

这个项目被命名为“琅嬛福地”是因为琅嬛福地在原著中是一个收藏天下武学典籍的地方。我想去帮助我和其他人去构建一个社区用来存放我们的文章和思想。虽然在世界上已经有许多和我想法一致并且技术含量比我好得多的项目，但我仍然想自己构建这个项目，并且真切地感到高兴和满足。

## Config

- Go  >=1.14
- Mysql
- Redis
- Docker

## Architecture

![](/docs/static-files/Architecture.png)

## Usage

```
git clone https://github.com/KuangjuX/Langhuan-Blessed-Land.git
cd Langhuan-Blessed-Land
docker build -t Langhuan-Blessed-Land .
docker run -p 8081:8081 -d Langhuan-Blessed-Land
```

You can visit `http://baseurl:8081` to use the app.

In addition, you can also use `nginx.conf` to carry out Proxy forwarding.

## Api

**Please see docs/api.**