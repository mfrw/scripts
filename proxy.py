import sys
import socket
import threading

def server_loop(lhost, lport, rhost, rport, recieve_first):
    server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    try:
        server.bind((lhost, lport))
    except:
        print('[!!] Failed to listen on %s:%d'%(lhost, lport))
        print('[!!] Check for other listening sockets or correct permissions.')
        sys.exit(0)

    print('[*] Listening on %s:%d'%(lhost, lport))

    server.listen(5)
    while True:
        csocket, addr = server.accept()
        print('[==>] Recieved incoming connection from %s:%d'%(addr[0], addr[1]))
        thread = threading.Thread(target=handle, args=(csocket, rhost, rport, recieve_first))
        thread.start()

def main():
    if len(sys.argv[1:]) != 5:
        print("Usage ./proxy [localhost] [localport] [remotehost] [remoteport] [revieve_first]")
        print("Example: ./proxy 127.0.0.1 9000 10.12.132.1 9000 True")
        sys.exit(0)
    lhost = sys.argv[1]
    lport = int(sys.argv[2])

    rhost = sys.argv[3]
    rport = int(sys.argv[4])

    recieve_first = sys.argv[5]

    if 'True' in recieve_first:
        recieve_first = True
    else:
        recieve_first = False

    server_loop(lhost, lport, rhost, rport, recieve_first)

if __name__ == '__main__':
    main()
