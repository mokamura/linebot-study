# linebot-study

## これはなに

LINE BOT & golangの勉強用

## Setup
### development

1. git clone
```Bash
$ git clone git@github.com:mokamura/linebot-study.git
```

2. run
```Bash
$ docker-compose up
$ open http://localhost:8080
```

### deploy

2. setup heroku
```Bash
$ heroku login
$ heroku git:remote -a ${my-app-name-on-heroku}
$ heroku config:set CHANNEL_SECRET=${channel-secret-on-line} CHANNEL_ACCESS_TOKEN=${channel-access-token-on-line}
```

3. deploy
```Bash
$ ./scripts/deploy.sh
```

## 参考記事
- https://speakerdeck.com/yagieng/go-to-line-bot-ni-matometeru-men-suru
- https://qiita.com/sunnyG/items/95f61a39ca3303a6c85f

