import socket
import os
import threading

# 处理客户端请求
def handle_client_request(new_socket):
    # 代码执行到此，说明连接建立成功
    recv_client_data = new_socket.recv(4097)
    if len(recv_client_data) == 0:
        print('关闭浏览器了!')
        new_socket.close()
        return

    # 对二进制数据进行解码
    recv_client_content = recv_client_data.decode('utf-8')
    print(recv_client_content)
    # 根据指定字符串进行分割， 最大分割次数指定2
    request_list = recv_client_content.split(" ", maxsplit=2)

    # 获取请求资源路径
    request_path = request_list[1]
    print(request_path)

    # 判断请求的是否是根目录，如果是根目录设置返回的信息
    if request_path == "/":
        request_path = "/index.html"

    # 1. os.path.exits
    # os.path.exists("static/" + request_path)
    # 2. try-except
    try:
        # 打开文件读取文件中的数据, 提示：这里使用rb模式，兼容打开图片文件
        with open("static" + request_path, "rb") as file:  # 这里的file表示打开文件的对象
            file_data = file.read()
        # 提示： with open 关闭文件这步操作不用程序员来完成，系统帮我们来完成
    except Exception as e:
        # 代码执行到此，说明没有请求的该文件，返回404状态信息
        # 响应行
        response_line = "HTTP/1.1 404 Not Found\r\n"
        # 响应头
        response_header = "Server: PWS/1.0\r\n"
        # 读取404页面数据
        with open("static/error.html", "rb") as file:
            file_data = file.read()

        # 响应体
        response_body = file_data

        # 把数据封装成http 响应报文格式的数据
        response = (response_line +
                    response_header +
                    "\r\n").encode("utf-8") + response_body
        # 发送给浏览器的响应报文数据
        new_socket.send(response)

    else:
        # 代码执行到此，说明文件存在，返回200状态信息
        # 响应行
        response_line = "HTTP/1.1 200 OK\r\n"
        # 响应头
        response_header = "Server: PWS/1.0\r\n"
        # 响应体
        response_body = file_data

        # 把数据封装成http 响应报文格式的数据
        response = (response_line +
                    response_header +
                    "\r\n").encode("utf-8") + response_body

        # 发送给浏览器的响应报文数据
        new_socket.send(response)
    finally:
        # 关闭服务于客户端的套接字
        new_socket.close()

# 代码执行到此，说明没有请求的该文件，返回404状态信息
# 响应行


def main():
    # 创建tcp服务端套接字
    tcp_server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    # 设置端口号复用，程序退出端口号立即释放
    tcp_server_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, True)
    # 绑定端口号
    tcp_server_socket.bind(("", 8000))
    # 设置监听
    tcp_server_socket.listen(128)
    # 循环等待接受客户端的连接请求
    while True:
        # 等待接受客户端的连接请求
        new_socket, ip_port = tcp_server_socket.accept()
        # 代码执行到此，说明连接建立成功
        print(ip_port)
        # 当客户端和服务器建立连接程，创建子线程
        sub_thread = threading.Thread(target=handle_client_request, args=(new_socket,))
        # 设置守护主线程
        sub_thread.setDaemon(True)
        # 启动子线程执行对应的任务
        sub_thread.start()


# 判断是否是主模块的代码
if __name__ == '__main__':
    main()
