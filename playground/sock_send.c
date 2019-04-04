#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <sys/socket.h>

int main() {
	int fd = atoi(getenv("WASI_FD"));
	if (fd <= 0) {
		perror("could not find WASI_FD environment variable");
		return 1;
	}

	printf("WASI_FD = %d\n", fd);

	const char* msg = "Hello host! This is wasm program sending a message over Unix socket.        -> Over <-\n";
	int len = strlen(msg);

	if (send(fd, msg, len, 0) < 0) {
		perror("send error");
		return 1;
	}
	return 0;
}
