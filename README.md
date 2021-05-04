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

## How to build
### Build dependencies
* cc
* pkg-config
* [telebot](https://github.com/tucnak/telebot)
* [tag](https://github.com/dhowden/tag)

To build on Linux or macOS systems, you need above dependencies pre-installed.

To install on Fedora / RHEL / CentOS
```
dnf install golang gcc pkg-config taglib-devel
```
To install on Ubuntu / Debian
```
apt install golang gcc pkg-config libtagc0-dev
```

Now you can start to build the go module
```
git clone https://github.com/star-39/moe_music_bot && cd moe_music_bot
go mode init moe_music_bot
go get -u gopkg.in/tucnak/telebot.v2
go get -u https://github.com/dhowden/tag
go build
```


## License
The GPL V3 License

![image](http://www.gnu.org/graphics/gplv3-127x51.png)
