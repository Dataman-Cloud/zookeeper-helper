docker run -d -e ZOOKEEPER_HELPER_HOST=10.3.10.36 -e ZOOKEEPER_HELPER_PORT=5096 --name zookeeper-helper  --net host testregistry.dataman.io/centos7/zookeeper-helper:omega.v0.11
