from flask import Flask
from redis import Redis,RedisError
import os
import socket

# Connect to Redis   连接redis数据库
redis = Redis(host="redis", db=0,socket_connect_timeout=2, socket_timeout=2)

app = Flask(__name__)

# 访问这台机器的根"/"
@app.route("/")
def hello():
	try:
		# 若是有人访问，会往redis数据库里增加一个值
		visits = redis.incr("counter")
	except RedisError:
		# 这个visits是从redis数据库里获取的
		visits = "<i>cannot connect to REdis, counter disabled</i>"

	html = "<h3>Hello {name}!</h3>" \
			"<b>Hostname:</b>{hostname}<br/>" \
			"<b>Visits:</b> {visits}"
	# 返回这个主机名(socket.gethostname())和访问的次数(visits)
	return html.format(name=os.getenv("NAME", "world"), hostname=socket.gethostname(),visits=visits)

if __name__ == "__main__":
	app.run(host='0.0.0.0', port=80)