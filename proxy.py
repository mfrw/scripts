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
        thread = threading.Thread(target=proxy, args=(csocket, rhost, rport, recieve_first))
        thread.start()

def proxy(csocket, rhost, rport, recieve_first):
    rsocket = socket.socket()
    rsocket.connect((rhost, rport))
    if recieve_first:
        rbuff = receive_from(rsocket)
        hexdump(rbuff)
        rbuff = response_handler(rbuff)
        if len(rbuff):
            print('[<==] Sending %d bytes to localhost.'%len(rbuff))
            csocket.send(rbuff)
    while True:
        lbuff = receive_from(csocket)
        if len(lbuff):
            print('[==>] Recieved %d bytes from localhost.'%len(lbuff))
            hexdump(lbuff)
        lbuff = request_handler(lbuff)
        rsocket.send(lbuff)
        print('[==>] Sent to remote.')
        rbuff = receive_from(rsocket)
        if(len(rbuff)):
            print('[<==] Received %d bytes from remote.' %len(rbuff))
            hexdump(rbuff)
            rbuff = response_handler(rbuff)
            csocket.send(rbuff)
            print('[<==] Sent to localhost.')
        if not len(lbuff) or not len(rbuff):
            csocket.close()
            rsocket.close()
            print('[*] No more data. Closing connection.')
            break

def hexdump(src, length=16):
    result = []
    digits = 4 if isinstance(src, unicode) else 2
    for i in xrange(0, len(src), length):
        s = src[i: i+length]
        hexa = b' '.join(['%o*X'%(digits, ord(x)) for x in s])
        text = b''.join([x if 0x20 <= ord(x) < 0x7f else b'.' for x in x])
        result.append(b'%o4X %-*s %s'%(i, length*(digits + 1), hexa, text))
    print b'\n'.join(result)

def receive_from(connection):
    buf = ''
    connection.settimeout(2)
    try:
        while True:
            data = connection.recv(4096)
            if not data:
                break
            buf += data
    except:
        pass
    return buf

def request_handler(buff):
    return buff

def response_handler(buf):
    return buf

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
