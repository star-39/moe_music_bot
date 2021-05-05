GO=go


all: build

init:
	$(GO) mod init moe_music_bot

build:
	$(GO) build

with-ffmpeg:
	$(GO) build -tags embed_ffmpeg -ldflags="-X main.embedFFMpeg=ON"

clean:
	rm moe_music_bot
