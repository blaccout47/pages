+++
date = "2016-10-13T21:08:26+09:00"
title = "Code review"
tags = ["blog"]
+++

こういうコードレビュー系の記事はたくさんありますが、大手もスタートアップも経験した私からコードレビューついて今まで悩んてたところと気を付けてたところをリストしてみようと思います。

## レビューが終わってるかどうかわからない問題
PR出して半日経ったけど、コメントがない。やっとコメントが入り半日経過。よっし、レビューの修正とコメントに返信しようと思い修正なり、コメントに返事する。しかし、レビュー終わってなかった。また残りのレビューでコメントが入る。この繰り返しが数日続く時もある。こういうレビューの反映はほかのタスクやってうちに割り込みで対応するのでおそらくコードのクオリティよくないでしょう。割り込みで本来のタスクにも影響がでますよね。なのでレビュアーは時間をかけてもいいから一気にレビューして最後にレビュー終わりましたの合図を出すべきです。

## 言葉使いの問題
「こうしてください。」「こうしたほうがいい。」「これ意味分からないｗｗｗ」などなど、理由もなしにこういう言い方されるとイラッとしますよね？こうしてくださいの代わりにこういう理由があってこうした方がいいと思います。この部分の意図が理解できないので説明してくれると助かります。などなど、代わりの言葉はいくらでもあると思います。敬語はもちろんで、です、ますだけではなく言い方にも気を付けましょう。

## アサインされても気付かない問題
PR出したけど、なかなか見てくれない。しょうがなくダイレクトメールかチャットでURL送ってお願いする場合もありますね。この問題はツールで解決できる問題だと思います。一日2,3回自分にアサインされて見てないPRをメールか、チャットで催促など、方法はいくらでもあるので、問題ないはずです。しかし、気付いたけどわざと見ないのはその人達間の人間関係の問題で解決方法は別途探しましょう。

## 対象ではないことろに触れる問題
プロジェクト全体のコードが理想の形になってない状態で1つの小さなPRでプロジェクト全体に対してのツッコミが入る時があります。それは今回のPRと関係ないです！と返す人もプロジェクト全体のコードを綺麗にしたいレビュアーもどっちも悪くはない思います。やる気あって且つデキる人はわかった全部直してやる！黙々と今回のPRと関係ないところも修正してPRを出し直す人もいれば、この機能のリリースが先だ、そんなの知らないよ！と返す人もいると思います。この辺は優先順位を決めれる人を捕まえて話し合いが必要ですね。

## 土日レビューしてくれる人がいない問題
若者が（若者に限らないが）土日やる気出して1つの機能追加してリリースしたいのにレビュー出してもレビューしてくれそうな人がいない時があると思います。月曜日まで待ちたくない、今すぐリリースしたい、こういう時どうすればいいのか？自分でPR作って自分でマージする。コードレビューの意味がなくなってる。そもそもそういう若者をフォロー出来ないのは組織の問題なので、ここではその辺の話には触れたくないです。
