

# 在项目根目录下执行，构建镜像
docker build -t learn-gin .

# 构建镜像完成之后 可以查看
docker images

# 运行程序
docker run -p 8080:8080 learn-gin