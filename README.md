# Android ICON converter(converticon)

## What is this?

Resize icon for playstore(android). 144x144, 96x96, ...

command interface.

```shell
$ converticon -p <project path> -i <image icon path>
```

for example,

```shell
$ converticon -p lifeticking -i lifeticking/icon.png
generate /Users/satouhayabusa/git/github.com/satoshun/lifeticking/app/src/main/res/drawable-xhdpi/ic_launcher.png
generate /Users/satouhayabusa/git/github.com/satoshun/lifeticking/app/src/main/res/drawable-hdpi/ic_launcher.png
generate /Users/satouhayabusa/git/github.com/satoshun/lifeticking/app/src/main/res/drawable-mdpi/ic_launcher.png
generate /Users/satouhayabusa/git/github.com/satoshun/lifeticking/app/src/main/res/drawable-xxhdpi/ic_launcher.png
```

very easy!


## How to install

Use Homebrew

```shell
$ brew tap satoshun/commands
$ brew install converticon
```


## Memo

- dependency ImageMagick


### License

Under MIT License(http://opensource.org/licenses/mit-license.php)
