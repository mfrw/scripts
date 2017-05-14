#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <string.h>
#include <sys/types.h>

int main() {
	int i;
	unsigned char toSend[102],mac[6];
	struct sockaddr_in udpClient, udpServer;
	int yes = 1 ;

	int udpSocket = socket(AF_INET, SOCK_DGRAM, 0);

	if (setsockopt(udpSocket, SOL_SOCKET, SO_BROADCAST, &yes, sizeof(yes)) == -1) {
		perror("setsockopt (SO_BROADCAST)");
		exit(EXIT_FAILURE);
	}
	udpClient.sin_family = AF_INET;
	udpClient.sin_addr.s_addr = INADDR_ANY;
	udpClient.sin_port = 0;

	//Binding the socket
	bind(udpSocket, (struct sockaddr*)&udpClient, sizeof(udpClient));

	for (i=0; i<6; i++)
		toSend[i] = 0xFF;

	mac[0] = 0xc4;  // 1st octet of the MAC Address
	mac[1] = 0x54;  // 2nd octet of the MAC Address
	mac[2] = 0x44;  // 3rd octet of the MAC Address
	mac[3] = 0x03;  // 4th octet of the MAC Address
	mac[4] = 0xf0;  // 5th octet of the MAC Address
	mac[5] = 0xeb;  // 6th octet of the MAC Address

	for (i=1; i<=16; i++)
		memcpy(&toSend[i*6], &mac, 6*sizeof(unsigned char));

	udpServer.sin_family = AF_INET;

	// Braodcast address
	udpServer.sin_addr.s_addr = inet_addr("192.168.79.255");
	udpServer.sin_port = htons(9);

	sendto(udpSocket, &toSend, sizeof(unsigned char) * 102, 0, (struct sockaddr*)&udpServer, sizeof(udpServer));
	return 0;
}
