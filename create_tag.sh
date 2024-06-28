#!/bin/bash

VERSION=$1
if [ ! "$VERSION" ] ;then
    echo "Version cannot empty."
    exit
fi
echo "Git tag version $VERSION."

git tag contrib/trace/otlphttp/$VERSION -m "add $VERSION tag for otlphttp"
git tag contrib/trace/otlpgrpc/$VERSION -m "add $VERSION tag for otlpgrpc"
git tag contrib/trace/jaeger/$VERSION -m "add $VERSION tag for jaeger"
git tag contrib/sdk/httpclient/$VERSION -m "add $VERSION tag for httpclient"
git tag contrib/rpc/grpcx/$VERSION -m "add $VERSION tag for grpcx"
git tag contrib/registry/zookeeper/$VERSION -m "add $VERSION tag for zookeeper"
git tag contrib/registry/polaris/$VERSION -m "add $VERSION tag for polaris"
git tag contrib/registry/nacos/$VERSION -m "add $VERSION tag for nacos"
git tag contrib/registry/file/$VERSION -m "add $VERSION tag for file"
git tag contrib/registry/etcd/$VERSION -m "add $VERSION tag for etcd"
git tag contrib/nosql/redis/$VERSION -m "add $VERSION tag for redis"
git tag contrib/metric/otelmetric/$VERSION -m "add $VERSION tag for otelmetric"
git tag contrib/drivers/sqlitecgo/$VERSION -m "add $VERSION tag for sqlitecgo"
git tag contrib/drivers/sqlite/$VERSION -m "add $VERSION tag for sqlite"
git tag contrib/drivers/pgsql/$VERSION -m "add $VERSION tag for pgsql"
git tag contrib/drivers/oracle/$VERSION -m "add $VERSION tag for oracle"
git tag contrib/drivers/mysql/$VERSION -m "add $VERSION tag for mysql"
git tag contrib/drivers/mssql/$VERSION -m "add $VERSION tag for mssql"
git tag contrib/drivers/dm/$VERSION -m "add $VERSION tag for dm"
git tag contrib/drivers/clickhouse/$VERSION -m "add $VERSION tag for clickhouse"
git tag contrib/config/polaris/$VERSION -m "add $VERSION tag for polaris"
git tag contrib/config/nacos/$VERSION -m "add $VERSION tag for nacos"
git tag contrib/config/kubecm/$VERSION -m "add $VERSION tag for kubecm"
git tag contrib/config/consul/$VERSION -m "add $VERSION tag for consul"
git tag contrib/config/apollo/$VERSION -m "add $VERSION tag for apollo"
git tag cmd/gf/$VERSION -m "add $VERSION tag for gf"
git tag cmd/gf/internal/cmd/testdata/build/varmap/$VERSION -m "add $VERSION tag for varmap"

echo "Git tag push..."

git push origin contrib/trace/otlphttp/$VERSION
git push origin contrib/trace/otlpgrpc/$VERSION
git push origin contrib/trace/jaeger/$VERSION
git push origin contrib/sdk/httpclient/$VERSION
git push origin contrib/rpc/grpcx/$VERSION
git push origin contrib/registry/zookeeper/$VERSION
git push origin contrib/registry/polaris/$VERSION
git push origin contrib/registry/nacos/$VERSION
git push origin contrib/registry/file/$VERSION
git push origin contrib/registry/etcd/$VERSION
git push origin contrib/nosql/redis/$VERSION
git push origin contrib/metric/otelmetric/$VERSION
git push origin contrib/drivers/sqlitecgo/$VERSION
git push origin contrib/drivers/sqlite/$VERSION
git push origin contrib/drivers/pgsql/$VERSION
git push origin contrib/drivers/oracle/$VERSION
git push origin contrib/drivers/mysql/$VERSION
git push origin contrib/drivers/mssql/$VERSION
git push origin contrib/drivers/dm/$VERSION
git push origin contrib/drivers/clickhouse/$VERSION
git push origin contrib/config/polaris/$VERSION
git push origin contrib/config/nacos/$VERSION
git push origin contrib/config/kubecm/$VERSION
git push origin contrib/config/consul/$VERSION
git push origin contrib/config/apollo/$VERSION
git push origin cmd/gf/$VERSION
git push origin cmd/gf/internal/cmd/testdata/build/varmap/$VERSION