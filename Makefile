GO=go


all: build

init:
	$(GO) mod init moe_music_bot

build:
	$(GO) build

with-ffmpeg:
	$(GO) build -tags embed_ffmpeg

clean:
	rm moe_music_bot
