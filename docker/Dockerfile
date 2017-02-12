FROM golang:1.6
MAINTAINER peter-mueller

# Install Chrome WebDriver (lib... required for chromedriver)
RUN apt-get update -yqq && \
    apt-get -yqq install curl unzip && \
    apt-get -yqq install xvfb && \
    apt-get -yqq install libxi6 libgconf-2-4 && \
    rm -rf /var/lib/apt/lists/*

RUN CHROMEDRIVER_VERSION=`curl -sS chromedriver.storage.googleapis.com/LATEST_RELEASE` && \
    mkdir -p /opt/chromedriver-$CHROMEDRIVER_VERSION && \
    curl -sS -o /tmp/chromedriver_linux64.zip http://chromedriver.storage.googleapis.com/$CHROMEDRIVER_VERSION/chromedriver_linux64.zip && \
    unzip -qq /tmp/chromedriver_linux64.zip -d /opt/chromedriver-$CHROMEDRIVER_VERSION && \
    rm /tmp/chromedriver_linux64.zip && \
    chmod +x /opt/chromedriver-$CHROMEDRIVER_VERSION/chromedriver && \
    ln -fs /opt/chromedriver-$CHROMEDRIVER_VERSION/chromedriver /usr/local/bin/chromedriver

# Install Asciidoctor and Chromium
RUN apt-get update -yqq && \
    apt-get install -yqq ruby asciidoctor chromium && \
    rm -rf /var/lib/apt/lists/*

# Get rhythmtool into the container
RUN go get -u -d -t github.com/peter-mueller/rhythmtool/...
RUN cd src/github.com/peter-mueller/rhythmtool

# Start Xvfb
ENV DISPLAY :99
CMD Xvfb :99 -nolisten tcp & \
    cd src/github.com/peter-mueller/rhythmtool && \
    bash
