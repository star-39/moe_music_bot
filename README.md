# moe_music_bot
Get all your musics to Telegram without effort!

![Screenshot 2021-05-04 184640](https://user-images.githubusercontent.com/75669297/116986761-290d9900-ad09-11eb-949b-ce7fa025ab3e.png)


## Usage
Download pre-built binary in [Release](https://github.com/star-39/moe_music_bot/releases)
```
./moe_music_bot [YOUR_BOT_API]
```

## Telegram Use Example
> /help

> /set transcode yes

> /set bitrate 256k

> /up /home/wsl/Music

## How to Build
* Go 1.13 or above is required, if you want to embed ffmpeg, 1.16 or above is required.
```
git clone https://github.com/star-39/moe_music_bot && cd moe_music_bot
go mode init moe_music_bot
go get -u gopkg.in/tucnak/telebot.v2
go get -u https://github.com/dhowden/tag
go build
```

If you are on Linux or macOS, you can simply use the Makefile
```
make
```

You can also embed a ffmpeg executable binary, put file `ffmpeg` in current directory and run:
```
go build -tags embed_ffmpeg
```
or
```
make with-ffmpeg
```


## License
The GPL V3 License

![image](http://www.gnu.org/graphics/gplv3-127x51.png)
