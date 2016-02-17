FROM alpine:3.3 
MAINTAINER Will <zhguo.dataman-inc.com>

ADD zookeeper-helper.yaml.sample /etc/omega/zookeeper-helper.yaml
ADD . /go/src/github.com/Dataman-Cloud/zookeeper-helper

RUN cd /go/src/github.com/Dataman-Cloud/zookeeper-helper && sh build.sh
VOLUME /var/log/zookeeper-helper

EXPOSE 5096 

CMD /zookeeper-helper
