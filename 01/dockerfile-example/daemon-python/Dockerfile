FROM python

LABEL author="librant"

# 我们在 docker 容器的工作目录。即进入 docker 容器的时候，所在的目录
WORKDIR /app

# 复制当前目录下的所有内容到容器里边 /app 目录下
ADD . /app

# 定义一个数据卷
VOLUME ["/data_flask"]

# requirements.txt会写明我这个项目会需要一些什么python库；\
# 这条命令是告诉容器把requirements.txt里需要的库到pypi.python.org这个网站里边下载
# 这条命令是在制作镜像的时候要运行的，而不是在启动容器的时候运行的，这个要记住很容易混淆
RUN pip install  --trusted-host pypi.python.org -r requirements.txt

# 暴露我们的端口号，即容器对外开放哪个端口
EXPOSE 80

# 定义环境变量
ENV NAME World

# 容器起来之后运行的第一个程序，这个是容器启动的时候运行，要和上边区分开来
CMD ["python", "app.py"]