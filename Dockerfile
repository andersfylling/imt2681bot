# Docker only
FROM golang:1.8
MAINTAINER https://github.com/andersfylling

WORKDIR /go/src/github.com/andersfylling/imt2681bot
COPY . .

# Get Glide for package management
RUN curl https://glide.sh/get | sh
RUN glide install

# The URL needs to be replaced. Please see extra details given with this submission.
ENV IMT_BOT_TOKEN replace_me_using_a_valid_discord_bot_token
ENV IMT_BOT_COMMAND_PREFIX imt!
