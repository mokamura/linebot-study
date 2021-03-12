# LINE bot & ビーコンメモ

## ビーコンデバイスの識別
- 公式アカウントに紐付いたLINE Simple Beacon用のHWIDを発行、デバイス側でAdvertiseDataに乗せることで識別する
- LINE Official Account ManagerやLINE Developersをざっと探してもリンクが見つからなかったけども、どこから到達できるんだろう。。
  - [直リンク](https://manager.line.biz/beacon/register)
- LINEがビーコンを受け取るとLINEのサーバに通知がいき、そのときHWIDをもとに通知先のbotサーバが決定されるしくみ
  - HWIDだけで判断してるっぽいので、なりすましできてしまう
  - HWIDは1アカウント最大10個までのようで、同種のデバイスは同じものを使う前提？
    - 端末を識別したければ、HWIDとは別に何かしらIDの用意が必要
    - AdvertiseData中に最大13byteの開発者が自由に使える領域(DeviceMessage)があるので、そこを使う？
    - 端末識別だけならまだしも、位置や気温など付随データを入れようとするときつい

## ユーザ(スマホ)の識別
- botサーバにビーコンイベント通知がくる際、メッセージ送信のためのトークンがついてくるので、それを使うことでビーコンに反応したスマホだけにメッセージを送れる
  - イベントからユーザIDも取れる
  - [やろうと思えばアカウント連携もできて](https://developers.line.biz/ja/docs/messaging-api/linking-accounts/)、商品を買ったらメッセージを送ることも可能

## 送信可能なメッセージ
- [およそLINEっぽいメッセージは大体送れる模様](https://developers.line.biz/ja/docs/messaging-api/message-types/)
- 凝ったUIが必要なら[Flex Message](https://developers.line.biz/ja/docs/messaging-api/using-flex-messages/)を使う感じか
- [ビーコンバナー](https://developers.line.biz/ja/docs/messaging-api/using-beacons/#beacon-banner)は法人ユーザ向けなので個人で調査は難しい

## 参考
- [Line Simple Beacon仕様](https://github.com/line/line-simple-beacon/blob/master/README.ja.md)
- [ラズベリーパイでLINE Beaconが作成可能に！「LINE Simple Beacon」仕様を公開しました](https://engineering.linecorp.com/ja/blog/detail/117/)
- [LINE Developers ドキュメント(ビーコンを使う)](https://developers.line.biz/ja/docs/messaging-api/using-beacons/)