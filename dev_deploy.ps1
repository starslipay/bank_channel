$MODULE_NAME = "bank_channel"
$VERSION = "v1.0.0"
$IMAGE_NAME = "${MODULE_NAME}:${VERSION}"

docker rm -f $MODULE_NAME
docker rmi -f $IMAGE_NAME
docker build -t $IMAGE_NAME .
docker run -d --name $MODULE_NAME --network dev_pay_net -p 30886:8080 -p 9101:9090 $IMAGE_NAME
# docker run -d --name bank_channel --network dev_pay_net -p 30886:8080 bank_channel:v1.0.0
docker ps
docker logs $MODULE_NAME
