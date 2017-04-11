# stereogram-swapper
## これはなに？
立体視画像の左右を入れ替えることで交差法⇔平行法を変換するツール

## つかいかた
```
Usage of stereogram-swapper:
  -i string
        input directory (default ".")
  -o string
        output directory (default ".")
```

input directory内にある全てのjpg画像の左右を入れ替えてoutput directoryに出力する。

```
$ stereogram-swapper -i ./input -o ./output
```
