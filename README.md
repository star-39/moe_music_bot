# moe_music_bot
Get all your musics to Telegram without effort!

## Dependencies
* cc
* pkg-config
* taglib
* [telebot](https://github.com/tucnak/telebot)

## Usage
Download pre-built binary in [Release](https://github.com/star-39/moe_music_bot/releases)
```
./moe_music_bot [YOUR_BOT_API]
```
Please note that it runs ONLY on Linux or macOS systems!

If you are using Windows, please use [WSL](https://docs.microsoft.com/en-us/windows/wsl/install-win10).


## How to build
To build on Linux or macOS systems, you need above dependencies pre-installed.

To install on Fedora / RHEL / CentOS
```
dnf install gcc pkg-config taglib-devel
```
To install on Ubuntu / Debian
```
apt install gcc pkg-config libtagc0-dev
```

Now you can start to build the go module
```
git clone https://github.com/star-39/moe_music_bot && cd moe_music_bot
go get -u gopkg.in/tucnak/telebot.v2
go build
```


## License
The GPL V3 License

![image](http://www.gnu.org/graphics/gplv3-127x51.png)
